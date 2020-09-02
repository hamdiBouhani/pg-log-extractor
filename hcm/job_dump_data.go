package hcm

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/common"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/config"
)

func (s *Server) DumpJobIntoBigQuery(c *gin.Context) {
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
		s.DumpData(session, "hcm", "job", "job", "job_title", "ASC", ParseValueToJob)

	}()

	common.ResponseData(c, true)
}

func (s *Server) DumpJobCompetencyIntoBigQuery(c *gin.Context) {
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
		s.DumpData(session, "hcm", "job_competency", "job_competency", "job_id", "ASC", ParseValueToJobCompetency)

	}()

	common.ResponseData(c, true)
}

func (s *Server) DumpJobEducationSpecializationIntoBigQuery(c *gin.Context) {
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
		s.DumpData(session, "hcm", "job_education_specialization", "job_education_specialization", "job_id", "ASC", ParseValueToJobEducationSpecialization)

	}()

	common.ResponseData(c, true)
}

func (s *Server) DumpJobLanguageIntoBigQuery(c *gin.Context) {
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
		s.DumpData(session, "hcm", "job_language", "job_language", "job_ID", "ASC", ParseValueToJobLanguage)

	}()

	common.ResponseData(c, true)
}

func (s *Server) DumpJobNationalityIntoBigQuery(c *gin.Context) {
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
		s.DumpData(session, "hcm", "job_nationality", "job_nationality", "job_id", "ASC", ParseValueToJobNationality)

	}()

	common.ResponseData(c, true)
}

func (s *Server) DumpJobDegreeLevelIntoBigQuery(c *gin.Context) {
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
		s.DumpData(session, "hcm", "degree_level", "degree_level", "degree_name", "ASC", ParseValueToDegreeLevel)
	}()

	common.ResponseData(c, true)
}
