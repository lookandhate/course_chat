package api_tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/lookandhate/course_chat/internal/api/chat"
	"github.com/lookandhate/course_chat/internal/service"
	serviceMocks "github.com/lookandhate/course_chat/internal/service/mocks"
	"github.com/lookandhate/course_chat/internal/service/model"
	chatAPI "github.com/lookandhate/course_chat/pkg/chat_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	var (
		ctx = context.Background()

		mc          = minimock.NewController(t)
		userID      = gofakeit.Int64()
		chatID      = gofakeit.Int64()
		messageText = gofakeit.Letter()
		now         = gofakeit.Date()
	)

	req := chatAPI.SendMessageRequest{
		From:      userID,
		Text:      messageText,
		Timestamp: timestamppb.New(now),
		ChatId:    chatID,
	}
	info := model.CreateMessage{
		ChatID:    int(chatID),
		AuthorID:  int(userID),
		Content:   messageText,
		Timestamp: now,
	}

	type args struct {
		ctx context.Context
		req *chatAPI.SendMessageRequest
	}

	tests := []struct {
		name            string
		args            args
		expectedResult  *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: &req,
			},
			expectedResult: &emptypb.Empty{},
			err:            nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, &info).Return(nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewChatServer(chatServiceMock)

			newID, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.expectedResult, newID)
			require.Equal(t, tt.err, err)
		})
	}
}
