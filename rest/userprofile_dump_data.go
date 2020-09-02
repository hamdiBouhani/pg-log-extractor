package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/common"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/config"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/types"
)

func (s *Server) DumpData(
	session *types.Session,
	dataset string,
	bigQueryTable string,
	tableName string,
	ColumnName string,
	order string,
	fn common.ParseValueFunc,
) error {

	Data, err := snapshotData(session, &types.SnapshotDataJSON{
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

func (s *Server) DumpUserProfileIntoBigQuery(c *gin.Context) {

	go func() {
		session, err := config.InitSession(s.DbName, s.PgUser, s.PgPass, s.PgHost, s.PgPort)
		if err != nil {
			e := fmt.Sprintf("unable to init session")
			s.Logger.WithError(err).Error(e)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": e,
			})
			return
		}
		s.DumpData(session, DatasetID, "user_profile", "user_profile", "full_name", "ASC", ParseValueToUserProfile)

	}()

	common.ResponseData(c, true)
}
func (s *Server) DumpCareerAspirationIntoBigQuery(c *gin.Context) {
	go func() {
		session, err := config.InitSession(s.DbName, s.PgUser, s.PgPass, s.PgHost, s.PgPort)
		if err != nil {
			e := fmt.Sprintf("unable to init session")
			s.Logger.WithError(err).Error(e)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": e,
			})
			return
		}
		s.DumpData(session, DatasetID, "user_profile_career_aspiration", "user_career_aspiration", "user_id", "ASC", ParseValueToUserCareerAspiration)

	}()

	common.ResponseData(c, true)
}

func (s *Server) DumpExperienceIntoBigQuery(c *gin.Context) {
	go func() {
		session, err := config.InitSession(s.DbName, s.PgUser, s.PgPass, s.PgHost, s.PgPort)
		if err != nil {
			e := fmt.Sprintf("unable to init session")
			s.Logger.WithError(err).Error(e)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": e,
			})
			return
		}
		s.DumpData(session, DatasetID, "experience", "experience", "user_id", "ASC", ParseValueToExperience)

	}()
	common.ResponseData(c, true)
}

func (s *Server) DumpUserCompetencyIntoBigQuery(c *gin.Context) {

	go func() {
		session, err := config.InitSession(s.DbName, s.PgUser, s.PgPass, s.PgHost, s.PgPort)
		if err != nil {
			e := fmt.Sprintf("unable to init session")
			s.Logger.WithError(err).Error(e)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": e,
			})
			return
		}
		s.DumpData(session, DatasetID, "user_competency", "user_competency", "competency_id", "ASC", ParseValueToUserCompetency)
	}()
	common.ResponseData(c, true)
}

func (s *Server) DumpEducationSpecializationIntoBigQuery(c *gin.Context) {

	go func() {
		session, err := config.InitSession(s.DbName, s.PgUser, s.PgPass, s.PgHost, s.PgPort)
		if err != nil {
			e := fmt.Sprintf("unable to init session")
			s.Logger.WithError(err).Error(e)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": e,
			})
			return
		}
		s.DumpData(session, DatasetID, "education_specialization", "education_specialization", "major_subject", "ASC", ParseValueToEducationSpecialization)
	}()
	common.ResponseData(c, true)
}

func (s *Server) DumpUserLanguageIntoBigQuery(c *gin.Context) {
	go func() {
		session, err := config.InitSession(s.DbName, s.PgUser, s.PgPass, s.PgHost, s.PgPort)
		if err != nil {
			e := fmt.Sprintf("unable to init session")
			s.Logger.WithError(err).Error(e)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": e,
			})
			return
		}
		s.DumpData(session, DatasetID, "language_bridge", "user_language", "user_id", "ASC", ParseValueToUserLanguage)
	}()
	common.ResponseData(c, true)
}

func (s *Server) DumpUserEducationIntoBigQuery(c *gin.Context) {
	go func() {
		session, err := config.InitSession(s.DbName, s.PgUser, s.PgPass, s.PgHost, s.PgPort)
		if err != nil {
			e := fmt.Sprintf("unable to init session")
			s.Logger.WithError(err).Error(e)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": e,
			})
			return
		}
		s.DumpData(session, DatasetID, "user_education", "user_education", "user_id", "ASC", ParseValueToUserEducation)
	}()
	common.ResponseData(c, true)
}

func (s *Server) DumpDegreeLevelIntoBigQuery(c *gin.Context) {
	go func() {
		session, err := config.InitSession(s.DbName, s.PgUser, s.PgPass, s.PgHost, s.PgPort)
		if err != nil {
			e := fmt.Sprintf("unable to init session")
			s.Logger.WithError(err).Error(e)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": e,
			})
			return
		}
		s.DumpData(session, DatasetID, "degree_level", "degree_level", "degree_name", "ASC", ParseValueToDegreeLevel)
	}()

	common.ResponseData(c, true)
}
