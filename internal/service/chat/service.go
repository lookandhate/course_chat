package service

import (
	"github.com/lookandhate/course_chat/internal/client/db"
	"github.com/lookandhate/course_chat/internal/repository"
)

type Service struct {
	repo      repository.ChatRepository
	txManager db.TxManager
}

// NewService creates Service with given repo.
func NewService(repo repository.ChatRepository, manager db.TxManager) *Service {
	return &Service{repo: repo, txManager: manager}
}
