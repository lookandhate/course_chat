package convertor

import (
	"database/sql"

	repoModel "github.com/lookandhate/course_chat/internal/repository/model"
	serviceModel "github.com/lookandhate/course_chat/internal/service/model"
)

// ServiceCreateChatToRepoCreateChat converts from service model to repo model.
func ServiceCreateChatToRepoCreateChat(chat *serviceModel.CreateChat) *repoModel.CreateChatModel {
	return &repoModel.CreateChatModel{
		UserIDs: chat.UserIDs,
	}
}

// ServiceCreateMessageToRepoMessageCreate converts from service chat creation to repo model.
func ServiceCreateMessageToRepoMessageCreate(message *serviceModel.CreateMessage) *repoModel.CreateMessageModel {
	return &repoModel.CreateMessageModel{
		ChatID:    message.ChatID,
		Content:   message.Content,
		AuthorID:  message.AuthorID,
		Timestamp: message.Timestamp,
	}
}

func RepoChatModelToServiceChatModel(chat *repoModel.ChatModel) *serviceModel.ChatModel {
	return &serviceModel.ChatModel{
		UserIDs:   chat.UserIDs,
		ChatID:    chat.ID,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: sql.NullTime{Time: chat.UpdatedAt.Time, Valid: false},
	}
}

func RepoMessageModelToServiceMessageModel(message *repoModel.MessageModel) *serviceModel.MessageModel {
	return &serviceModel.MessageModel{
		ID:        message.ID,
		ChatID:    message.ChatID,
		AuthorID:  message.Author,
		Content:   message.Content,
		Timestamp: message.CreatedAt,
	}
}
