package chat

import (
	"context"

	"github.com/lookandhate/course_chat/internal/service/convertor"
	chatAPI "github.com/lookandhate/course_chat/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateChat(ctx context.Context, request *chatAPI.CreateChatRequest) (*chatAPI.CreateChatResponse, error) {
	id, err := s.chatService.CreateChat(ctx, convertor.CreateChatFromProto(request))
	if err != nil {
		return nil, err
	}

	return &chatAPI.CreateChatResponse{Id: int64(id)}, nil
}

func (s *Server) SendMessage(ctx context.Context, request *chatAPI.SendMessageRequest) (*emptypb.Empty, error) {
	err := s.chatService.SendMessage(ctx, convertor.SendMessageFromProto(request))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}
