package service

import (
	"context"

	"github.com/lookandhate/course_chat/internal/service"
)

// checkChatExists - checks if chat.go exists and returns err if it does not exist.
func (s Service) checkChatExists(ctx context.Context, chatID int) error {
	isExists, err := s.repo.ChatExists(ctx, chatID)
	if err != nil {
		return err
	}
	if !isExists {
		return service.ErrChatDoesNotExist
	}

	return nil
}

// validateID validates given chat.go id.
func (s Service) validateID(_ context.Context, chatID int) error {
	if chatID <= 0 {
		return service.ErrInvalidID
	}

	return nil
}
