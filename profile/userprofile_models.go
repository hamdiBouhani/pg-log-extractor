package profile

import "gitlab.com/target-smart-data-ai-search/pg-log-extractor/common"

// UserProfile represents a row from 'user_profile'.
type UserProfile struct {
	UserID               int64  `bigquery:"user_id"`                 // user_id
	Title                string `bigquery:"title"`                   // title
	FullName             string `bigquery:"full_name"`               // full_name
	AboutMe              string `bigquery:"about_me"`                // about_me
	Nationality          string `bigquery:"nationality"`             // nationality
	DateOfBirth          string `bigquery:"date_of_birth"`           // date_of_birth
	Gender               string `bigquery:"gender"`                  // gender
	ResidentCardNo       string `bigquery:"resident_card_no"`        // resident_card_no
	PhoneMobile          string `bigquery:"phone_mobile"`            // phone_mobile
	PhoneHome            string `bigquery:"phone_home"`              // phone_home
	Subject              string `bigquery:"subject"`                 // subject
	Email                string `bigquery:"email"`                   // email
	LastLoginDate        string `bigquery:"last_login_date"`         // last_login_date
	IsSearchable         bool   `bigquery:"is_searchable"`           // is_searchable
	IsActive             bool   `bigquery:"is_active"`               // is_active
	IsLocked             bool   `bigquery:"is_locked"`               // is_locked
	BkgImg               string `bigquery:"bkg_img"`                 // bkg_img
	Photo                int64  `bigquery:"photo"`                   // photo
	CountryID            int64  `bigquery:"country_id"`              // country_id
	CityID               int64  `bigquery:"city_id"`                 // city_id
	StateID              int64  `bigquery:"state_id"`                // state_id
	CreatedDate          string `bigquery:"created_date"`            // created_date
	ChangedDate          string `bigquery:"changed_date"`            // changed_date
	DeletedDate          string `bigquery:"deleted_date"`            // deleted_date
	Groups               string `bigquery:"groups"`                  // groups
	DisplayName          string `bigquery:"display_name"`            // display_name
	Disabilities         string `bigquery:"disabilities"`            // disabilities
	MaritalStatus        string `bigquery:"marital_status"`          // marital_status
	DrivingLicenseNo     string `bigquery:"driving_license_no"`      // driving_license_no
	DrivingLicenseDate   string `bigquery:"driving_license_date"`    // driving_license_date
	DrivingLicenseType   string `bigquery:"driving_license_type"`    // driving_license_type
	PassportNo           string `bigquery:"passport_no"`             // passport_no
	TaxNo                string `bigquery:"tax_no"`                  // tax_no
	VatNo                string `bigquery:"vat_no"`                  // vat_no
	StreetAddress        string `bigquery:"street_address"`          // street_address
	PostalCode           string `bigquery:"postal_code"`             // postal_code
	PhoneWork            string `bigquery:"phone_work"`              // phone_work
	Metadata             string `bigquery:"metadata"`                // metadata
	IsContactInfoVisible bool   `bigquery:"is_contact_info_visible"` // is_contact_info_visible
	PrimaryLanguage      int64  `bigquery:"primary_language"`        // primary_language
	DeviceToken          string `bigquery:"device_token"`            // device_token
	AccessTokenCount     int64  `bigquery:"access_token_count"`      // access_token_count
	PictureURL           string `bigquery:"picture_url"`             // picture_url
	DependentsCount      int64  `bigquery:"dependents_count"`        // dependents_count
}

func ParseValueToUserProfile(user map[string]interface{}) interface{} {

	var item UserProfile

	common.GetMapInt64Value(user, "user_id", &item.UserID)
	common.GetMapStringValue(user, "title", &item.Title)
	common.GetMapStringValue(user, "full_name", &item.FullName)
	common.GetMapStringValue(user, "about_me", &item.AboutMe)
	common.GetMapStringValue(user, "nationality", &item.Nationality)
	common.GetMapStringValue(user, "date_of_birth", &item.DateOfBirth)
	common.GetMapStringValue(user, "gender", &item.Gender)
	common.GetMapStringValue(user, "resident_card_no", &item.ResidentCardNo)
	common.GetMapStringValue(user, "phone_mobile", &item.PhoneMobile)
	common.GetMapStringValue(user, "phone_home", &item.PhoneHome)
	common.GetMapStringValue(user, "subject", &item.Subject)
	common.GetMapStringValue(user, "email", &item.Email)
	common.GetMapStringValue(user, "last_login_date", &item.LastLoginDate)
	common.GetMapBoolValue(user, "is_searchable", &item.IsSearchable)
	common.GetMapBoolValue(user, "is_active", &item.IsActive)
	common.GetMapBoolValue(user, "is_locked", &item.IsLocked)
	common.GetMapStringValue(user, "bkg_img", &item.BkgImg)
	common.GetMapInt64Value(user, "photo", &item.Photo)
	common.GetMapInt64Value(user, "country_id", &item.CountryID)
	common.GetMapInt64Value(user, "city_id", &item.CityID)
	common.GetMapInt64Value(user, "state_id", &item.StateID)
	common.GetMapStringValue(user, "created_date", &item.CreatedDate)
	common.GetMapStringValue(user, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(user, "deleted_date", &item.DeletedDate)
	common.GetMapStringValue(user, "groups", &item.Groups)
	common.GetMapStringValue(user, "display_name", &item.DisplayName)
	common.GetMapStringValue(user, "disabilities", &item.Disabilities)
	common.GetMapStringValue(user, "marital_status", &item.MaritalStatus)
	common.GetMapStringValue(user, "driving_license_no", &item.DrivingLicenseNo)
	common.GetMapStringValue(user, "bkg_idriving_license_date", &item.DrivingLicenseDate)
	common.GetMapStringValue(user, "driving_license_type", &item.DrivingLicenseType)
	common.GetMapStringValue(user, "passport_no", &item.PassportNo)
	common.GetMapStringValue(user, "tax_no", &item.TaxNo)
	common.GetMapStringValue(user, "vat_no", &item.VatNo)
	common.GetMapStringValue(user, "street_address", &item.StreetAddress)
	common.GetMapStringValue(user, "postal_code", &item.PostalCode)
	common.GetMapStringValue(user, "phone_work", &item.PhoneWork)
	common.GetMapStringValue(user, "metadata", &item.Metadata)
	common.GetMapBoolValue(user, "is_contact_info_visible", &item.IsContactInfoVisible)
	common.GetMapInt64Value(user, "primary_language", &item.PrimaryLanguage)
	common.GetMapStringValue(user, "device_token", &item.DeviceToken)
	common.GetMapInt64Value(user, "access_token_count", &item.AccessTokenCount)
	common.GetMapStringValue(user, "picture_url", &item.PictureURL)
	common.GetMapInt64Value(user, "dependents_count", &item.DependentsCount)

	return item
}

type UserCompetency struct {
	ID            int64  `bigquery:"id"`
	UserID        int64  `bigquery:"user_id"`
	CompetencyID  int64  `bigquery:"competency_id"`
	LevelID       int64  `bigquery:"level_id"`
	CreatedBy     int64  `bigquery:"created_by"`
	CreatedDate   string `bigquery:"created_date"`
	ChangedDate   string `bigquery:"changed_date"`
	DeletedDate   string `bigquery:"deleted_date"`
	IsGained      bool   `bigquery:"is_gained"`
	CertificateID int64  `bigquery:"certificate_id"`
}

func ParseValueToUserCompetency(competency map[string]interface{}) interface{} {

	var item UserCompetency

	common.GetMapInt64Value(competency, "id", &item.ID)
	common.GetMapInt64Value(competency, "user_id", &item.UserID)
	common.GetMapInt64Value(competency, "competency_id", &item.CompetencyID)
	common.GetMapInt64Value(competency, "level_id", &item.LevelID)
	common.GetMapInt64Value(competency, "created_by", &item.CreatedBy)
	common.GetMapStringValue(competency, "created_date", &item.CreatedDate)
	common.GetMapStringValue(competency, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(competency, "deleted_date", &item.DeletedDate)
	common.GetMapBoolValue(competency, "is_gained", &item.IsGained)
	common.GetMapInt64Value(competency, "certificate_id", &item.CertificateID)

	return item
}

// UserCareerAspiration represents a row from 'user_career_aspiration'.
type UserCareerAspiration struct {
	ID               int64   `bigquery:"id"`                  // id
	UserID           int64   `bigquery:"user_id"`             // user_id
	IndustryID       int64   `bigquery:"industry_id"`         // industry_id
	EmploymentTypeID int64   `bigquery:"employment_type_id" ` // employment_type_id
	ContractTypeID   int64   `bigquery:"contract_type_id"`    // contract_type_id
	FrequencyOfTrips string  `bigquery:"frequency_of_trips"`  // frequency_of_trips
	PreferedCountry  int64   `bigquery:"prefered_country"`    // prefered_country
	PreferedCity     int64   `bigquery:"prefered_city"`       // prefered_city
	SalaryRangeMin   float64 `bigquery:"salary_range_min"`    // salary_range_min
	SalaryRangeMax   float64 `bigquery:"salary_range_max"`    // salary_range_max
	CreatedDate      string  `bigquery:"created_date"`        // created_date
	ChangedDate      string  `bigquery:"changed_date"`        // changed_date
	DeletedDate      string  `bigquery:"deleted_date"`        // deleted_date
	PreferedState    int64   `bigquery:"prefered_state"`      // prefered_state
	SalaryCurrency   int64   `bigquery:"salary_currency"`     // salary_currency
}

func ParseValueToUserCareerAspiration(careerAspiration map[string]interface{}) interface{} {

	var item UserCareerAspiration

	common.GetMapInt64Value(careerAspiration, "id", &item.ID)
	common.GetMapInt64Value(careerAspiration, "user_id", &item.UserID)
	common.GetMapInt64Value(careerAspiration, "industry_id", &item.IndustryID)
	common.GetMapInt64Value(careerAspiration, "employment_type_id", &item.EmploymentTypeID)
	common.GetMapInt64Value(careerAspiration, "contract_type_id", &item.ContractTypeID)
	common.GetMapStringValue(careerAspiration, "frequency_of_trips", &item.FrequencyOfTrips)
	common.GetMapInt64Value(careerAspiration, "prefered_country", &item.PreferedCountry)
	common.GetMapInt64Value(careerAspiration, "prefered_city", &item.PreferedCity)
	common.GetMapFloat64Value(careerAspiration, "salary_range_min", &item.SalaryRangeMin)
	common.GetMapFloat64Value(careerAspiration, "salary_range_max", &item.SalaryRangeMax)
	common.GetMapStringValue(careerAspiration, "created_date", &item.CreatedDate)
	common.GetMapStringValue(careerAspiration, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(careerAspiration, "deleted_date", &item.DeletedDate)
	common.GetMapInt64Value(careerAspiration, "prefered_state", &item.PreferedState)
	common.GetMapInt64Value(careerAspiration, "salary_currency", &item.SalaryCurrency)
	return item
}

// Experience represents a row from 'experience'.
type Experience struct {
	ID                 int64   `bigquery:"id"`                   // id
	UserID             int64   `bigquery:"user_id"`              // user_id
	JobTitle           string  `bigquery:"job_title"`            // job_title
	JobDescription     string  `bigquery:"job_description"`      // job_description
	CompanyName        string  `bigquery:"company_name"`         // company_name
	CompanyLogo        string  `bigquery:"company_logo"`         // company_logo
	DateFrom           string  `bigquery:"date_from"`            // date_from
	DateTo             string  `bigquery:"date_to"`              // date_to
	IsCurrent          bool    `bigquery:"is_current"`           // is_current
	CountryID          int64   `bigquery:"country_id"`           // country_id
	CityID             int64   `bigquery:"city_id"`              // city_id
	StateID            int64   `bigquery:"state_id"`             // state_id
	CreatedDate        string  `bigquery:"created_date"`         // created_date
	ChangedDate        string  `bigquery:"changed_date"`         // changed_date
	DeletedDate        string  `bigquery:"deleted_date"`         // deleted_date
	Experience         float64 `bigquery:"experience"`           // experience
	DesignationName    string  `bigquery:"designation_name"`     // designation_name
	Department         string  `bigquery:"department"`           // department
	Responsibilities   string  `bigquery:"responsibilities"`     // responsibilities
	Email              string  `bigquery:"email"`                // email
	PhoneOffice        string  `bigquery:"phone_office"`         // phone_office
	FaxNumber          string  `bigquery:"fax_number"`           // fax_number
	IsTerminated       bool    `bigquery:"is_terminated"`        // is_terminated
	CurrentEmployee    bool    `bigquery:"current_employee"`     // current_employee
	AdditionalNotes    string  `bigquery:"additional_notes"`     // additional_notes
	TaxonomyPositionID int64   `bigquery:"taxonomy_position_id"` // taxonomy_position_id
	Attachments        string  `bigquery:"attachments"`          // attachments
	CreatedBy          int64   `bigquery:"created_by"`           // created_by
}

func ParseValueToExperience(exp map[string]interface{}) interface{} {

	var item Experience

	common.GetMapInt64Value(exp, "id", &item.ID)
	common.GetMapInt64Value(exp, "user_id", &item.UserID)
	common.GetMapStringValue(exp, "job_title", &item.JobTitle)
	common.GetMapStringValue(exp, "job_description", &item.JobDescription)
	common.GetMapStringValue(exp, "company_name", &item.CompanyName)
	common.GetMapStringValue(exp, "company_logo", &item.CompanyLogo)
	common.GetMapStringValue(exp, "date_from", &item.DateFrom)
	common.GetMapStringValue(exp, "date_to", &item.DateTo)
	common.GetMapBoolValue(exp, "is_current", &item.IsCurrent)
	common.GetMapInt64Value(exp, "country_id", &item.CountryID)
	common.GetMapInt64Value(exp, "city_id", &item.CityID)
	common.GetMapInt64Value(exp, "state_id", &item.StateID)
	common.GetMapStringValue(exp, "created_date", &item.CreatedDate)
	common.GetMapStringValue(exp, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(exp, "deleted_date", &item.DeletedDate)
	common.GetMapFloat64Value(exp, "experience", &item.Experience)
	common.GetMapStringValue(exp, "designation_name", &item.DesignationName)
	common.GetMapStringValue(exp, "department", &item.Department)
	common.GetMapStringValue(exp, "responsibilities", &item.Responsibilities)
	common.GetMapStringValue(exp, "email", &item.Email)
	common.GetMapStringValue(exp, "phone_office", &item.PhoneOffice)
	common.GetMapStringValue(exp, "fax_number", &item.FaxNumber)
	common.GetMapBoolValue(exp, "is_terminated", &item.IsTerminated)
	common.GetMapBoolValue(exp, "current_employee", &item.CurrentEmployee)
	common.GetMapStringValue(exp, "additional_notes", &item.AdditionalNotes)
	common.GetMapInt64Value(exp, "taxonomy_position_id", &item.TaxonomyPositionID)
	common.GetMapStringValue(exp, "attachments", &item.Attachments)
	common.GetMapInt64Value(exp, "created_by", &item.CreatedBy)

	return item
}

// EducationSpecialization represents a row from 'education_specialization'.
type EducationSpecialization struct {
	ID           int64  `bigquery:"id"`            // id
	MinorSubject string `bigquery:"minor_subject"` // minor_subject
	MajorSubject string `bigquery:"major_subject"` // major_subject
	Description  string `bigquery:"description"`   // description
	CreatedDate  string `bigquery:"created_date"`  // created_date
	ChangedDate  string `bigquery:"changed_date"`  // changed_date
	DeletedDate  string `bigquery:"deleted_date"`  // deleted_date
}

func ParseValueToEducationSpecialization(education map[string]interface{}) interface{} {

	var item EducationSpecialization

	common.GetMapInt64Value(education, "id", &item.ID)
	common.GetMapStringValue(education, "minor_subject", &item.MinorSubject)
	common.GetMapStringValue(education, "major_subject", &item.MajorSubject)
	common.GetMapStringValue(education, "description", &item.Description)
	common.GetMapStringValue(education, "created_date", &item.CreatedDate)
	common.GetMapStringValue(education, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(education, "deleted_date", &item.DeletedDate)

	return item
}

// UserLanguage represents a row from 'user_language'.
type UserLanguage struct {
	ID          int64  `bigquery:"id"`           // id
	UserID      int64  `bigquery:"user_id"`      // user_id
	LanguageID  int64  `bigquery:"language_id"`  // language_id
	CreatedDate string `bigquery:"created_date"` // created_date
	ChangedDate string `bigquery:"changed_date"` // changed_date
	DeletedDate string `bigquery:"deleted_date"` // deleted_date
}

func ParseValueToUserLanguage(language map[string]interface{}) interface{} {

	var item UserLanguage

	common.GetMapInt64Value(language, "id", &item.ID)
	common.GetMapInt64Value(language, "user_id", &item.UserID)
	common.GetMapInt64Value(language, "language_id", &item.LanguageID)
	common.GetMapStringValue(language, "created_date", &item.CreatedDate)
	common.GetMapStringValue(language, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(language, "deleted_date", &item.DeletedDate)

	return item
}

// UserEducation represents a row from 'user_education'.
type UserEducation struct {
	ID               int64   `bigquery:"id"`                // id
	UserID           int64   `bigquery:"user_id"`           // user_id
	DegreeLevelID    int64   `bigquery:"degree_level_id"`   // degree_level_id
	DegreeTitle      string  `bigquery:"degree_title"`      // degree_title
	UniversityName   string  `bigquery:"university_name"`   // university_name
	StartDate        string  `bigquery:"start_date"`        // start_date
	EndDate          string  `bigquery:"end_date"`          // end_date
	CountryID        int64   `bigquery:"country_id"`        // country_id
	CityID           int64   `bigquery:"city_id"`           // city_id
	StateID          int64   `bigquery:"state_id"`          // state_id
	CreatedDate      string  `bigquery:"created_date"`      // created_date
	ChangedDate      string  `bigquery:"changed_date"`      // changed_date
	DeletedDate      string  `bigquery:"deleted_date"`      // deleted_date
	Gpa              float64 `bigquery:"gpa"`               // gpa
	Percentage       float64 `bigquery:"percentage"`        // percentage
	SpecializationID int64   `bigquery:"specialization_id"` // specialization_id
	PassingYear      string  `bigquery:"passing_year"`      // passing_year
	OutOf            float64 `bigquery:"out_of"`            // out_of
	Grade            string  `bigquery:"grade"`             // grade
	Attachments      string  `bigquery:"attachments"`       // attachments
	FinalProject     string  `bigquery:"final_project"`     // final_project
	Description      string  `bigquery:"description"`       // description
	CreatedBy        int64   `bigquery:"created_by"`        // created_by
	DocID            int64   `bigquery:"doc_id"`            // doc_id
}

func ParseValueToUserEducation(education map[string]interface{}) interface{} {

	var item UserEducation

	common.GetMapInt64Value(education, "id", &item.ID)
	common.GetMapInt64Value(education, "user_id", &item.UserID)
	common.GetMapInt64Value(education, "degree_level_id", &item.DegreeLevelID)
	common.GetMapStringValue(education, "degree_title", &item.DegreeTitle)
	common.GetMapStringValue(education, "university_name", &item.UniversityName)
	common.GetMapStringValue(education, "start_date", &item.StartDate)
	common.GetMapStringValue(education, "end_date", &item.EndDate)
	common.GetMapInt64Value(education, "country_id", &item.CountryID)
	common.GetMapInt64Value(education, "city_id", &item.CityID)
	common.GetMapInt64Value(education, "state_id", &item.StateID)
	common.GetMapStringValue(education, "created_date", &item.CreatedDate)
	common.GetMapStringValue(education, "changed_date", &item.ChangedDate)
	common.GetMapStringValue(education, "deleted_date", &item.DeletedDate)
	common.GetMapFloat64Value(education, "gpa", &item.Gpa)
	common.GetMapFloat64Value(education, "percentage", &item.Percentage)
	common.GetMapInt64Value(education, "specialization_id", &item.SpecializationID)
	common.GetMapStringValue(education, "passing_year", &item.PassingYear)
	common.GetMapFloat64Value(education, "out_of", &item.OutOf)
	common.GetMapStringValue(education, "grade", &item.Grade)
	common.GetMapStringValue(education, "attachments", &item.Attachments)
	common.GetMapStringValue(education, "final_project", &item.FinalProject)
	common.GetMapStringValue(education, "description", &item.Description)
	common.GetMapInt64Value(education, "created_by", &item.CreatedBy)
	common.GetMapInt64Value(education, "doc_id", &item.DocID)

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
