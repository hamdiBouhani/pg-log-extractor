package profile

import (
	"context"

	"cloud.google.com/go/bigquery"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
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

func NewServer(
	logger *logrus.Logger,
	dbName string,
	pgUser string,
	pgPass string,
	pgHost string,
	pgPort int,
	projectID string, /* Google Cloud Platform project ID.*/
	serviceAccount string, /*"service-account Path to Service Account .json file"*/

) (*Server, error) {
	if logger == nil {
		logger = logrus.New()
		logger.SetLevel(logrus.InfoLevel)
	}

	config.CreateConfig(dbName, pgUser, pgPass, pgHost, pgPort)

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
	session.BigQueryClient = bigQueryClient

	v1 := r.Group("/v1/api")
	{

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

	}

	return server, nil
}
