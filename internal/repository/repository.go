package repository

import (
	"context"

	"github.com/lookandhate/microservice-courese/chat/internal/repository/model"
)

type ChatRepository interface {
	CreateChat(context.Context, *model.CreateChatModel) (*model.ChatModel, error)
	CreateMessage(context.Context, *model.CreateMessageModel) (*model.MessageModel, error)
	Delete(context.Context, int64) error
	ChatExists(ctx context.Context, chatID int) (bool, error)
}
