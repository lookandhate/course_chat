package api_test

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
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	var (
		ctx = context.Background()

		mc        = minimock.NewController(t)
		userCount = gofakeit.Number(1, 100)
		users     = make([]int64, 0, userCount)
		chatID    = gofakeit.Int64()
	)
	// Generate list of user ids
	for range userCount {
		users = append(users, gofakeit.Int64())
	}

	req := chatAPI.CreateChatRequest{UserIds: users}
	info := model.CreateChat{
		UserIDs: users,
	}

	type args struct {
		ctx context.Context
		req *chatAPI.CreateChatRequest
	}

	tests := []struct {
		name            string
		args            args
		expectedResult  *chatAPI.CreateChatResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: &req,
			},
			expectedResult: &chatAPI.CreateChatResponse{
				Id: chatID,
			},
			err: nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateChatMock.Expect(ctx, &info).Return(int(chatID), nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewChatServer(chatServiceMock)

			newID, err := api.CreateChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.expectedResult, newID)
			require.Equal(t, tt.err, err)
		})
	}
}
