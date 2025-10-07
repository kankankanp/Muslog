package repository

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

type BandRecruitmentFilter struct {
	Keyword  string
	Genre    string
	Location string
	Status   string
	Page     int
	PerPage  int
}

type BandRecruitmentRepository interface {
	Create(ctx context.Context, recruitment *entity.BandRecruitment) error
	Update(ctx context.Context, recruitment *entity.BandRecruitment) error
	FindByID(ctx context.Context, id string) (*entity.BandRecruitment, error)
	FindByIDForUser(ctx context.Context, id string, userID string) (*entity.BandRecruitment, error)
	Search(ctx context.Context, filter BandRecruitmentFilter) ([]*entity.BandRecruitment, int64, error)
}
