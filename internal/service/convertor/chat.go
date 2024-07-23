package convertor

import (
	"time"

	"github.com/lookandhate/microservice-courese/chat/internal/service/model"
	"github.com/lookandhate/microservice-courese/chat/pkg/chat_v1"
)

// CreateChatFromProto converts data from protobuf to service layer model.
func CreateChatFromProto(chat *chat_v1.CreateRequest) *model.CreateChatRequest {
	return &model.CreateChatRequest{
		UserIDs: chat.UserIds,
	}
}

// SendMessageFromProto converts message creation data from proto to service layer model.
func SendMessageFromProto(message *chat_v1.SendMessageRequest) *model.SendMessageRequest {
	return &model.SendMessageRequest{
		ChatID:    int(message.ChatId),
		AuthorID:  int(message.From),
		Content:   message.Text,
		Timestamp: time.Now(),
	}
}
