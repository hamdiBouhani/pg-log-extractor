package profile

import (
	"context"

	"cloud.google.com/go/bigquery"
	"github.com/gin-gonic/gin"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/common"
)

var (
	// DatasetID represents a database from 'user_profile'.
	DatasetID = "profile"
)

func (s *Server) InitUserProfileBigQueryShema(c *gin.Context) {
	//user_profile
	userProfileSchema := bigquery.Schema{
		{Name: "user_id", Type: bigquery.IntegerFieldType, Required: true},
		{Name: "title", Type: bigquery.StringFieldType},
		{Name: "full_name", Type: bigquery.StringFieldType},
		{Name: "about_me", Type: bigquery.StringFieldType},
		{Name: "nationality", Type: bigquery.StringFieldType},
		{Name: "date_of_birth", Type: bigquery.StringFieldType},
		{Name: "gender", Type: bigquery.StringFieldType},
		{Name: "resident_card_no", Type: bigquery.StringFieldType},
		{Name: "phone_mobile", Type: bigquery.StringFieldType},
		{Name: "phone_home", Type: bigquery.StringFieldType},
		{Name: "subject", Type: bigquery.StringFieldType},
		{Name: "email", Type: bigquery.StringFieldType},
		{Name: "last_login_date", Type: bigquery.StringFieldType},
		{Name: "is_searchable", Type: bigquery.BooleanFieldType},
		{Name: "is_active", Type: bigquery.BooleanFieldType},
		{Name: "is_locked", Type: bigquery.BooleanFieldType},
		{Name: "bkg_img", Type: bigquery.StringFieldType},
		{Name: "photo", Type: bigquery.IntegerFieldType},
		{Name: "country_id", Type: bigquery.IntegerFieldType},
		{Name: "city_id", Type: bigquery.IntegerFieldType},
		{Name: "state_id", Type: bigquery.IntegerFieldType},
		{Name: "created_date", Type: bigquery.StringFieldType},
		{Name: "changed_date", Type: bigquery.StringFieldType},
		{Name: "deleted_date", Type: bigquery.StringFieldType},
		{Name: "groups", Type: bigquery.StringFieldType},
		{Name: "display_name", Type: bigquery.StringFieldType},
		{Name: "disabilities", Type: bigquery.StringFieldType},
		{Name: "marital_status", Type: bigquery.StringFieldType},
		{Name: "driving_license_no", Type: bigquery.StringFieldType},
		{Name: "driving_license_date", Type: bigquery.StringFieldType},
		{Name: "driving_license_type", Type: bigquery.StringFieldType},
		{Name: "passport_no", Type: bigquery.StringFieldType},
		{Name: "tax_no", Type: bigquery.StringFieldType},
		{Name: "vat_no", Type: bigquery.StringFieldType},
		{Name: "street_address", Type: bigquery.StringFieldType},
		{Name: "postal_code", Type: bigquery.StringFieldType},
		{Name: "phone_work", Type: bigquery.StringFieldType},
		{Name: "metadata", Type: bigquery.StringFieldType},
		{Name: "is_contact_info_visible", Type: bigquery.BooleanFieldType},
		{Name: "primary_language", Type: bigquery.IntegerFieldType},
		{Name: "device_token", Type: bigquery.StringFieldType},
		{Name: "access_token_count", Type: bigquery.IntegerFieldType},
		{Name: "picture_url", Type: bigquery.StringFieldType},
		{Name: "dependents_count", Type: bigquery.IntegerFieldType},
	}
	//user_profile_career_aspiration
	userProfileCareerAspirationSchema := bigquery.Schema{
		{Name: "id", Type: bigquery.IntegerFieldType, Required: true},
		{Name: "user_id", Type: bigquery.IntegerFieldType},
		{Name: "industry_id", Type: bigquery.IntegerFieldType},
		{Name: "employment_type_id", Type: bigquery.IntegerFieldType},
		{Name: "contract_type_id", Type: bigquery.IntegerFieldType},
		{Name: "frequency_of_trips", Type: bigquery.StringFieldType},
		{Name: "prefered_country", Type: bigquery.IntegerFieldType},
		{Name: "prefered_city", Type: bigquery.IntegerFieldType},
		{Name: "salary_range_min", Type: bigquery.FloatFieldType},
		{Name: "salary_range_max", Type: bigquery.FloatFieldType},
		{Name: "created_date", Type: bigquery.StringFieldType},
		{Name: "changed_date", Type: bigquery.StringFieldType},
		{Name: "deleted_date", Type: bigquery.StringFieldType},
		{Name: "prefered_state", Type: bigquery.IntegerFieldType},
		{Name: "salary_currency", Type: bigquery.IntegerFieldType},
	}
	//experience
	experienceSchema := bigquery.Schema{
		{Name: "id", Type: bigquery.IntegerFieldType, Required: true},
		{Name: "user_id", Type: bigquery.IntegerFieldType},
		{Name: "job_title", Type: bigquery.StringFieldType},
		{Name: "job_description", Type: bigquery.StringFieldType},
		{Name: "company_name", Type: bigquery.StringFieldType},
		{Name: "company_logo", Type: bigquery.StringFieldType},
		{Name: "date_from", Type: bigquery.StringFieldType},
		{Name: "date_to", Type: bigquery.StringFieldType},
		{Name: "is_current", Type: bigquery.BooleanFieldType},
		{Name: "country_id", Type: bigquery.IntegerFieldType},
		{Name: "city_id", Type: bigquery.IntegerFieldType},
		{Name: "state_id", Type: bigquery.IntegerFieldType},
		{Name: "created_date", Type: bigquery.StringFieldType},
		{Name: "changed_date", Type: bigquery.StringFieldType},
		{Name: "deleted_date", Type: bigquery.StringFieldType},
		{Name: "experience", Type: bigquery.FloatFieldType},
		{Name: "designation_name", Type: bigquery.StringFieldType},
		{Name: "department", Type: bigquery.StringFieldType},
		{Name: "responsibilities", Type: bigquery.StringFieldType},
		{Name: "email", Type: bigquery.StringFieldType},
		{Name: "phone_office", Type: bigquery.StringFieldType},
		{Name: "fax_number", Type: bigquery.StringFieldType},
		{Name: "is_terminated", Type: bigquery.BooleanFieldType},
		{Name: "current_employee", Type: bigquery.BooleanFieldType},
		{Name: "additional_notes", Type: bigquery.StringFieldType},
		{Name: "taxonomy_position_id", Type: bigquery.IntegerFieldType},
		{Name: "attachments", Type: bigquery.StringFieldType},
		{Name: "created_by", Type: bigquery.IntegerFieldType},
	}
	//user_profile_competency
	userCompetencySchema := bigquery.Schema{
		{Name: "id", Type: bigquery.IntegerFieldType, Required: true},
		{Name: "user_id", Type: bigquery.IntegerFieldType},
		{Name: "competency_id", Type: bigquery.IntegerFieldType},
		{Name: "level_id", Type: bigquery.IntegerFieldType},
		{Name: "created_date", Type: bigquery.StringFieldType},
		{Name: "changed_date", Type: bigquery.StringFieldType},
		{Name: "deleted_date", Type: bigquery.StringFieldType},
		{Name: "is_gained", Type: bigquery.BooleanFieldType},
		{Name: "certificate_id", Type: bigquery.IntegerFieldType},
	}
	//user_profile_education_specialization
	userProfileEducationSpecializationSchema := bigquery.Schema{
		{Name: "id", Type: bigquery.IntegerFieldType, Required: true},
		{Name: "minor_subject", Type: bigquery.StringFieldType},
		{Name: "major_subject", Type: bigquery.StringFieldType},
		{Name: "description", Type: bigquery.StringFieldType},
		{Name: "created_date", Type: bigquery.StringFieldType},
		{Name: "changed_date", Type: bigquery.StringFieldType},
		{Name: "deleted_date", Type: bigquery.StringFieldType},
	}
	//language_bridge
	languageBridgeSchema := bigquery.Schema{
		{Name: "id", Type: bigquery.IntegerFieldType, Required: true},
		{Name: "user_id", Type: bigquery.IntegerFieldType},
		{Name: "created_date", Type: bigquery.StringFieldType},
		{Name: "changed_date", Type: bigquery.StringFieldType},
		{Name: "deleted_date", Type: bigquery.StringFieldType},
	}
	//user_profile_education
	userEducationSchema := bigquery.Schema{
		{Name: "id", Type: bigquery.IntegerFieldType, Required: true},
		{Name: "user_id", Type: bigquery.IntegerFieldType},
		{Name: "degree_level_id", Type: bigquery.IntegerFieldType},
		{Name: "university_name", Type: bigquery.StringFieldType},
		{Name: "start_date", Type: bigquery.StringFieldType},
		{Name: "end_date", Type: bigquery.StringFieldType},
		{Name: "country_id", Type: bigquery.IntegerFieldType},
		{Name: "city_id", Type: bigquery.IntegerFieldType},
		{Name: "state_id", Type: bigquery.IntegerFieldType},
		{Name: "created_date", Type: bigquery.StringFieldType},
		{Name: "changed_date", Type: bigquery.StringFieldType},
		{Name: "deleted_date", Type: bigquery.StringFieldType},
		{Name: "gpa", Type: bigquery.FloatFieldType},
		{Name: "percentage", Type: bigquery.FloatFieldType},
		{Name: "specialization_id", Type: bigquery.IntegerFieldType},
		{Name: "passing_year", Type: bigquery.StringFieldType},
		{Name: "out_of", Type: bigquery.FloatFieldType},
		{Name: "grade", Type: bigquery.StringFieldType},
		{Name: "attachments", Type: bigquery.StringFieldType},
		{Name: "final_project", Type: bigquery.StringFieldType},
		{Name: "description", Type: bigquery.StringFieldType},
		{Name: "created_by", Type: bigquery.IntegerFieldType},
		{Name: "doc_id", Type: bigquery.IntegerFieldType},
	}
	//degree_level
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
	if err := s.BigQueryClient.Dataset(DatasetID).Create(ctx, meta); err != nil {
		s.Logger.Errorf("Failed to Create bigQuery dataset : %v", err)
		return
	}

	userProfileMetaData := &bigquery.TableMetadata{
		Schema: userProfileSchema,
	}
	userProfileRef := s.BigQueryClient.Dataset(DatasetID).Table("user_profile")
	if err := userProfileRef.Create(ctx, userProfileMetaData); err != nil {
		s.Logger.Errorf("Failed to Create bigQuery table : %v", err)
		return
	}

	userProfileCareerAspirationMetaData := &bigquery.TableMetadata{
		Schema: userProfileCareerAspirationSchema,
	}
	userProfileCareerAspirationRef := s.BigQueryClient.Dataset(DatasetID).Table("user_profile_career_aspiration")
	if err := userProfileCareerAspirationRef.Create(ctx, userProfileCareerAspirationMetaData); err != nil {
		s.Logger.Errorf("Failed to Create bigQuery table : %v", err)
		return
	}

	experienceMetaData := &bigquery.TableMetadata{
		Schema: experienceSchema,
	}
	experienceRef := s.BigQueryClient.Dataset(DatasetID).Table("experience")
	if err := experienceRef.Create(ctx, experienceMetaData); err != nil {
		s.Logger.Errorf("Failed to Create bigQuery table : %v", err)
		return
	}

	userCompetencyMetaData := &bigquery.TableMetadata{
		Schema: userCompetencySchema,
	}
	userCompetencyRef := s.BigQueryClient.Dataset(DatasetID).Table("user_competency")
	if err := userCompetencyRef.Create(ctx, userCompetencyMetaData); err != nil {
		s.Logger.Errorf("Failed to Create bigQuery table : %v", err)
		return
	}

	userProfileEducationSpecializationMetaData := &bigquery.TableMetadata{
		Schema: userProfileEducationSpecializationSchema,
	}
	userProfileEducationSpecializationRef := s.BigQueryClient.Dataset(DatasetID).Table("education_specialization")
	if err := userProfileEducationSpecializationRef.Create(ctx, userProfileEducationSpecializationMetaData); err != nil {
		s.Logger.Errorf("Failed to Create bigQuery table : %v", err)
		return
	}

	languageBridgeMetaData := &bigquery.TableMetadata{
		Schema: languageBridgeSchema,
	}
	languageBridgeRef := s.BigQueryClient.Dataset(DatasetID).Table("language_bridge")
	if err := languageBridgeRef.Create(ctx, languageBridgeMetaData); err != nil {
		s.Logger.Errorf("Failed to Create bigQuery table : %v", err)
		return
	}

	userEducationMetaData := &bigquery.TableMetadata{
		Schema: userEducationSchema,
	}
	userEducationRef := s.BigQueryClient.Dataset(DatasetID).Table("user_education")
	if err := userEducationRef.Create(ctx, userEducationMetaData); err != nil {
		s.Logger.Errorf("Failed to Create bigQuery table : %v", err)
		return
	}

	degreeLevelMetaData := &bigquery.TableMetadata{
		Schema: degreeLevelSchema,
	}
	degreeLevelRef := s.BigQueryClient.Dataset(DatasetID).Table("degree_level")
	if err := degreeLevelRef.Create(ctx, degreeLevelMetaData); err != nil {
		s.Logger.Errorf("Failed to Create bigQuery table : %v", err)
		return
	}

	/*
			 SELECT
		       * EXCEPT(is_typed)
		     FROM
		     profile.INFORMATION_SCHEMA.TABLES
	*/

	common.ResponseData(c, nil)
}
