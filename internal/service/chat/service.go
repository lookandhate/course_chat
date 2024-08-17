package service

import (
	"github.com/lookandhate/course_chat/internal/cache"
	"github.com/lookandhate/course_chat/internal/repository"
	"github.com/lookandhate/course_platform_lib/pkg/db"
)

type Service struct {
	repo      repository.ChatRepository
	txManager db.TxManager
	cache     cache.ChatCache
}

// NewService creates Service with given repo.
func NewService(repo repository.ChatRepository, manager db.TxManager, userCache cache.ChatCache) *Service {
	return &Service{
		repo:      repo,
		txManager: manager,
		cache:     userCache,
	}
}
