package chat

import (
	"context"

	"github.com/lookandhate/course_chat/internal/service/convertor"
	chatAPI "github.com/lookandhate/course_chat/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Create(ctx context.Context, request *chatAPI.CreateRequest) (*chatAPI.CreateResponse, error) {
	id, err := s.chatService.Create(ctx, convertor.CreateChatFromProto(request))

	return &chatAPI.CreateResponse{Id: int64(id)}, err
}

func (s *Server) SendMessage(ctx context.Context, request *chatAPI.SendMessageRequest) (*emptypb.Empty, error) {
	err := s.chatService.SendMessage(ctx, convertor.SendMessageFromProto(request))

	return &emptypb.Empty{}, err
}
