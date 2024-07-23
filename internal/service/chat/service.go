package service

import (
	"github.com/lookandhate/microservice-courese/chat/internal/repository"
)

type Service struct {
	repo repository.ChatRepository
}

// NewService creates Service with given repo.
func NewService(repo repository.ChatRepository) *Service {
	return &Service{repo: repo}
}
