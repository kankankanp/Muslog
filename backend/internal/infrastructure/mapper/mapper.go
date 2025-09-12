package mapper

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/kankankanp/Muslog/internal/infrastructure/model"
)

// User
func ToUserEntity(m *model.UserModel) *entity.User {
	if m == nil {
		return nil
	}
	return &entity.User{
		ID:              m.ID,
		Name:            m.Name,
		Email:           m.Email,
		Password:        m.Password,
		GoogleID:        m.GoogleID,
		ProfileImageUrl: m.ProfileImageUrl,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func FromUserEntity(e *entity.User) *model.UserModel {
	if e == nil {
		return nil
	}
	return &model.UserModel{
		ID:              e.ID,
		Name:            e.Name,
		Email:           e.Email,
		Password:        e.Password,
		GoogleID:        e.GoogleID,
		ProfileImageUrl: e.ProfileImageUrl,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
	}
}

// Post
func ToPostEntity(m *model.PostModel) *entity.Post {
	if m == nil {
		return nil
	}
	tracks := make([]entity.Track, 0, len(m.Tracks))
	for _, t := range m.Tracks {
		tracks = append(tracks, *ToTrackEntity(&t))
	}

	tags := make([]entity.Tag, 0, len(m.Tags))
	for _, tg := range m.Tags {
		tags = append(tags, *ToTagEntity(&tg))
	}

	return &entity.Post{
		ID:             m.ID,
		Title:          m.Title,
		Description:    m.Description,
		UserID:         m.UserID,
		HeaderImageUrl: m.HeaderImageUrl,
		Tracks:         tracks,
		Tags:           tags,
		LikesCount:     m.LikesCount,
		IsLiked:        m.IsLiked,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}

func FromPostEntity(e *entity.Post) *model.PostModel {
	if e == nil {
		return nil
	}
	tracks := make([]model.TrackModel, 0, len(e.Tracks))
	for _, t := range e.Tracks {
		tracks = append(tracks, *FromTrackEntity(&t))
	}

	tags := make([]model.TagModel, 0, len(e.Tags))
	for _, tg := range e.Tags {
		tags = append(tags, *FromTagEntity(&tg))
	}

	return &model.PostModel{
		ID:             e.ID,
		Title:          e.Title,
		Description:    e.Description,
		UserID:         e.UserID,
		HeaderImageUrl: e.HeaderImageUrl,
		Tracks:         tracks,
		Tags:           tags,
		LikesCount:     e.LikesCount,
		IsLiked:        e.IsLiked,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
}

// Track
func ToTrackEntity(m *model.TrackModel) *entity.Track {
	if m == nil {
		return nil
	}
	return &entity.Track{
		ID:            m.ID,
		SpotifyID:     m.SpotifyID,
		Name:          m.Name,
		ArtistName:    m.ArtistName,
		AlbumImageUrl: m.AlbumImageUrl,
		PostID:        m.PostID,
	}
}

func FromTrackEntity(e *entity.Track) *model.TrackModel {
	if e == nil {
		return nil
	}
	return &model.TrackModel{
		ID:            e.ID,
		SpotifyID:     e.SpotifyID,
		Name:          e.Name,
		ArtistName:    e.ArtistName,
		AlbumImageUrl: e.AlbumImageUrl,
		PostID:        e.PostID,
	}
}

// Tag
func ToTagEntity(m *model.TagModel) *entity.Tag {
	if m == nil {
		return nil
	}
	return &entity.Tag{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func FromTagEntity(e *entity.Tag) *model.TagModel {
	if e == nil {
		return nil
	}
	return &model.TagModel{
		ID:        e.ID,
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func ToPostTagEntity(m *model.PostTagModel) *entity.PostTag {
	if m == nil {
		return nil
	}
	return &entity.PostTag{
		PostID:    m.PostID,
		TagID:     m.TagID,
		CreatedAt: m.CreatedAt,
	}
}

func FromPostTagEntity(e *entity.PostTag) *model.PostTagModel {
	if e == nil {
		return nil
	}
	return &model.PostTagModel{
		PostID:    e.PostID,
		TagID:     e.TagID,
		CreatedAt: e.CreatedAt,
	}
}

// Message
func ToMessageEntity(m *model.MessageModel) *entity.Message {
	if m == nil {
		return nil
	}
	return &entity.Message{
		ID:          m.ID,
		CommunityID: m.CommunityID,
		SenderID:    m.SenderID,
		Content:     m.Content,
		CreatedAt:   m.CreatedAt,
	}
}

func FromMessageEntity(e *entity.Message) *model.MessageModel {
	if e == nil {
		return nil
	}
	return &model.MessageModel{
		ID:          e.ID,
		CommunityID: e.CommunityID,
		SenderID:    e.SenderID,
		Content:     e.Content,
		CreatedAt:   e.CreatedAt,
	}
}

// Like
func ToLikeEntity(m *model.LikeModel) *entity.Like {
	if m == nil {
		return nil
	}
	return &entity.Like{
		ID:        m.ID,
		PostID:    m.PostID,
		UserID:    m.UserID,
		CreatedAt: m.CreatedAt,
	}
}

func FromLikeEntity(e *entity.Like) *model.LikeModel {
	if e == nil {
		return nil
	}
	return &model.LikeModel{
		ID:        e.ID,
		PostID:    e.PostID,
		UserID:    e.UserID,
		CreatedAt: e.CreatedAt,
	}
}

// Community
func ToCommunityEntity(m *model.CommunityModel) *entity.Community {
	if m == nil {
		return nil
	}
	return &entity.Community{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		CreatorID:   m.CreatorID,
		CreatedAt:   m.CreatedAt,
	}
}

func FromCommunityEntity(e *entity.Community) *model.CommunityModel {
	if e == nil {
		return nil
	}
	return &model.CommunityModel{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		CreatorID:   e.CreatorID,
		CreatedAt:   e.CreatedAt,
	}
}
