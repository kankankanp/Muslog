package repository

import "context"

type TransactionManager interface {
	Do(ctx context.Context, fn func(txRepo RepositoryProvider) error) error
}

type RepositoryProvider interface {
	PostRepository() PostRepository
	TagRepository() TagRepository
	UserRepository() UserRepository
	MessageRepository() MessageRepository
	CommunityRepository() CommunityRepository
	LikeRepository() LikeRepository
	BandRecruitmentRepository() BandRecruitmentRepository
	BandApplicationRepository() BandApplicationRepository
}
