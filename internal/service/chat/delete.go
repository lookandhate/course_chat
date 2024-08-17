package service

import (
	"context"
	"log"
)

func (s Service) DeleteChat(ctx context.Context, chatID int) error {
	if err := s.validateID(ctx, chatID); err != nil {
		return err
	}

	err := s.cache.DeleteChat(ctx, chatID)
	if err != nil {
		log.Printf("failed to delete chat from cache %d: %v", chatID, err)
	}

	return s.repo.Delete(ctx, int64(chatID))
}
