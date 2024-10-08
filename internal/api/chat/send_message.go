package chat

import (
	"context"

	"github.com/lookandhate/course_chat/internal/api/convertor"
	chatAPI "github.com/lookandhate/course_chat/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) SendMessage(ctx context.Context, request *chatAPI.SendMessageRequest) (*emptypb.Empty, error) {
	err := s.chatService.CreateMessage(ctx, convertor.SendMessageFromProto(request))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}
