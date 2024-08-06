package service

import (
	"context"

	"github.com/lookandhate/course_chat/internal/repository/convertor"
	"github.com/lookandhate/course_chat/internal/service"
	"github.com/lookandhate/course_chat/internal/service/model"
)

// CreateChat creates chat with given users.
func (s Service) CreateChat(ctx context.Context, chat *model.CreateChatRequest) (int, error) {
	createdChat, err := s.repo.CreateChat(ctx, convertor.CreateChatRequestToChatCreateRepo(chat))
	if err != nil {
		return 0, err
	}

	return createdChat.ID, nil
}

// SendMessage sends message to the chat.
func (s Service) SendMessage(ctx context.Context, message *model.SendMessageRequest) error {
	if err := s.validateID(ctx, message.ChatID); err != nil {
		return err
	}

	if err := s.checkChatExists(ctx, message.ChatID); err != nil {
		return service.ErrChatDoesNotExist
	}
	// TODO add check that user has access to the chat

	_, err := s.repo.CreateMessage(ctx, convertor.CreateMessageRequestToMessageCreateRepo(message))

	return err
}
