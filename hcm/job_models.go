package hcm

import "gitlab.com/target-smart-data-ai-search/pg-log-extractor/common"

// Job represents a row from 'job'.
type Job struct {
	ID                      int64   `bigquery:"id"`                         // id
	JobTitle                string  `bigquery:"job_title"`                  // job_title
	JobDescription          string  `bigquery:"job_description"`            // job_description
	Skills                  string  `bigquery:"skills"`                     // skills
	FunctionalArea          string  `bigquery:"functional_area"`            // functional_area
	TotalPosition           int64   `bigquery:"total_position"`             // total_position
	Shift                   string  `bigquery:"shift"`                      // shift
	Gender                  string  `bigquery:"gender"`                     // gender
	SpecificDegreeTitle     string  `bigquery:"specific_degree_title"`      // specific_degree_title
	MajorSubject            string  `bigquery:"major_subject"`              // major_subject
	MinorSubject            string  `bigquery:"minor_subject"`              // minor_subject
	Responsibilities        string  `bigquery:"responsibilities"`           // responsibilities
	MinExperience           int64   `bigquery:"min_experience"`             // min_experience
	MaxExperience           int64   `bigquery:"max_experience"`             // max_experience
	ExperienceDetail        string  `bigquery:"experience_detail"`          // experience_detail
	IsHideSalary            bool    `bigquery:"is_hide_salary"`             // is_hide_salary
	Gpa                     float64 `bigquery:"gpa"`                        // gpa
	OutOf                   float64 `bigquery:"out_of"`                     // out_of
	Grade                   string  `bigquery:"grade"`                      // grade
	Percentage              float64 `bigquery:"percentage"`                 // percentage
	PrimaryLocation         string  `bigquery:"primary_location"`           // primary_location
	OtherLocation           string  `bigquery:"other_location"`             // other_location
	PostingDate             string  `bigquery:"posting_date"`               // posting_date
	AppliedBefore           string  `bigquery:"applied_before"`             // applied_before
	Activities              string  `bigquery:"activities"`                 // activities
	Challanges              string  `bigquery:"challanges"`                 // challanges
	CompanyID               int64   `bigquery:"company_id"`                 // company_id
	PostingImmediatly       bool    `bigquery:"posting_immediatly"`         // posting_immediatly
	MinAge                  int64   `bigquery:"min_age"`                    // min_age
	MaxAge                  int64   `bigquery:"max_age"`                    // max_age
	IsPublish               bool    `bigquery:"is_publish"`                 // is_publish
	IsActive                bool    `bigquery:"is_active"`                  // is_active
	CreatedBy               int64   `bigquery:"created_by"`                 // created_by
	CreatedDate             string  `bigquery:"created_date"`               // created_date
	ChangedDate             string  `bigquery:"changed_date"`               // changed_date
	DeletedDate             string  `bigquery:"deleted_date"`               // deleted_date
	VacancyCode             int64   `bigquery:"vacancy_code"`               // vacancy_code
	VacancyReferenseNumber  int64   `bigquery:"vacancy_referense_number"`   // vacancy_referense_number
	MaleRequiredCount       int64   `bigquery:"male_required_count"`        // male_required_count
	FemaleRequiredCount     int64   `bigquery:"female_required_count"`      // female_required_count
	BothGenderRequiredCount int64   `bigquery:"both_gender_required_count"` // both_gender_required_count
	UniversityID            int64   `bigquery:"university_id"`              // university_id
	DegreeID                int64   `bigquery:"degree_id"`                  // degree_id
	SpecializationID        int64   `bigquery:"specialization_id"`          // specialization_id
	IsPrivate               bool    `bigquery:"is_private"`                 // is_private
	IsPublicmatchOnly       bool    `bigquery:"is_publicmatch_only"`        // is_publicmatch_only
	IndustryID              int64   `bigquery:"industry_id"`                // industry_id
	EmploymentTypeID        int64   `bigquery:"employment_type_id"`         // employment_type_id
	CareerLevelID           int64   `bigquery:"career_level_id"`            // career_level_id
	MinSalary               float64 `bigquery:"min_salary"`                 // min_salary
	MaxSalary               float64 `bigquery:"max_salary"`                 // max_salary
	CurrencySalary          string  `bigquery:"currency_salary"`            // currency_salary
	CountryID               int64   `bigquery:"country_id"`                 // country_id
	StateID                 int64   `bigquery:"state_id"`                   // state_id
	CityID                  int64   `bigquery:"city_id"`                    // city_id
	JobNo                   string  `bigquery:"job_no"`                     // job_no
	JobTypeID               int64   `bigquery:"job_type_id"`                // job_type_id
	CreatorSubject          string  `bigquery:"creator_subject"`            // creator_subject
	CompanyName             string  `bigquery:"company_name"`               // company_name
	AccountId               int64   `bigquery:"account_id"`
	ContactIds              string  `bigquery:"contact_ids"`
	OrgID                   int64   `bigquery:"org_id"`
}

func ParseValueToJob(job map[string]interface{}) interface{} {

	var item Job

	common.GetMapInt64Value(job, "id", &item.ID)
	common.GetMapStringValue(job, "job_title", &item.JobTitle)
	common.GetMapStringValue(job, "job_description", &item.JobDescription)
	common.GetMapStringValue(job, "skills", &item.Skills)
	common.GetMapStringValue(job, "functional_area", &item.FunctionalArea)
	common.GetMapInt64Value(job, "total_position", &item.TotalPosition)
	common.GetMapStringValue(job, "shift", &item.Shift)
	common.GetMapStringValue(job, "gender", &item.Gender)
	common.GetMapStringValue(job, "specific_degree_title", &item.SpecificDegreeTitle)
	common.GetMapStringValue(job, "major_subject", &item.MajorSubject)
	common.GetMapStringValue(job, "minor_subject", &item.MinorSubject)
	common.GetMapStringValue(job, "responsibilities", &item.Responsibilities)
	common.GetMapInt64Value(job, "min_experience", &item.MinExperience)
	common.GetMapInt64Value(job, "max_experience", &item.MaxExperience)
	common.GetMapStringValue(job, "experience_detail", &item.ExperienceDetail)
	common.GetMapBoolValue(job, "is_hide_salary", &item.IsHideSalary)
	common.GetMapFloat64Value(job, "gpa", &item.Gpa)
	common.GetMapFloat64Value(job, "out_of", &item.OutOf)
	common.GetMapStringValue(job, "grade", &item.Grade)
	common.GetMapFloat64Value(job, "percentage", &item.Percentage)
	common.GetMapStringValue(job, "primary_location", &item.PrimaryLocation)
	common.GetMapStringValue(job, "other_location", &item.OtherLocation)
	common.GetMapStringValue(job, "posting_date", &item.PostingDate)
	common.GetMapStringValue(job, "applied_before", &item.AppliedBefore)
	common.GetMapStringValue(job, "activities", &item.Activities)
	common.GetMapStringValue(job, "challanges", &item.Challanges)
	common.GetMapInt64Value(job, "company_id", &item.CompanyID)
	common.GetMapBoolValue(job, "posting_immediatly", &item.PostingImmediatly)
	common.GetMapInt64Value(job, "min_age", &item.MinAge)
	common.GetMapInt64Value(job, "max_age", &item.MaxAge)
	common.GetMapBoolValue(job, "is_publish", &item.IsPublish)
	common.GetMapBoolValue(job, "is_active", &item.IsActive)
	common.GetMapInt64Value(job, "created_by", &item.CreatedBy)
	common.GetMapStringValue(job, "created_date", &item.CreatedDate)
	common.GetMapStringValue(job, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(job, "deleted_date", &item.DeletedDate)
	common.GetMapInt64Value(job, "vacancy_code", &item.VacancyCode)
	common.GetMapInt64Value(job, "vacancy_referense_number", &item.VacancyReferenseNumber)
	common.GetMapInt64Value(job, "male_required_count", &item.MaleRequiredCount)
	common.GetMapInt64Value(job, "female_required_count", &item.FemaleRequiredCount)
	common.GetMapInt64Value(job, "both_gender_required_count", &item.BothGenderRequiredCount)
	common.GetMapInt64Value(job, "university_id", &item.UniversityID)
	common.GetMapInt64Value(job, "degree_id", &item.DegreeID)
	common.GetMapInt64Value(job, "specialization_id", &item.SpecializationID)
	common.GetMapBoolValue(job, "is_private", &item.IsPrivate)
	common.GetMapBoolValue(job, "is_publicmatch_only", &item.IsPublicmatchOnly)
	common.GetMapInt64Value(job, "industry_id", &item.IndustryID)
	common.GetMapInt64Value(job, "employment_type_id", &item.EmploymentTypeID)
	common.GetMapInt64Value(job, "career_level_id", &item.CareerLevelID)
	common.GetMapFloat64Value(job, "min_salary", &item.MinSalary)
	common.GetMapFloat64Value(job, "max_salary", &item.MaxSalary)
	common.GetMapStringValue(job, "currency_salary", &item.CurrencySalary)
	common.GetMapInt64Value(job, "country_id", &item.CountryID)
	common.GetMapInt64Value(job, "state_id", &item.StateID)
	common.GetMapInt64Value(job, "city_id", &item.CityID)
	common.GetMapStringValue(job, "job_no", &item.JobNo)
	common.GetMapInt64Value(job, "job_type_id", &item.JobTypeID)
	common.GetMapStringValue(job, "creator_subject", &item.CreatorSubject)
	common.GetMapStringValue(job, "company_name", &item.CompanyName)
	common.GetMapInt64Value(job, "account_id", &item.AccountId)
	common.GetMapStringValue(job, "contact_ids", &item.ContactIds)
	common.GetMapInt64Value(job, "org_id", &item.OrgID)

	return item
}

// JobCompetency represents a row from 'job_competency'.
type JobCompetency struct {
	ID           int64  `bigquery:"id"`            // id
	JobID        int64  `bigquery:"job_id"`        // job_id
	CompetencyID int64  `bigquery:"competency_id"` // competency_id
	LevelID      int64  `bigquery:"level_id"`      // level_id
	IsGained     bool   `bigquery:"is_gained"`     // is_gained
	CreatedDate  string `bigquery:"created_date"`  // created_date
	ChangedDate  string `bigquery:"changed_date"`  // changed_date
	DeletedDate  string `bigquery:"deleted_date"`  // deleted_date
}

func ParseValueToJobCompetency(job map[string]interface{}) interface{} {

	var item JobCompetency

	common.GetMapInt64Value(job, "id", &item.ID)
	common.GetMapInt64Value(job, "job_id", &item.JobID)
	common.GetMapInt64Value(job, "level_id", &item.LevelID)
	common.GetMapInt64Value(job, "competency_id", &item.CompetencyID)
	common.GetMapBoolValue(job, "is_gained", &item.IsGained)
	common.GetMapStringValue(job, "created_date", &item.CreatedDate)
	common.GetMapStringValue(job, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(job, "deleted_date", &item.DeletedDate)

	return item
}

// JobEducationSpecialization represents a row from 'job_education_specialization'.
type JobEducationSpecialization struct {
	ID               int64  `json:"id"`                // id
	JobID            int64  `json:"job_id"`            // job_id
	SpecializationID int64  `json:"specialization_id"` // specialization_id
	CreatedDate      string `json:"created_date"`      // created_date
	ChangedDate      string `json:"changed_date"`      // changed_date
	DeletedDate      string `json:"deleted_date"`      // deleted_date
}

func ParseValueToJobEducationSpecialization(job map[string]interface{}) interface{} {

	var item JobEducationSpecialization

	common.GetMapInt64Value(job, "id", &item.ID)
	common.GetMapInt64Value(job, "job_id", &item.JobID)
	common.GetMapInt64Value(job, "specialization_id", &item.SpecializationID)
	common.GetMapStringValue(job, "created_date", &item.CreatedDate)
	common.GetMapStringValue(job, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(job, "deleted_date", &item.DeletedDate)

	return item
}

// JobLanguage represents a row from 'job_language'.
type JobLanguage struct {
	ID          int64  `json:"id"`           // id
	JobID       int64  `json:"job_id"`       // job_id
	LanguageID  int64  `json:"language_id"`  // language_id
	CreatedDate string `json:"created_date"` // created_date
	ChangedDate string `json:"changed_date"` // changed_date
	DeletedDate string `json:"deleted_date"` // deleted_date
}

func ParseValueToJobLanguage(job map[string]interface{}) interface{} {

	var item JobLanguage

	common.GetMapInt64Value(job, "id", &item.ID)
	common.GetMapInt64Value(job, "job_id", &item.JobID)
	common.GetMapInt64Value(job, "language_id", &item.LanguageID)
	common.GetMapStringValue(job, "created_date", &item.CreatedDate)
	common.GetMapStringValue(job, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(job, "deleted_date", &item.DeletedDate)

	return item
}

// JobNationality represents a row from 'job_nationality'.
type JobNationality struct {
	ID          int64  `json:"id"`           // id
	JobID       int64  `json:"job_id"`       // job_id
	CountryID   int64  `json:"country_id""`  // country_id
	CreatedDate string `json:"created_date"` // created_date
	ChangedDate string `json:"changed_date"` // changed_date
	DeletedDate string `json:"deleted_date"` // deleted_date
}

func ParseValueToJobNationality(job map[string]interface{}) interface{} {

	var item JobNationality

	common.GetMapInt64Value(job, "id", &item.ID)
	common.GetMapInt64Value(job, "job_id", &item.JobID)
	common.GetMapInt64Value(job, "country_id", &item.CountryID)
	common.GetMapStringValue(job, "created_date", &item.CreatedDate)
	common.GetMapStringValue(job, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(job, "deleted_date", &item.DeletedDate)

	return item
}

// DegreeLevel represents a row from 'degree_level'.
type DegreeLevel struct {
	ID             int64  `bigquery:"id"`               // id
	DegreeName     string `bigquery:"degree_name"`      // degree_name
	Description    string `bigquery:"description"`      // description
	CreatedDate    string `bigquery:"created_date"`     // created_date
	ChangedDate    string `bigquery:"changed_date"`     // changed_date
	DeletedDate    string `bigquery:"deleted_date"`     // deleted_date
	YearsOfStudies int64  `bigquery:"years_of_studies"` // years_of_studies
}

func ParseValueToDegreeLevel(degree map[string]interface{}) interface{} {

	var item DegreeLevel

	common.GetMapInt64Value(degree, "id", &item.ID)
	common.GetMapStringValue(degree, "degree_name", &item.DegreeName)
	common.GetMapStringValue(degree, "description", &item.Description)
	common.GetMapStringValue(degree, "created_date", &item.CreatedDate)
	common.GetMapStringValue(degree, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(degree, "deleted_date", &item.DeletedDate)
	common.GetMapInt64Value(degree, "years_of_studies", &item.YearsOfStudies)

	return item
}
