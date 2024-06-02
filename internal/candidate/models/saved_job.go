package models

import (
	"github.com/google/uuid"
)

type SavedJob struct {
	ID          uuid.UUID `gorm:"primarykey"`
	CandidateID string    `gorm:"not null; index:,unique,composite:saved_job_key"`
	JobID       uuid.UUID `gorm:"not null; index:,unique,composite:saved_job_key"`

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
