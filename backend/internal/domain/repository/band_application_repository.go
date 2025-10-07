package repository

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

type BandApplicationRepository interface {
	Create(ctx context.Context, application *entity.BandApplication) error
	CountByRecruitmentIDs(ctx context.Context, recruitmentIDs []string) (map[string]int64, error)
	HasApplied(ctx context.Context, recruitmentID, applicantID string) (bool, error)
	FindAppliedRecruitmentIDs(ctx context.Context, recruitmentIDs []string, applicantID string) (map[string]bool, error)
	FindRecruitmentsByApplicant(ctx context.Context, applicantID string) ([]*entity.BandRecruitment, error)
}
