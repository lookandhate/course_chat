package convertor

import (
	"database/sql"

	cacheModel "github.com/lookandhate/course_chat/internal/cache/model"
	repoModel "github.com/lookandhate/course_chat/internal/repository/model"
	serviceModel "github.com/lookandhate/course_chat/internal/service/model"
)

// CreateChatRequestToChatCreateRepo converts from service model to repo model.
func CreateChatRequestToChatCreateRepo(chat *serviceModel.CreateChat) *repoModel.CreateChatModel {
	return &repoModel.CreateChatModel{
		UserIDs: chat.UserIDs,
	}
}

// CreateMessageRequestToMessageCreateRepo converts from service chat creation to repo model.
func CreateMessageRequestToMessageCreateRepo(message *serviceModel.CreateMessage) *repoModel.CreateMessageModel {
	return &repoModel.CreateMessageModel{
		ChatID:   message.ChatID,
		Content:  message.Content,
		AuthorID: message.AuthorID,
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

// ServiceChatModelToCacheChatModel converts from service layer to cache layer.
func ServiceChatModelToCacheChatModel(chat *serviceModel.ChatModel) *cacheModel.ChatModel {
	var updatedAtNS int64
	if chat.UpdatedAt.Valid {
		updatedAtNS = chat.UpdatedAt.Time.UnixNano()
	} else {
		updatedAtNS = 0
	}
	return &cacheModel.ChatModel{
		ID:          chat.ChatID,
		UserIDs:     chat.UserIDs,
		CreatedAtNs: chat.CreatedAt.UnixNano(),
		UpdatedAtNs: updatedAtNS,
	}
}

func ServiceMessageModelToCacheMessageModel(message *serviceModel.MessageModel) *cacheModel.MessageModel {
	return &cacheModel.MessageModel{
		ID:          message.ID,
		Content:     message.Content,
		Author:      message.AuthorID,
		ChatID:      message.ChatID,
		TimestampNS: message.Timestamp.UnixNano(),
	}
}
