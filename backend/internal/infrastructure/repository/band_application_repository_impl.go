package repository

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"github.com/kankankanp/Muslog/internal/infrastructure/mapper"
	"github.com/kankankanp/Muslog/internal/infrastructure/model"
	"gorm.io/gorm"
)

type bandApplicationRepositoryImpl struct {
	DB *gorm.DB
}

func NewBandApplicationRepository(db *gorm.DB) domainRepo.BandApplicationRepository {
	return &bandApplicationRepositoryImpl{DB: db}
}

func (r *bandApplicationRepositoryImpl) Create(ctx context.Context, application *entity.BandApplication) error {
	m := mapper.FromBandApplicationEntity(application)
	if err := r.DB.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	application.ID = m.ID
	application.CreatedAt = m.CreatedAt
	return nil
}

func (r *bandApplicationRepositoryImpl) CountByRecruitmentIDs(ctx context.Context, recruitmentIDs []string) (map[string]int64, error) {
	counts := make(map[string]int64)
	if len(recruitmentIDs) == 0 {
		return counts, nil
	}

	type result struct {
		BandRecruitmentID string
		Count             int64
	}
	var rows []result

	if err := r.DB.WithContext(ctx).
		Model(&model.BandApplicationModel{}).
		Select("band_recruitment_id, COUNT(*) as count").
		Where("band_recruitment_id IN ?", recruitmentIDs).
		Group("band_recruitment_id").
		Find(&rows).Error; err != nil {
		return nil, err
	}

	for _, row := range rows {
		counts[row.BandRecruitmentID] = row.Count
	}

	return counts, nil
}

func (r *bandApplicationRepositoryImpl) HasApplied(ctx context.Context, recruitmentID, applicantID string) (bool, error) {
	if recruitmentID == "" || applicantID == "" {
		return false, nil
	}

	var count int64
	if err := r.DB.WithContext(ctx).
		Model(&model.BandApplicationModel{}).
		Where("band_recruitment_id = ? AND applicant_id = ?", recruitmentID, applicantID).
		Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *bandApplicationRepositoryImpl) FindAppliedRecruitmentIDs(ctx context.Context, recruitmentIDs []string, applicantID string) (map[string]bool, error) {
	result := make(map[string]bool)
	if len(recruitmentIDs) == 0 || applicantID == "" {
		return result, nil
	}

	var rows []string
	if err := r.DB.WithContext(ctx).
		Model(&model.BandApplicationModel{}).
		Select("band_recruitment_id").
		Where("band_recruitment_id IN ? AND applicant_id = ?", recruitmentIDs, applicantID).
		Group("band_recruitment_id").
		Find(&rows).Error; err != nil {
		return nil, err
	}

	for _, id := range rows {
		result[id] = true
	}

	return result, nil
}

func (r *bandApplicationRepositoryImpl) FindRecruitmentsByApplicant(ctx context.Context, applicantID string) ([]*entity.BandRecruitment, error) {
	if applicantID == "" {
		return []*entity.BandRecruitment{}, nil
	}

	var models []model.BandRecruitmentModel
	if err := r.DB.WithContext(ctx).
		Model(&model.BandRecruitmentModel{}).
		Distinct("band_recruitments.*").
		Joins("JOIN band_applications ON band_applications.band_recruitment_id = band_recruitments.id").
		Where("band_applications.applicant_id = ?", applicantID).
		Order("band_recruitments.created_at DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	recruitments := make([]*entity.BandRecruitment, 0, len(models))
	for i := range models {
		recruitments = append(recruitments, mapper.ToBandRecruitmentEntity(&models[i]))
	}

	return recruitments, nil
}
