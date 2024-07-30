package chat

import (
	"github.com/lookandhate/course_chat/internal/service"
	chatAPI "github.com/lookandhate/course_chat/pkg/chat_v1"
)

type Server struct {
	chatAPI.UnimplementedChatServer
	chatService service.ChatService
}

// NewChatServer returns Server.
func NewChatServer(service service.ChatService) *Server {
	return &Server{
		chatService: service,
	}
}
