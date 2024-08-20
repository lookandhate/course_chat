package cache

import (
	"context"

	"github.com/lookandhate/course_chat/internal/service/model"
)

type ChatCache interface {
	CreateChat(ctx context.Context, model *model.ChatModel) error
	CreateMessage(ctx context.Context, model *model.MessageModel) error
	DeleteMessage(ctx context.Context, id int) error
	DeleteChat(ctx context.Context, id int) error
}
