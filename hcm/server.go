package hcm

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/common"
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

func (s *Server) DumpData(
	session *types.Session,
	dataset string,
	bigQueryTable string,
	tableName string,
	ColumnName string,
	order string,
	fn common.ParseValueFunc,
) error {

	Data, err := config.SnapshotData(session, &types.SnapshotDataJSON{
		SlotName: session.SlotName,
		Table:    tableName,
		OrderBy: &types.OrderBy{
			Column: ColumnName,
			Order:  order,
		},
		Limit:  nil,
		Offset: nil,
	})
	if err != nil {
		e := fmt.Sprintf("unable to get snapshot data")
		s.Logger.WithError(err).Error(e)
		return errors.Wrap(err, "unable to get snapshot data")
	}

	inseter := s.BigQueryClient.Dataset(dataset).Table(bigQueryTable).Inserter()
	for _, d := range Data {
		item := fn(d)

		if err := inseter.Put(context.Background(), item); err != nil {
			e := fmt.Sprintf("unable to insert %s data", tableName)
			s.Logger.WithError(err).Error(e)
			return errors.Wrap(err, "unable to get snapshot data")
		}
	}

	return nil
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
		lrUserProfile := v1.Group("hcm/bigquery")
		{
			lrUserProfile.GET("/init", server.InitJobBigQuerySchema)
			lrUserProfile.GET("/Dump-job", server.DumpJobIntoBigQuery)
			lrUserProfile.GET("/Dump-job-competency", server.DumpJobIntoBigQuery)
			lrUserProfile.GET("/Dump-job-education-specialization", server.DumpJobEducationSpecializationIntoBigQuery)
			lrUserProfile.GET("/Dump-job-language", server.DumpJobLanguageIntoBigQuery)
			lrUserProfile.GET("/Dump-job-nationality", server.DumpJobNationalityIntoBigQuery)
			lrUserProfile.GET("/Dump-degree-level", server.DumpJobDegreeLevelIntoBigQuery)
			lrUserProfile.GET("/Stream", server.StreamJobIntoBigQuery)
		}
	}

	return server, nil
}
