package convertor

import (
	repoModel "github.com/lookandhate/microservice-courese/chat/internal/repository/model"
	"github.com/lookandhate/microservice-courese/chat/internal/service/model"
)

// CreateChatRequestToChatCreateRepo converts from service model to repo model.
func CreateChatRequestToChatCreateRepo(chat *model.CreateChatRequest) *repoModel.CreateChatModel {
	return &repoModel.CreateChatModel{
		UserIDs: chat.UserIDs,
	}
}

// CreateMessageRequestToMessageCreateRepo converts from service chat creation to repo model.
func CreateMessageRequestToMessageCreateRepo(message *model.SendMessageRequest) *repoModel.CreateMessageModel {
	return &repoModel.CreateMessageModel{
		ChatID:   message.ChatID,
		Content:  message.Content,
		AuthorID: message.AuthorID,
	}
}
