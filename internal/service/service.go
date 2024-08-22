package service

import (
	"context"

	"github.com/lookandhate/course_chat/internal/service/model"
)

type ChatService interface {
	CreateChat(context.Context, *model.CreateChat) (int, error)
	DeleteChat(context.Context, int) error
	CreateMessage(context.Context, *model.CreateMessage) error
}
