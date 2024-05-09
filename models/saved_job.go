package models

import (
	"time"

	"gorm.io/gorm"
)

type SavedJob struct {
	gorm.Model
	EmployerID   uint   `gorm:"not null"`
	EnterpriseID uint   `gorm:"not null"`
	Title        string `gorm:"not null"`
	JobType      string `gorm:"not null"`
	WorkShift    string `gorm:"not null"`
	Visa         bool   `gorm:"not null"`
	Experience   string `gorm:"not null"`
	StartDate    time.Time
	Currency     string `gorm:"not null"`
	ExactSalary  float64
	RangeSalary  string
	Description  string `gorm:"not null"`
	CreatedAt    time.Time
	ExpireAt     time.Time
}

func MigrateSavedJobs(db *gorm.DB) error {
	err := db.AutoMigrate(&SavedJob{})
	return err
}
