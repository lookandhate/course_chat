package chat

import (
	"context"

	chatAPI "github.com/lookandhate/course_chat/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Delete(ctx context.Context, request *chatAPI.DeleteRequest) (*emptypb.Empty, error) {
	err := s.chatService.Delete(ctx, int(request.Id))

	return &emptypb.Empty{}, err
}
