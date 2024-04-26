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
	ProfileImage       pgtype.Text   `json:"profile_image" json:"profile_image"`
	FirstNameProfile   pgtype.Text   `json:"first_name_profile" json:"first_name_profile"`
	LastNameProfile    pgtype.Text   `json:"last_name_profile" json:"last_name_profile"`
	Phone              pgtype.Text   `json:"phone" gorm:"column:phone"`
	PhoneNumberCountry pgtype.Text   `json:"phone_number_country" gorm:"column:phone_number_country"`
	Address            pgtype.Text   `json:"address" gorm:"column:address"`
	CurrentLocation    pgtype.Text   `json:"current_location" gorm:"column:current_location"`
	PrivacySetting     pgtype.Text   `json:"privacy_setting" json:"privacy_setting"`
	WorkEligibility    []byte        `json:"work_eligibility" json:"work_eligibility"`
	ResumeLink         pgtype.Text   `json:"resume_link" json:"resume_link"`
	CurrentRole        pgtype.Text   `json:"current_role" json:"current_role"`
	WorkWhenever       pgtype.Bool   `json:"work_whenever" json:"work_whenever"`
	WorkShift          []byte        `json:"work_shift" json:"work_shift"`
	LocationLat        pgtype.Float8 `json:"location_lat" json:"location_lat"`
	LocationLon        pgtype.Float8 `json:"location_lon" json:"location_lon"`
	Visa               pgtype.Bool   `json:"visa" json:"visa"`
	Description        pgtype.Text   `json:"description" json:"description"`
	Position           pgtype.Text   `json:"position" json:"position"`
	StartDate          pgtype.Date   `json:"start_date" json:"start_date"`
	ShareProfile       pgtype.Bool   `json:"share_profile" json:"share_profile"`
	UpdatedAt          time.Time     `json:"updated_at" json:"updated_at"`
	CreatedAt          time.Time     `json:"created_at" json:"created_at"`
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
	LastName                    string                 `form:"last_name"`
	FirstName                   string                 `form:"first_name"`
	PhoneNumberCountry          string                 `form:"phone_number_country"`
	PhoneNumber                 string                 `form:"phone"`
	CurrentLocation             string                 `form:"current_location"`
	LocationLon                 string                 `form:"location_lon"`
	LocationLat                 string                 `form:"location_lat"`
	AddressComponentsSerialized string                 `form:"address_components_serialized"`
	WorkEligibility             map[string]string      `form:"work_eligibility"`
	AboutMe                     string                 `form:"about_me"`
	CurrentRole                 string                 `form:"current_role"`
	StartDate                   string                 `form:"start_date"`
	WorkShift                   map[string]interface{} `form:"work_shift"`
	PrivacySetting              string                 `form:"privacy_setting"`
	Resume                      multipart.FileHeader   `form:"resume"`
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
	WorkEligibility    interface{} `json:"work_eligibility"`
	ResumeLink         string      `json:"resume_link"`
	CurrentRole        string      `json:"current_role"`
	WorkWhenever       pgtype.Bool `json:"work_whenever"`
	WorkShift          interface{} `json:"work_shift"`
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
