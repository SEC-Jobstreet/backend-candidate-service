package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"mime/multipart"
	"time"
)

type CandidateProfile struct {
	UserID             string        `json:"user_id" gorm:"column:user_id"`
	GoogleID           pgtype.Int8   `json:"google_id" gorm:"column:google_id"`
	Email              pgtype.Text   `json:"email" gorm:"column:email"`
	FirstName          pgtype.Text   `json:"first_name" gorm:"column:first_name"`
	LastName           pgtype.Text   `json:"last_name" gorm:"column:last_name"`
	ProfileImage       pgtype.Text   `json:"profile_image" gorm:"profile_image"`
	FirstNameProfile   pgtype.Text   `json:"first_name_profile" gorm:"first_name_profile"`
	LastNameProfile    pgtype.Text   `json:"last_name_profile" gorm:"last_name_profile"`
	Phone              pgtype.Text   `json:"phone" gorm:"column:phone"`
	PhoneNumberCountry pgtype.Text   `json:"phone_number_country" gorm:"column:phone_number_country"`
	Address            pgtype.Text   `json:"address" gorm:"column:address"`
	CurrentLocation    pgtype.Text   `json:"current_location" gorm:"column:current_location"`
	PrivacySetting     pgtype.Text   `json:"privacy_setting" gorm:"privacy_setting"`
	WorkEligibility    pgtype.Text   `json:"work_eligibility" gorm:"work_eligibility"`
	ResumeLink         pgtype.Text   `json:"resume_link" gorm:"resume_link"`
	Resume             pgtype.Text   `json:"resume" gorm:"resume"`
	CurrentRole        pgtype.Text   `json:"current_role" gorm:"current_role"`
	WorkWhenever       pgtype.Bool   `json:"work_whenever" gorm:"work_whenever"`
	WorkShift          pgtype.Text   `json:"work_shift" gorm:"work_shift"`
	LocationLat        pgtype.Float8 `json:"location_lat" gorm:"location_lat"`
	LocationLon        pgtype.Float8 `json:"location_lon" gorm:"location_lon"`
	Visa               pgtype.Bool   `json:"visa" gorm:"visa"`
	Description        pgtype.Text   `json:"description" gorm:"description"`
	Position           pgtype.Text   `json:"position" gorm:"position"`
	StartDate          pgtype.Date   `json:"start_date" gorm:"start_date"`
	ShareProfile       pgtype.Bool   `json:"share_profile" gorm:"share_profile"`
	UpdatedAt          time.Time     `json:"updated_at" gorm:"updated_at"`
	CreatedAt          time.Time     `json:"created_at" gorm:"created_at"`
}

func (CandidateProfile) TableName() string {
	return "candidate_profile"
}

type ProfileResponse struct {
	ProfileUpdated     bool              `json:"profileUpdated"`
	Profile            Profile           `json:"profile"`
	MetaTitle          string            `json:"metaTitle"`
	Site               Site              `json:"site"`
	IndustriesAndRoles []IndustryAndRole `json:"industriesAndRoles"`
}

type Profile struct {
	GivenName                  string            `json:"givenName"`
	SurName                    string            `json:"surName"`
	PhoneNumber                string            `json:"phoneNumber"`
	CurrentLocation            string            `json:"currentLocation"`
	PrivacySetting             string            `json:"privacySetting"`
	WorkEligibility            map[string]string `json:"workEligibility"`
	Resume                     string            `json:"resume,omitempty"`
	CurrentRole                string            `json:"currentRole"`
	CurrentRoleStartDate       string            `json:"currentRoleStartDate,omitempty"`
	AboutMe                    string            `json:"aboutMe,omitempty"`
	ShiftAvailability          ShiftAvailability `json:"shiftAvailability"`
	CreatedAt                  time.Time         `json:"createdAt"`
	UpdatedAt                  time.Time         `json:"updatedAt"`
	CurrentLocationCoordinates Coordinates       `json:"currentLocationCoordinates"`
	CandidateId                string            `json:"candidateId"`
	Email                      string            `json:"email"`
	PhoneNumberCountryAlpha2   string            `json:"phoneNumberCountryAlpha2"`
}

type ShiftAvailability struct {
	AnyTimeShiftAvailability  bool                  `json:"anyTimeShiftAvailability"`
	SpecificShiftAvailability map[string]ShiftTimes `json:"specificShiftAvailability"`
}

type ShiftTimes struct {
	Morning   bool `json:"morning"`
	Afternoon bool `json:"afternoon"`
	Evening   bool `json:"evening"`
}

type Coordinates struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Site struct {
	ID        string    `json:"id"`
	Hosts     []string  `json:"hosts"`
	Country   Country   `json:"country"`
	Brand     Brand     `json:"brand"`
	Analytics Analytics `json:"analytics"`
}

type Country struct {
	IsoCode         string   `json:"isoCode"`
	Name            string   `json:"name"`
	NameWithArticle string   `json:"nameWithArticle"`
	LanguageTags    []string `json:"languageTags"`
}

type Brand struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CopyrightName string `json:"copyrightName"`
}

type Analytics struct {
	Prod    string `json:"prod"`
	Sandbox string `json:"sandbox,omitempty"`
}

type IndustryAndRole struct {
	Value       string `json:"value"`
	DisplayName string `json:"displayName"`
	Roles       []Role `json:"roles"`
}

type Role struct {
	Value       string `json:"value"`
	DisplayName string `json:"displayName"`
}

type UserProfileEditRequest struct {
	LastName                    string               `json:"last_name" form:"last_name" validate:"required"`
	FirstName                   string               `json:"first_name" form:"first_name" validate:"required"`
	PhoneNumberCountry          string               `json:"phone_number_country" form:"phone_number_country" validate:"required"`
	PhoneNumber                 string               `json:"phone" form:"phone" validate:"required"`
	CurrentLocation             string               `json:"current_location" form:"current_location" validate:"required"`
	LocationLon                 string               `json:"location_lon" form:"location_lon"`
	LocationLat                 string               `json:"location_lat" form:"location_lat"`
	AddressComponentsSerialized string               `json:"address_components_serialized" form:"address_components_serialized"`
	WorkEligibility             string               `json:"work_eligibility" form:"work_eligibility"`
	AboutMe                     string               `json:"about_me" form:"about_me"`
	CurrentRole                 string               `json:"current_role" form:"current_role"`
	StartDate                   string               `json:"start_date" form:"start_date"`
	WorkShift                   string               `json:"work_shift" form:"work_shift"`
	PrivacySetting              string               `json:"privacy_setting" form:"privacy_setting"`
	Visa                        bool                 `json:"visa" form:"visa"`
	Resume                      multipart.FileHeader `json:"resume" form:"resume"`
}

type UserProfileCreateRequest struct {
	UserId string `json:"userId"`
}

type GetCandidateProfilesResponse struct {
	UserID             string      `json:"user_id"`
	GoogleID           pgtype.Int8 `json:"google_id"`
	Email              string      `json:"email"`
	FirstName          string      `json:"first_name"`
	LastName           string      `json:"last_name"`
	ProfileImage       string      `json:"profile_image"`
	FirstNameProfile   string      `json:"first_name_profile"`
	LastNameProfile    string      `json:"last_name_profile"`
	Phone              string      `json:"phone"`
	PhoneNumberCountry string      `json:"phone_number_country"`
	Address            string      `json:"address"`
	CurrentLocation    string      `json:"current_location"`
	PrivacySetting     string      `json:"privacy_setting"`
	ResumeLink         string      `json:"resume_link"`
	Resume             string      `json:"resume"`
	CurrentRole        string      `json:"current_role"`
	WorkWhenever       pgtype.Bool `json:"work_whenever"`
	WorkShift          string      `json:"work_shift"`
	LocationLat        float64     `json:"location_lat"`
	LocationLon        float64     `json:"location_lon"`
	Visa               pgtype.Bool `json:"visa"`
	Description        string      `json:"description"`
	Position           string      `json:"position"`
	StartDate          pgtype.Date `json:"start_date"`
	ShareProfile       pgtype.Bool `json:"share_profile"`
	UpdatedAt          time.Time   `json:"updated_at"`
	CreatedAt          time.Time   `json:"created_at"`
}
