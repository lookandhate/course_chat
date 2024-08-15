package service

import (
	"github.com/lookandhate/course_chat/internal/repository"
	"github.com/lookandhate/course_platform_lib/pkg/db"
)

type Service struct {
	repo      repository.ChatRepository
	txManager db.TxManager
}

// NewService creates Service with given repo.
func NewService(repo repository.ChatRepository, manager db.TxManager) *Service {
	return &Service{repo: repo, txManager: manager}
}
