package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Applications struct {
	ID          uuid.UUID `gorm:"primarykey"`
	CandidateID string    `gorm:"not null; index:,unique,composite:candidate_job_key"`
	JobID       uuid.UUID `gorm:"not null; index:,unique,composite:candidate_job_key"`
	Status      string
	// EmployerID  string

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}

func MigrateApplications(db *gorm.DB) error {
	err := db.AutoMigrate(&Applications{})
	return err
}
