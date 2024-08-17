package service

import (
	"context"
	"log"

	"github.com/lookandhate/course_chat/internal/service"
	"github.com/lookandhate/course_chat/internal/service/convertor"
	"github.com/lookandhate/course_chat/internal/service/model"
)

// CreateChat creates chat.go with given users.
func (s Service) CreateChat(ctx context.Context, chat *model.CreateChat) (int, error) {
	var createdChat *model.ChatModel
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error
		createdChatRepo, txErr := s.repo.CreateChat(ctx, convertor.CreateChatRequestToChatCreateRepo(chat))
		if txErr != nil {
			return txErr
		}
		createdChat = convertor.RepoChatModelToServiceChatModel(createdChatRepo)

		return nil
	})
	if err != nil {
		return 0, err
	}

	// Do not think that we need to raise cache error above, just log it
	err = s.cache.CreateChat(ctx, convertor.ServiceChatModelToCacheChatModel(createdChat))
	if err != nil {
		log.Default().Printf("Error when saving chat to cache: %v", err)
	}

	return createdChat.ChatID, nil
}

// SendMessage sends message to the chat.go.
func (s Service) SendMessage(ctx context.Context, message *model.CreateMessage) error {
	if err := s.validateID(ctx, message.ChatID); err != nil {
		return err
	}

	if err := s.checkChatExists(ctx, message.ChatID); err != nil {
		return service.ErrChatDoesNotExist
	}
	// TODO add check that user has access to the chat

	createdMessageRepo, err := s.repo.CreateMessage(ctx, convertor.CreateMessageRequestToMessageCreateRepo(message))
	if err != nil {
		return err
	}

	err = s.cache.CreateMessage(ctx, convertor.ServiceMessageModelToCacheMessageModel(convertor.RepoMessageModelToServiceMessageModel(createdMessageRepo)))
	if err != nil {
		log.Default().Printf("Error when saving message to cache: %v", err)
	}

	return nil
}
