package response

import (
	"time"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

type BandRecruitmentResponse struct {
	ID                string     `json:"id"`
	Title             string     `json:"title"`
	Description       string     `json:"description"`
	Genre             string     `json:"genre"`
	Location          string     `json:"location"`
	RecruitingParts   []string   `json:"recruitingParts"`
	SkillLevel        string     `json:"skillLevel"`
	Contact           string     `json:"contact"`
	Deadline          *time.Time `json:"deadline"`
	Status            string     `json:"status"`
	UserID            string     `json:"userId"`
	ApplicationsCount int64      `json:"applicationsCount"`
	HasApplied        bool       `json:"hasApplied"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
}

type BandRecruitmentListResponse struct {
	Message      string                    `json:"message"`
	Recruitments []BandRecruitmentResponse `json:"recruitments"`
	TotalCount   int64                     `json:"totalCount"`
	Page         int                       `json:"page"`
	PerPage      int                       `json:"perPage"`
}

type BandRecruitmentDetailResponse struct {
	Message     string                  `json:"message"`
	Recruitment BandRecruitmentResponse `json:"recruitment"`
}

func ToBandRecruitmentResponse(e *entity.BandRecruitment) BandRecruitmentResponse {
	if e == nil {
		return BandRecruitmentResponse{}
	}

	return BandRecruitmentResponse{
		ID:                e.ID,
		Title:             e.Title,
		Description:       e.Description,
		Genre:             e.Genre,
		Location:          e.Location,
		RecruitingParts:   append([]string(nil), e.RecruitingParts...),
		SkillLevel:        e.SkillLevel,
		Contact:           e.Contact,
		Deadline:          e.Deadline,
		Status:            e.Status,
		UserID:            e.UserID,
		ApplicationsCount: e.ApplicationsCount,
		HasApplied:        e.HasApplied,
		CreatedAt:         e.CreatedAt,
		UpdatedAt:         e.UpdatedAt,
	}
}

func ToBandRecruitmentResponses(items []*entity.BandRecruitment) []BandRecruitmentResponse {
	if len(items) == 0 {
		return []BandRecruitmentResponse{}
	}
	responses := make([]BandRecruitmentResponse, 0, len(items))
	for _, item := range items {
		responses = append(responses, ToBandRecruitmentResponse(item))
	}
	return responses
}
