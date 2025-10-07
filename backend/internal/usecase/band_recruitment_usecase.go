package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"gorm.io/gorm"
)

var (
	ErrApplyOwnRecruitment    = errors.New("cannot apply to own recruitment")
	ErrRecruitmentClosed      = errors.New("recruitment is closed")
	ErrAlreadyApplied         = errors.New("already applied to this recruitment")
	ErrRecruitmentNotFound    = errors.New("recruitment not found")
	ErrRecruitmentNoPrivilege = errors.New("no permission for this recruitment")
)

type BandRecruitmentFilterInput struct {
	Keyword  string
	Genre    string
	Location string
	Status   string
	Page     int
	PerPage  int
}

type CreateBandRecruitmentInput struct {
	Title           string
	Description     string
	Genre           string
	Location        string
	RecruitingParts []string
	SkillLevel      string
	Contact         string
	Deadline        *time.Time
	Status          string
	UserID          string
}

type UpdateBandRecruitmentInput struct {
	ID              string
	Title           string
	Description     string
	Genre           string
	Location        string
	RecruitingParts []string
	SkillLevel      string
	Contact         string
	Deadline        *time.Time
	Status          string
	UserID          string
}

type ApplyBandRecruitmentInput struct {
	BandRecruitmentID string
	ApplicantID       string
	Message           string
}

type BandRecruitmentUsecase interface {
	SearchBandRecruitments(ctx context.Context, filter BandRecruitmentFilterInput, userID string) ([]*entity.BandRecruitment, int64, error)
	GetBandRecruitmentByID(ctx context.Context, id string, userID string) (*entity.BandRecruitment, error)
	CreateBandRecruitment(ctx context.Context, input CreateBandRecruitmentInput) (*entity.BandRecruitment, error)
	UpdateBandRecruitment(ctx context.Context, input UpdateBandRecruitmentInput) (*entity.BandRecruitment, error)
	ApplyToBandRecruitment(ctx context.Context, input ApplyBandRecruitmentInput) error
}

type bandRecruitmentUsecaseImpl struct {
	recruitmentRepo domainRepo.BandRecruitmentRepository
	applicationRepo domainRepo.BandApplicationRepository
	txManager       domainRepo.TransactionManager
}

func NewBandRecruitmentUsecase(
	recruitmentRepo domainRepo.BandRecruitmentRepository,
	applicationRepo domainRepo.BandApplicationRepository,
	txManager domainRepo.TransactionManager,
) BandRecruitmentUsecase {
	return &bandRecruitmentUsecaseImpl{
		recruitmentRepo: recruitmentRepo,
		applicationRepo: applicationRepo,
		txManager:       txManager,
	}
}

func (u *bandRecruitmentUsecaseImpl) SearchBandRecruitments(ctx context.Context, filter BandRecruitmentFilterInput, userID string) ([]*entity.BandRecruitment, int64, error) {
	recruitFilter := domainRepo.BandRecruitmentFilter{
		Keyword:  filter.Keyword,
		Genre:    filter.Genre,
		Location: filter.Location,
		Status:   filter.Status,
		Page:     filter.Page,
		PerPage:  filter.PerPage,
	}

	recruitments, total, err := u.recruitmentRepo.Search(ctx, recruitFilter)
	if err != nil {
		return nil, 0, err
	}

	if err := u.enrichRecruitments(ctx, recruitments, userID); err != nil {
		return nil, 0, err
	}

	return recruitments, total, nil
}

func (u *bandRecruitmentUsecaseImpl) GetBandRecruitmentByID(ctx context.Context, id string, userID string) (*entity.BandRecruitment, error) {
	recruitment, err := u.recruitmentRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecruitmentNotFound
		}
		return nil, err
	}

	if err := u.enrichRecruitments(ctx, []*entity.BandRecruitment{recruitment}, userID); err != nil {
		return nil, err
	}

	return recruitment, nil
}

func (u *bandRecruitmentUsecaseImpl) CreateBandRecruitment(ctx context.Context, input CreateBandRecruitmentInput) (*entity.BandRecruitment, error) {
	recruitment := &entity.BandRecruitment{
		Title:           input.Title,
		Description:     input.Description,
		Genre:           input.Genre,
		Location:        input.Location,
		RecruitingParts: append([]string(nil), input.RecruitingParts...),
		SkillLevel:      input.SkillLevel,
		Contact:         input.Contact,
		Deadline:        input.Deadline,
		Status:          defaultStatus(input.Status),
		UserID:          input.UserID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := u.txManager.Do(ctx, func(txRepo domainRepo.RepositoryProvider) error {
		return txRepo.BandRecruitmentRepository().Create(ctx, recruitment)
	}); err != nil {
		return nil, err
	}

	created, err := u.recruitmentRepo.FindByID(ctx, recruitment.ID)
	if err != nil {
		return recruitment, nil
	}
	if err := u.enrichRecruitments(ctx, []*entity.BandRecruitment{created}, input.UserID); err != nil {
		return created, nil
	}
	return created, nil
}

func (u *bandRecruitmentUsecaseImpl) UpdateBandRecruitment(ctx context.Context, input UpdateBandRecruitmentInput) (*entity.BandRecruitment, error) {
	var updatedRecruitment *entity.BandRecruitment

	err := u.txManager.Do(ctx, func(txRepo domainRepo.RepositoryProvider) error {
		recruitment, err := txRepo.BandRecruitmentRepository().FindByIDForUser(ctx, input.ID, input.UserID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrRecruitmentNoPrivilege
			}
			return err
		}

		recruitment.Title = input.Title
		recruitment.Description = input.Description
		recruitment.Genre = input.Genre
		recruitment.Location = input.Location
		recruitment.RecruitingParts = append([]string(nil), input.RecruitingParts...)
		recruitment.SkillLevel = input.SkillLevel
		recruitment.Contact = input.Contact
		recruitment.Deadline = input.Deadline
		if input.Status != "" {
			recruitment.Status = input.Status
		}
		recruitment.UpdatedAt = time.Now()

		if err := txRepo.BandRecruitmentRepository().Update(ctx, recruitment); err != nil {
			return err
		}

		updatedRecruitment = recruitment
		return nil
	})

	if err != nil {
		return nil, err
	}

	reloaded, err := u.recruitmentRepo.FindByID(ctx, updatedRecruitment.ID)
	if err != nil {
		return updatedRecruitment, nil
	}
	if err := u.enrichRecruitments(ctx, []*entity.BandRecruitment{reloaded}, updatedRecruitment.UserID); err != nil {
		return reloaded, nil
	}

	return reloaded, nil
}

func (u *bandRecruitmentUsecaseImpl) ApplyToBandRecruitment(ctx context.Context, input ApplyBandRecruitmentInput) error {
	return u.txManager.Do(ctx, func(txRepo domainRepo.RepositoryProvider) error {
		recruitment, err := txRepo.BandRecruitmentRepository().FindByID(ctx, input.BandRecruitmentID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrRecruitmentNotFound
			}
			return err
		}

		if recruitment.UserID == input.ApplicantID {
			return ErrApplyOwnRecruitment
		}
		if recruitment.Status == "closed" {
			return ErrRecruitmentClosed
		}

		alreadyApplied, err := txRepo.BandApplicationRepository().HasApplied(ctx, input.BandRecruitmentID, input.ApplicantID)
		if err != nil {
			return err
		}
		if alreadyApplied {
			return ErrAlreadyApplied
		}

		application := &entity.BandApplication{
			BandRecruitmentID: input.BandRecruitmentID,
			ApplicantID:       input.ApplicantID,
			Message:           input.Message,
			CreatedAt:         time.Now(),
		}

		return txRepo.BandApplicationRepository().Create(ctx, application)
	})
}

func (u *bandRecruitmentUsecaseImpl) enrichRecruitments(ctx context.Context, recruitments []*entity.BandRecruitment, userID string) error {
	if len(recruitments) == 0 {
		return nil
	}

	ids := make([]string, 0, len(recruitments))
	for _, r := range recruitments {
		if r != nil {
			ids = append(ids, r.ID)
		}
	}

	counts, err := u.applicationRepo.CountByRecruitmentIDs(ctx, ids)
	if err != nil {
		return err
	}

	appliedMap := make(map[string]bool)
	if userID != "" {
		appliedMap, err = u.applicationRepo.FindAppliedRecruitmentIDs(ctx, ids, userID)
		if err != nil {
			return err
		}
	}

	for _, r := range recruitments {
		if r == nil {
			continue
		}
		r.ApplicationsCount = counts[r.ID]
		r.HasApplied = appliedMap[r.ID]
	}

	return nil
}

func defaultStatus(status string) string {
	s := status
	if s == "" {
		s = "open"
	}
	return s
}
