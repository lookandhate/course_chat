package chat

import (
	"context"

	chatAPI "github.com/lookandhate/course_chat/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteChat(ctx context.Context, request *chatAPI.DeleteChatRequest) (*emptypb.Empty, error) {
	err := s.chatService.DeleteChat(ctx, int(request.Id))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
