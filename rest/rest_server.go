package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/bigquery"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/config"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/types"
)

type Server struct {
	*logrus.Logger
	*gin.Engine

	BigQueryClient *bigquery.Client

	DbName string
	PgUser string
	PgPass string
	PgHost string
	PgPort int
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewServer(
	logger *logrus.Logger,
	dbName string,
	pgUser string,
	pgPass string,
	pgHost string,
	pgPort int,
	topicName string, /*The cloud pubsub topic to use*/
	projectID string, /* Google Cloud Platform project ID.*/
	serviceAccount string, /*"service-account Path to Service Account .json file"*/
	mongoAddr string, /*eg: mongodb://localhost:27017 OR mongodb://mongo-rs0-1,mongo-rs0-2,mongo-rs0-3/?replicaSet=rs0*/
	repSetName string,
	mongoHosts string,
) (*Server, error) {
	if logger == nil {
		logger = logrus.New()
		logger.SetLevel(logrus.InfoLevel)
	}

	config.CreateConfig(dbName, pgUser, pgPass, pgHost, pgPort)

	PSConn := config.NewPubSubCon(topicName, projectID, serviceAccount)
	var Mconn *types.MongoConn
	if mongoAddr != "" && len(mongoAddr) > 0 {
		Mconn = config.NewMongoConn(mongoAddr, repSetName, strings.Split(mongoHosts, ","))
	}

	bigQueryClient, err := bigquery.NewClient(context.Background(), projectID)
	if err != nil {
		return nil, err
	}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// TODO uncomment + pass the argCORSHosts to the Header instead of *
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Mode, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	pprof.Register(r)
	// Create server
	server := &Server{
		Logger:         logger,
		Engine:         r,
		BigQueryClient: bigQueryClient,
		DbName:         dbName,
		PgUser:         pgUser,
		PgPass:         pgPass,
		PgHost:         pgHost,
		PgPort:         pgPort,
	}

	//WebSocket Test Client
	dir, err := os.Getwd()
	if err != nil {
		logrus.Info("unable to find test.html file !!!")
		log.Fatal(err)
	}

	r.LoadHTMLFiles(filepath.Join(dir, "rest", "view", "mongo-stream.html"), filepath.Join(dir, "rest", "view", "pg-stream.html"))

	r.GET("/pg-stream", func(c *gin.Context) {
		c.HTML(200, "pg-stream.html", nil)
	})

	r.GET("/mongo-stream", func(c *gin.Context) {
		c.HTML(200, "mongo-stream.html", nil)
	})

	// open accessed group
	openAccessed := r.Group("/")
	{
		// service info handler
		openAccessed.GET("/info", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"info": "UP",
			})
		})
	}
	session := &types.Session{}
	session.PSConn = PSConn
	session.BigQueryClient = bigQueryClient

	v1 := r.Group("/v1/api")
	{
		v1.GET("/init", func(c *gin.Context) {
			err := initDB(session)
			if err != nil {
				e := fmt.Sprintf("unable to init session")
				logger.WithError(err).Error(e)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": e,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"slotName": session.SlotName,
			})
		})

		snapshotRoute := v1.Group("/snapshot")
		{
			snapshotRoute.POST("/data", func(c *gin.Context) {
				if session.SnapshotName == "" {
					e := fmt.Sprintf("snapshot not available: call /init to initialize a new slot and snapshot")
					logger.Error(e)
					c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
						"error": e,
					})
					return
				}

				// get data with table, offset, limits
				var postData types.SnapshotDataJSON
				//err = c.MustBindWith(&postData, binding.JSON)
				err := c.ShouldBindJSON(&postData)
				if err != nil {
					e := fmt.Sprintf("invalid input JSON")
					logger.WithError(err).Error(e)
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": e,
					})
					return
				}

				err = validateSnapshotDataJSON(&postData)
				if err != nil {
					e := err.Error()
					logger.WithError(err).Error(e)
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": e,
					})
					return
				}

				logger.Infof("Snapshot data requested for table: %s, offset: %d, limit: %d", postData.Table, *postData.Offset, *postData.Limit)

				data, err := snapshotData(session, &postData)
				if err != nil {
					e := fmt.Sprintf("unable to get snapshot data")
					logger.WithError(err).Error(e)
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": e,
					})
					return
				}

				c.JSON(http.StatusOK, data)
			})
		}

		lrRoute := v1.Group("/lr")
		{
			lrRoute.GET("/stream", func(c *gin.Context) {
				slotName := c.Query("slotName")
				if slotName == "" {
					e := fmt.Sprintf("no slotName provided")
					logger.WithError(nil).Error(e)
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": e,
					})
					return
				}

				logger.Info("LR Stream requested for slot ", slotName)
				session.SlotName = slotName

				// now upgrade the HTTP  connection to a websocket connection
				wsConn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
				if err != nil {
					e := fmt.Sprintf("could not upgrade to websocket connection")
					logger.WithError(err).Error(e)
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": e,
					})
					return
				}
				session.WSConn = wsConn

				// begin streaming
				err = lrStream(session)
				if err != nil {
					e := fmt.Sprintf("could not create stream")
					logger.WithError(err).Error(e)
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": e,
					})
					return
				}
			})
		}

		lrUserProfile := v1.Group("user_profile/bigquery")
		{
			lrUserProfile.GET("/init", server.InitUserProfileBigQueryShema)
			lrUserProfile.GET("/Dump-user-profile", server.DumpUserProfileIntoBigQuery)
			lrUserProfile.GET("/Dump-career-aspiration", server.DumpCareerAspirationIntoBigQuery)
			lrUserProfile.GET("/Dump-experience", server.DumpExperienceIntoBigQuery)
			lrUserProfile.GET("/Dump-user-competency", server.DumpUserCompetencyIntoBigQuery)
			lrUserProfile.GET("/Dump-education-specialization", server.DumpEducationSpecializationIntoBigQuery)
			lrUserProfile.GET("/Dump-user-language", server.DumpUserLanguageIntoBigQuery)
			lrUserProfile.GET("/Dump-user-education", server.DumpUserEducationIntoBigQuery)
			lrUserProfile.GET("/Dump-degree-level", server.DumpDegreeLevelIntoBigQuery)
			lrUserProfile.GET("/Stream", server.StreamUserProfileIntoBigQuery)
		}
		mdbRoute := v1.Group("/mdb")
		{
			mdbRoute.GET("/stream", func(c *gin.Context) {
				database := c.Query("database")
				if database == "" {
					e := fmt.Sprintf("no database provided")
					logger.WithError(nil).Error(e)
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": e,
					})
					return
				}
				collection := c.Query("collection")
				if collection == "" {
					e := fmt.Sprintf("no collection provided")
					logger.WithError(nil).Error(e)
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": e,
					})
					return
				}
				// now upgrade the HTTP  connection to a websocket connection
				wsConn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
				if err != nil {
					e := fmt.Sprintf("could not upgrade to websocket connection")
					logger.WithError(err).Error(e)
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": e,
					})
					return
				}

				session.MbdConn = Mconn

				db := session.MbdConn.MongoClt.Database(database)
				if db == nil {
					e := fmt.Sprintf("faild to create  db")
					logger.WithError(errors.New("error : ")).Error(e)
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": e,
					})
					return
				}
				ctx := c.Request.Context()

				cs, err := db.Collection(collection).Watch(ctx, mongo.Pipeline{})
				if err != nil {
					e := fmt.Sprintf("could not create change stream")
					logger.WithError(err).Error(e)
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": e,
					})
					return
				}

				defer cs.Close(ctx)
				for {
					ok := cs.Next(ctx)
					if ok == false {
						break
					}
					next := cs.Current.String()
					b := []byte(next)
					// send message over ws
					wsConn.WriteMessage(websocket.BinaryMessage, b)
				}

				/** OR
					var doc bson.M
					for cur.Next(ctx) {
						if err = cur.Decode(&doc); err != nil {
							log.Fatal(err)
						}
						cb(doc)
					}
				**/

			})
		}
	}

	return server, nil
}
