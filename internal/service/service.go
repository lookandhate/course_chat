package service

import (
	"context"

	"github.com/lookandhate/microservice-courese/chat/internal/service/model"
)

type ChatService interface {
	Create(context.Context, *model.CreateChatRequest) (int, error)
	Delete(context.Context, int) error
	SendMessage(context.Context, *model.SendMessageRequest) error
}
