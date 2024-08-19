package api_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/lookandhate/course_chat/internal/api/chat"
	"github.com/lookandhate/course_chat/internal/service"
	serviceMocks "github.com/lookandhate/course_chat/internal/service/mocks"
	chatAPI "github.com/lookandhate/course_chat/pkg/chat_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	var (
		ctx = context.Background()

		mc     = minimock.NewController(t)
		chatID = gofakeit.Int64()
	)

	req := chatAPI.DeleteChatRequest{Id: chatID}
	info := chatID

	type args struct {
		ctx context.Context
		req *chatAPI.DeleteChatRequest
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
				mock.DeleteChatMock.Expect(ctx, int(info)).Return(nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewChatServer(chatServiceMock)

			newID, err := api.DeleteChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.expectedResult, newID)
			require.Equal(t, tt.err, err)
		})
	}
}
