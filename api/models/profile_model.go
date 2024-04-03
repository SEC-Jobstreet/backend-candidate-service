package models

import (
	"mime/multipart"
	"time"
)

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
	GivenName                            string               `form:"givenName"`
	SurName                              string               `form:"surName"`
	PhoneNumberCountry                   string               `form:"phoneNumberCountry"`
	PhoneNumber                          string               `form:"phoneNumber"`
	CurrentLocation                      string               `form:"currentLocation"`
	CurrentLocationCoordinatesSerialized string               `form:"currentLocationCoordinatesSerialized"`
	AddressComponentsSerialized          string               `form:"addressComponentsSerialized"`
	WorkEligibility                      map[string]string    `form:"workEligibility"`
	AboutMe                              string               `form:"aboutMe"`
	CurrentRole                          string               `form:"currentRole"`
	CurrentRoleStartDate                 string               `form:"currentRoleStartDate"`
	AnyTimeShiftAvailability             string               `form:"shiftAvailability.anyTimeShiftAvailability"`
	PrivacySetting                       string               `form:"privacySetting"`
	Resume                               multipart.FileHeader `form:"resume"` // Assume this is the path to the resume for simplicity
}
