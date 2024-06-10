package models

import (
	"github.com/google/uuid"
)

type Jobs struct {
	ID uuid.UUID `gorm:"primarykey" json:"id"`

	EmployerID         string `gorm:"index:employer_id" json:"employer_id"`
	Status             string `gorm:"not null; default: REVIEW" json:"status"` // REVIEW, POSTED, DENIED, CLOSED
	Title              string `gorm:"not null" json:"title"`
	Type               string `json:"type"`
	WorkWhenever       bool   `json:"work_whenever"`
	WorkShift          string `json:"work_shift"`
	Description        string `gorm:"not null" json:"description"`
	Visa               bool   `json:"visa"`
	Experience         uint32 `json:"experience"`
	StartDate          int64  `json:"start_date"`
	Currency           string `json:"currency"`
	SalaryLevelDisplay string `json:"salary_level_display"`
	PaidPeriod         string `json:"paid_period"`
	ExactSalary        uint32 `json:"exact_salary"`
	RangeSalary        string `json:"range_salary"`
	ExpiresAt          int64  `json:"expires_at"`

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`

	EnterpriseID      uuid.UUID `json:"enterprise_id"`
	EnterpriseName    string    `json:"enterprise_name"`
	EnterpriseAddress string    `json:"enterprise_address"`

	Crawl         bool   `gorm:"default: false" json:"crawl"`
	JobURL        string `json:"job_url"`
	JobSourceName string `json:"job_source_name"`
}
