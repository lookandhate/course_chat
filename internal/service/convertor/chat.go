package convertor

import (
	cacheModel "github.com/lookandhate/course_chat/internal/cache/model"
	serviceModel "github.com/lookandhate/course_chat/internal/service/model"
)

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
