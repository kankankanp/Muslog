package model

import (
	"time"

	"github.com/lib/pq"
)

type BandRecruitmentModel struct {
	ID              string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title           string
	Description     string `gorm:"type:text"`
	Genre           string
	Location        string
	RecruitingParts pq.StringArray `gorm:"type:text[]"`
	SkillLevel      string
	Contact         string
	Deadline        *time.Time
	Status          string
	UserID          string                 `gorm:"type:uuid"`
	Applications    []BandApplicationModel `gorm:"foreignKey:BandRecruitmentID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
