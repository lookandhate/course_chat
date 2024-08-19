package service_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/lookandhate/course_chat/internal/cache"
	cacheMocks "github.com/lookandhate/course_chat/internal/cache/mocks"
	"github.com/lookandhate/course_chat/internal/repository"
	repoMocks "github.com/lookandhate/course_chat/internal/repository/mocks"
	chatService "github.com/lookandhate/course_chat/internal/service/chat"
	"github.com/lookandhate/course_platform_lib/pkg/db"
	"github.com/lookandhate/course_platform_lib/pkg/db/mocks"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type chatRepoMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type chatCacheMockFunc func(mc *minimock.Controller) cache.ChatCache
	type txManagerMockFunc func(f func(context.Context) error, mc *minimock.Controller) db.TxManager

	var (
		mc            = minimock.NewController(t)
		chatID        = gofakeit.Uint32()
		ctx           = context.Background()
		deleteChatReq = chatID
	)
	type args struct {
		ctx context.Context
		req int
	}
	tests := []struct {
		name string
		args args
		err  error

		chatRepoMock  chatRepoMockFunc
		chatCacheMock chatCacheMockFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: int(deleteChatReq),
			},
			err: nil,
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, int64(deleteChatReq)).Return(nil)
				return mock
			},
			chatCacheMock: func(mc *minimock.Controller) cache.ChatCache {
				mock := cacheMocks.NewChatCacheMock(mc)
				mock.DeleteChatMock.Expect(ctx, int(deleteChatReq)).Return(nil)
				return mock
			},
			txManagerMock: func(_ func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := mocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Optional().Set(func(ctx context.Context, f db.Handler) error {
					return f(ctx)
				})
				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepoMock := tt.chatRepoMock(mc)
			chatCacheMock := tt.chatCacheMock(mc)
			txManagerMock := tt.txManagerMock(func(ctx context.Context) error {
				txErr := chatRepoMock.Delete(ctx, int64(deleteChatReq))
				if txErr != nil {
					return txErr
				}
				return nil
			}, mc)

			serviceTest := chatService.NewService(chatRepoMock, txManagerMock, chatCacheMock)
			err := serviceTest.DeleteChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
