package convertor

import (
	"github.com/lookandhate/course_chat/internal/service/model"
	"github.com/lookandhate/course_chat/pkg/chat_v1"
)

// CreateChatFromProto converts data from protobuf to service layer model.
func CreateChatFromProto(chat *chat_v1.CreateChatRequest) *model.CreateChat {
	return &model.CreateChat{
		UserIDs: chat.UserIds,
	}
}

// SendMessageFromProto converts message creation data from proto to service layer model.
func SendMessageFromProto(message *chat_v1.SendMessageRequest) *model.CreateMessage {
	return &model.CreateMessage{
		ChatID:    int(message.ChatId),
		AuthorID:  int(message.From),
		Content:   message.Text,
		Timestamp: message.Timestamp.AsTime(),
	}
}
