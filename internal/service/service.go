package service

import (
	"context"

	"github.com/lookandhate/course_chat/internal/service/model"
)

type ChatService interface {
	CreateChat(context.Context, *model.CreateChatRequest) (int, error)
	DeleteChat(context.Context, int) error
	SendMessage(context.Context, *model.SendMessageRequest) error
}
