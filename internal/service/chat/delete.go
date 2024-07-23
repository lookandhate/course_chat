package service

import (
	"context"
)

func (s Service) Delete(ctx context.Context, chatID int) error {
	if err := s.validateID(ctx, chatID); err != nil {
		return err
	}
	return s.repo.Delete(ctx, int64(chatID))
}
