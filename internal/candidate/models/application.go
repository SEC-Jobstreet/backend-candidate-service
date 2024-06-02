package models

import (
	"github.com/google/uuid"
)

type Application struct {
	ID          uuid.UUID `gorm:"primarykey"`
	CandidateID string    `gorm:"not null; index:,unique,composite:application_key"`
	JobID       uuid.UUID `gorm:"not null; index:,unique,composite:application_key"`
	Status      string

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
