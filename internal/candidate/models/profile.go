package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Profile struct {
	Username string `gorm:"primarykey; not null; index:username,unique" json:"username" bson:"username"`

	FirstName    string `gorm:"not null" json:"first_name" bson:"first_name"`
	LastName     string `gorm:"not null" json:"last_name" bson:"last_name"`
	CountryPhone string `bson:"CountryPhone"`
	Phone        string `bson:"Phone"`
	Address      string `bson:"Address"`
	Latitude     string `bson:"Latitude"`
	Longitude    string `bson:"Longitude"`
	Visa         bool   `bson:"Visa"`
	Description  string `bson:"Description"`

	CurrentPosition string `bson:"CurrentPosition"`
	StartDate       int64  `bson:"StartDate"`

	WorkWhenever bool   `bson:"WorkWhenever"`
	WorkShift    string `bson:"WorkShift"`

	ShareProfile bool   `bson:"ShareProfile"`
	ResumeLink   string `bson:"ResumeLink"`
	ResumeName   string `bson:"ResumeName"`

	UpdatedAt int64 `gorm:"autoUpdateTime" bson:"UpdatedAt"`
	CreatedAt int64 `gorm:"autoCreateTime" bson:"CreatedAt"`
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&Profile{}, &Application{}, &SavedJob{})
	return err
}

func (o *Profile) String() string {
	return fmt.Sprintf("Username: {%s}, FirstName: {%s}, LastName: {%s}, Visa: {%v}, "+
		"Description: {%s}, CurrentPosition: {%s}, StartDate: {%v}, WorkWhenever: {%v}, WorkShift: {%s}, ShareProfile: {%v}, ResumeLink: {%s}, ResumeName: {%s}",
		o.Username,
		o.FirstName,
		o.LastName,
		o.Visa,
		o.Description,
		o.CurrentPosition,
		o.StartDate,
		o.WorkWhenever,
		o.WorkShift,
		o.ShareProfile,
		o.ResumeLink,
		o.ResumeName,
	)
}
