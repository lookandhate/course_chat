package service

import (
	"context"
	"log"

	"github.com/lookandhate/course_chat/internal/service"
	"github.com/lookandhate/course_chat/internal/service/model"
)

// CreateChat creates chat.go with given users.
func (s Service) CreateChat(ctx context.Context, chat *model.CreateChat) (int, error) {
	var createdChat *model.ChatModel
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error
		createdChat, txErr = s.repo.CreateChat(ctx, chat)
		if txErr != nil {
			return txErr
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	// Do not think that we need to raise cache error above, just log it
	err = s.cache.CreateChat(ctx, createdChat)
	if err != nil {
		log.Default().Printf("Error when saving chat to cache: %v", err)
	}

	return createdChat.ChatID, nil
}

// CreateMessage creates message in chat.
func (s Service) CreateMessage(ctx context.Context, message *model.CreateMessage) error {
	if err := s.validateID(ctx, message.ChatID); err != nil {
		return err
	}

	if err := s.checkChatExists(ctx, message.ChatID); err != nil {
		return service.ErrChatDoesNotExist
	}
	// TODO add check that user has access to the chat

	createdMessage, err := s.repo.CreateMessage(ctx, message)
	if err != nil {
		return err
	}

	err = s.cache.CreateMessage(ctx, createdMessage)

	if err != nil {
		log.Default().Printf("Error when saving message to cache: %v", err)
	}

	return nil
}
