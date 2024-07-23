package chat

import (
	"context"

	"github.com/lookandhate/microservice-courese/chat/internal/service"
	"github.com/lookandhate/microservice-courese/chat/internal/service/convertor"
	chatAPI "github.com/lookandhate/microservice-courese/chat/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	chatAPI.UnimplementedChatServer
	chatService service.ChatService
}

func (s *Server) Create(ctx context.Context, request *chatAPI.CreateRequest) (*chatAPI.CreateResponse, error) {
	id, err := s.chatService.Create(ctx, convertor.CreateChatFromProto(request))

	return &chatAPI.CreateResponse{Id: int64(id)}, err
}

func (s *Server) Delete(ctx context.Context, request *chatAPI.DeleteRequest) (*emptypb.Empty, error) {
	err := s.chatService.Delete(ctx, int(request.Id))

	return &emptypb.Empty{}, err
}

func (s *Server) SendMessage(ctx context.Context, request *chatAPI.SendMessageRequest) (*emptypb.Empty, error) {
	err := s.chatService.SendMessage(ctx, convertor.SendMessageFromProto(request))

	return &emptypb.Empty{}, err
}

// NewChatServer returns GRPC server.
func NewChatServer(chatService service.ChatService) *Server {
	return &Server{chatService: chatService}
}
