package chat

import (
	"context"

	"github.com/lookandhate/course_chat/internal/api/convertor"
	chatAPI "github.com/lookandhate/course_chat/pkg/chat_v1"
)

func (s *Server) CreateChat(
	ctx context.Context,
	request *chatAPI.CreateChatRequest,
) (*chatAPI.CreateChatResponse, error) {
	id, err := s.chatService.CreateChat(ctx, convertor.CreateChatFromProto(request))
	if err != nil {
		return nil, err
	}

	return &chatAPI.CreateChatResponse{Id: int64(id)}, nil
}
