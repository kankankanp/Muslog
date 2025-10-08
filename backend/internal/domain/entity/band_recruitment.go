package entity

import "time"

type BandRecruitment struct {
	ID                string
	Title             string
	Description       string
	Genre             string
	Location          string
	RecruitingParts   []string
	SkillLevel        string
	Contact           string
	Deadline          *time.Time
	Status            string
	UserID            string
	ApplicationsCount int64
	HasApplied        bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type BandApplication struct {
	ID                string
	BandRecruitmentID string
	ApplicantID       string
	Message           string
	CreatedAt         time.Time
}
