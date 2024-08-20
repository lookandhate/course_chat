package repository

import (
	"context"

	"github.com/lookandhate/course_chat/internal/service/model"
)

type ChatRepository interface {
	CreateChat(context.Context, *model.CreateChat) (*model.ChatModel, error)
	CreateMessage(context.Context, *model.CreateMessage) (*model.MessageModel, error)
	Delete(context.Context, int64) error
	ChatExists(ctx context.Context, chatID int) (bool, error)
}
