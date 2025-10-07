package repository

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"github.com/kankankanp/Muslog/internal/infrastructure/mapper"
	"github.com/kankankanp/Muslog/internal/infrastructure/model"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type bandRecruitmentRepositoryImpl struct {
	DB *gorm.DB
}

func NewBandRecruitmentRepository(db *gorm.DB) domainRepo.BandRecruitmentRepository {
	return &bandRecruitmentRepositoryImpl{DB: db}
}

func (r *bandRecruitmentRepositoryImpl) Create(ctx context.Context, recruitment *entity.BandRecruitment) error {
	m := mapper.FromBandRecruitmentEntity(recruitment)
	if err := r.DB.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	recruitment.ID = m.ID
	recruitment.CreatedAt = m.CreatedAt
	recruitment.UpdatedAt = m.UpdatedAt
	return nil
}

func (r *bandRecruitmentRepositoryImpl) Update(ctx context.Context, recruitment *entity.BandRecruitment) error {
	data := map[string]interface{}{
		"title":            recruitment.Title,
		"description":      recruitment.Description,
		"genre":            recruitment.Genre,
		"location":         recruitment.Location,
		"recruiting_parts": pq.StringArray(append([]string(nil), recruitment.RecruitingParts...)),
		"skill_level":      recruitment.SkillLevel,
		"contact":          recruitment.Contact,
		"deadline":         recruitment.Deadline,
		"status":           recruitment.Status,
		"updated_at":       recruitment.UpdatedAt,
	}

	result := r.DB.WithContext(ctx).
		Model(&model.BandRecruitmentModel{}).
		Where("id = ? AND user_id = ?", recruitment.ID, recruitment.UserID).
		Updates(data)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *bandRecruitmentRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.BandRecruitment, error) {
	var m model.BandRecruitmentModel
	if err := r.DB.WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return mapper.ToBandRecruitmentEntity(&m), nil
}

func (r *bandRecruitmentRepositoryImpl) FindByIDForUser(ctx context.Context, id string, userID string) (*entity.BandRecruitment, error) {
	var m model.BandRecruitmentModel
	if err := r.DB.WithContext(ctx).First(&m, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return nil, err
	}
	return mapper.ToBandRecruitmentEntity(&m), nil
}

func (r *bandRecruitmentRepositoryImpl) Search(ctx context.Context, filter domainRepo.BandRecruitmentFilter) ([]*entity.BandRecruitment, int64, error) {
	page := filter.Page
	if page < 1 {
		page = 1
	}
	perPage := filter.PerPage
	if perPage <= 0 {
		perPage = 10
	}

	query := r.DB.WithContext(ctx).Model(&model.BandRecruitmentModel{})

	if filter.Keyword != "" {
		like := "%" + filter.Keyword + "%"
		query = query.Where("title ILIKE ? OR description ILIKE ?", like, like)
	}
	if filter.Genre != "" {
		query = query.Where("genre = ?", filter.Genre)
	}
	if filter.Location != "" {
		query = query.Where("location = ?", filter.Location)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	var total int64
	if err := query.Session(&gorm.Session{NewDB: true}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var models []model.BandRecruitmentModel
	if err := query.Order("created_at DESC").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&models).Error; err != nil {
		return nil, 0, err
	}

	recruitments := make([]*entity.BandRecruitment, 0, len(models))
	for i := range models {
		recruitments = append(recruitments, mapper.ToBandRecruitmentEntity(&models[i]))
	}

	return recruitments, total, nil
}
