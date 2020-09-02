package hcm

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/bigquery"
	"github.com/gin-gonic/gin"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/common"
)

func (s *Server) InitJobBigQuerySchema(c *gin.Context) {

	jobSchema := bigquery.Schema{
		{Name: "id", Type: bigquery.IntegerFieldType, Required: true},         // id
		{Name: "job_title", Type: bigquery.StringFieldType},                   // job_title
		{Name: "job_description", Type: bigquery.StringFieldType},             // job_description
		{Name: "skills", Type: bigquery.StringFieldType},                      // skills
		{Name: "functional_area", Type: bigquery.StringFieldType},             // functional_area
		{Name: "total_position", Type: bigquery.IntegerFieldType},             // total_position
		{Name: "shift", Type: bigquery.StringFieldType},                       // shift
		{Name: "gender", Type: bigquery.StringFieldType},                      // gender
		{Name: "specific_degree_title", Type: bigquery.StringFieldType},       // specific_degree_title
		{Name: "major_subject", Type: bigquery.StringFieldType},               // major_subject
		{Name: "minor_subject", Type: bigquery.StringFieldType},               // minor_subject
		{Name: "responsibilities", Type: bigquery.StringFieldType},            // responsibilities
		{Name: "min_experience", Type: bigquery.IntegerFieldType},             // min_experience
		{Name: "max_experience", Type: bigquery.IntegerFieldType},             // max_experience
		{Name: "experience_detail", Type: bigquery.StringFieldType},           // experience_detail
		{Name: "is_hide_salary", Type: bigquery.BooleanFieldType},             // is_hide_salary
		{Name: "gpa", Type: bigquery.FloatFieldType},                          // gpa
		{Name: "out_of", Type: bigquery.FloatFieldType},                       // out_of
		{Name: "grade", Type: bigquery.StringFieldType},                       // grade
		{Name: "percentage", Type: bigquery.FloatFieldType},                   // percentage
		{Name: "primary_location", Type: bigquery.StringFieldType},            // primary_location
		{Name: "other_location", Type: bigquery.StringFieldType},              // other_location
		{Name: "posting_date", Type: bigquery.StringFieldType},                // posting_date
		{Name: "applied_before", Type: bigquery.StringFieldType},              // applied_before
		{Name: "activities", Type: bigquery.StringFieldType},                  // activities
		{Name: "challanges", Type: bigquery.StringFieldType},                  // challanges
		{Name: "company_id", Type: bigquery.IntegerFieldType},                 // company_id
		{Name: "posting_immediatly", Type: bigquery.BooleanFieldType},         // posting_immediatly
		{Name: "min_age", Type: bigquery.IntegerFieldType},                    // min_age
		{Name: "max_age", Type: bigquery.IntegerFieldType},                    // max_age
		{Name: "is_publish", Type: bigquery.BooleanFieldType},                 // is_publish
		{Name: "is_active", Type: bigquery.BooleanFieldType},                  // is_active
		{Name: "created_by", Type: bigquery.IntegerFieldType},                 // created_by
		{Name: "created_date", Type: bigquery.StringFieldType},                // created_date
		{Name: "changed_date", Type: bigquery.StringFieldType},                // changed_date
		{Name: "deleted_date", Type: bigquery.StringFieldType},                // deleted_date
		{Name: "vacancy_code", Type: bigquery.IntegerFieldType},               // vacancy_code
		{Name: "vacancy_referense_number", Type: bigquery.IntegerFieldType},   // vacancy_referense_number
		{Name: "male_required_count", Type: bigquery.IntegerFieldType},        // male_required_count
		{Name: "female_required_count", Type: bigquery.IntegerFieldType},      // female_required_count
		{Name: "both_gender_required_count", Type: bigquery.IntegerFieldType}, // both_gender_required_count
		{Name: "university_id", Type: bigquery.IntegerFieldType},              // university_id
		{Name: "degree_id", Type: bigquery.IntegerFieldType},                  // degree_id
		{Name: "specialization_id", Type: bigquery.IntegerFieldType},          // specialization_id
		{Name: "is_private", Type: bigquery.BooleanFieldType},                 // is_private
		{Name: "is_publicmatch_only", Type: bigquery.BooleanFieldType},        // is_publicmatch_only
		{Name: "industry_id", Type: bigquery.IntegerFieldType},                // industry_id
		{Name: "employment_type_id", Type: bigquery.IntegerFieldType},         // employment_type_id
		{Name: "career_level_id", Type: bigquery.IntegerFieldType},            // career_level_id
		{Name: "min_salary", Type: bigquery.FloatFieldType},                   // min_salary
		{Name: "max_salary", Type: bigquery.FloatFieldType},                   // max_salary
		{Name: "currency_salary", Type: bigquery.StringFieldType},             // currency_salary
		{Name: "country_id", Type: bigquery.IntegerFieldType},                 // country_id
		{Name: "state_id", Type: bigquery.IntegerFieldType},                   // state_id
		{Name: "city_id", Type: bigquery.IntegerFieldType},                    // city_id
		{Name: "job_no", Type: bigquery.StringFieldType},                      // job_no
		{Name: "job_type_id", Type: bigquery.IntegerFieldType},                // job_type_id
		{Name: "creator_subject", Type: bigquery.StringFieldType},             // creator_subject
		{Name: "company_name", Type: bigquery.StringFieldType},                // company_name
		{Name: "account_id", Type: bigquery.IntegerFieldType},
		{Name: "contact_ids", Type: bigquery.StringFieldType},
		{Name: "org_id", Type: bigquery.IntegerFieldType},
	}

	jobCompetencySchema := bigquery.Schema{
		{Name: "id", Type: bigquery.IntegerFieldType, Required: true}, // id
		{Name: "job_id", Type: bigquery.IntegerFieldType},             // job_id
		{Name: "level_id", Type: bigquery.IntegerFieldType},           // level_id
		{Name: "competency_id", Type: bigquery.IntegerFieldType},      // competency_id
		{Name: "is_gained", Type: bigquery.BooleanFieldType},          // is_gained
		{Name: "created_date", Type: bigquery.StringFieldType},        // created_date
		{Name: "changed_date", Type: bigquery.StringFieldType},        // changed_date
		{Name: "deleted_date", Type: bigquery.StringFieldType},        // deleted_date
	}

	jobEducationSpecializationSchema := bigquery.Schema{
		{Name: "id", Type: bigquery.IntegerFieldType, Required: true}, // id
		{Name: "job_id", Type: bigquery.IntegerFieldType},             // job_id
		{Name: "specialization_id", Type: bigquery.IntegerFieldType},  // specialization_id
		{Name: "created_date", Type: bigquery.StringFieldType},        // created_date
		{Name: "changed_date", Type: bigquery.StringFieldType},        // changed_date
		{Name: "deleted_date", Type: bigquery.StringFieldType},        // deleted_date
	}

	jobLanguageSchema := bigquery.Schema{
		{Name: "id", Type: bigquery.IntegerFieldType, Required: true}, // id
		{Name: "job_id", Type: bigquery.IntegerFieldType},             // job_id
		{Name: "language_id", Type: bigquery.IntegerFieldType},        // language_id
		{Name: "created_date", Type: bigquery.StringFieldType},        // created_date
		{Name: "changed_date", Type: bigquery.StringFieldType},        // changed_date
		{Name: "deleted_date", Type: bigquery.StringFieldType},        // deleted_date
	}

	jobNationalitySchema := bigquery.Schema{
		{Name: "id", Type: bigquery.IntegerFieldType, Required: true}, // id
		{Name: "job_id", Type: bigquery.IntegerFieldType},             // job_id
		{Name: "country_id", Type: bigquery.IntegerFieldType},         // country_id
		{Name: "created_date", Type: bigquery.StringFieldType},        // created_date
		{Name: "changed_date", Type: bigquery.StringFieldType},        // changed_date
		{Name: "deleted_date", Type: bigquery.StringFieldType},        // deleted_date
	}

	degreeLevelSchema := bigquery.Schema{
		{Name: "id", Type: bigquery.IntegerFieldType, Required: true},
		{Name: "degree_name", Type: bigquery.StringFieldType},
		{Name: "description", Type: bigquery.StringFieldType},
		{Name: "created_date", Type: bigquery.StringFieldType},
		{Name: "changed_date", Type: bigquery.StringFieldType},
		{Name: "deleted_date", Type: bigquery.StringFieldType},
		{Name: "years_of_studies", Type: bigquery.IntegerFieldType},
	}

	ctx := context.Background()

	meta := &bigquery.DatasetMetadata{
		Location: "US", // Create the dataset in the US.
	}
	if err := s.BigQueryClient.Dataset("hcm").Create(ctx, meta); err != nil {
		e := fmt.Sprintf("Failed to Create bigQuery dataset")
		s.Logger.WithError(err).Error(e)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": e,
		})
		return
	}

	jobMetaData := &bigquery.TableMetadata{
		Schema: jobSchema,
	}
	jobRef := s.BigQueryClient.Dataset("hcm").Table("job")
	if err := jobRef.Create(ctx, jobMetaData); err != nil {
		e := fmt.Sprintf("Failed to Create bigQuery table")
		s.Logger.WithError(err).Error(e)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": e,
		})
		return
	}

	jobCompetencyMetaData := &bigquery.TableMetadata{
		Schema: jobCompetencySchema,
	}
	jobCompetencyRef := s.BigQueryClient.Dataset("hcm").Table("job_competency")
	if err := jobCompetencyRef.Create(ctx, jobCompetencyMetaData); err != nil {
		e := fmt.Sprintf("Failed to Create bigQuery table")
		s.Logger.WithError(err).Error(e)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": e,
		})
		return
	}

	jobEducationSpecializationMetaData := &bigquery.TableMetadata{
		Schema: jobEducationSpecializationSchema,
	}
	jobEducationSpecializationRef := s.BigQueryClient.Dataset("hcm").Table("job_education_specialization")
	if err := jobEducationSpecializationRef.Create(ctx, jobEducationSpecializationMetaData); err != nil {
		e := fmt.Sprintf("Failed to Create bigQuery table")
		s.Logger.WithError(err).Error(e)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": e,
		})
		return
	}

	jobLanguageMetaData := &bigquery.TableMetadata{
		Schema: jobLanguageSchema,
	}
	jobLanguageRef := s.BigQueryClient.Dataset("hcm").Table("job_language")
	if err := jobLanguageRef.Create(ctx, jobLanguageMetaData); err != nil {
		e := fmt.Sprintf("Failed to Create bigQuery table")
		s.Logger.WithError(err).Error(e)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": e,
		})
		return
	}

	jobNationalityMetaData := &bigquery.TableMetadata{
		Schema: jobNationalitySchema,
	}
	jobNationalityRef := s.BigQueryClient.Dataset("hcm").Table("job_nationality")
	if err := jobNationalityRef.Create(ctx, jobNationalityMetaData); err != nil {
		e := fmt.Sprintf("Failed to Create bigQuery table")
		s.Logger.WithError(err).Error(e)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": e,
		})
		return
	}

	degreeLevelMetaData := &bigquery.TableMetadata{
		Schema: degreeLevelSchema,
	}
	degreeLevelRef := s.BigQueryClient.Dataset("hcm").Table("degree_level")
	if err := degreeLevelRef.Create(ctx, degreeLevelMetaData); err != nil {
		e := fmt.Sprintf("Failed to Create bigQuery table")
		s.Logger.WithError(err).Error(e)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": e,
		})
		return
	}

	/*
			 SELECT
		       * EXCEPT(is_typed)
		     FROM
		     profile.INFORMATION_SCHEMA.TABLES
	*/

	common.ResponseData(c, true)
}
