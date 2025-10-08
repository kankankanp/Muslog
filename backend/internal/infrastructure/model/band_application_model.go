package model

import "time"

type BandApplicationModel struct {
	ID                string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BandRecruitmentID string `gorm:"type:uuid"`
	ApplicantID       string `gorm:"type:uuid"`
	Message           string `gorm:"type:text"`
	CreatedAt         time.Time
}
