package service_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/lookandhate/course_chat/internal/cache"
	cacheMocks "github.com/lookandhate/course_chat/internal/cache/mocks"
	"github.com/lookandhate/course_chat/internal/repository"
	repoMocks "github.com/lookandhate/course_chat/internal/repository/mocks"
	chatService "github.com/lookandhate/course_chat/internal/service/chat"
	"github.com/lookandhate/course_chat/internal/service/model"
	"github.com/lookandhate/course_platform_lib/pkg/db"
	"github.com/lookandhate/course_platform_lib/pkg/db/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreateChat(t *testing.T) {
	t.Parallel()

	type chatRepoMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type chatCacheMockFunc func(mc *minimock.Controller) cache.ChatCache
	type txManagerMockFunc func(f func(context.Context) error, mc *minimock.Controller) db.TxManager

	var (
		mc = minimock.NewController(t)

		userCount = gofakeit.Number(1, 100)
		users     = make([]int64, 0, userCount)
		chatID    = gofakeit.Int64()
		ctx       = context.Background()
		timeNow   = time.Now()
	)

	for range userCount {
		users = append(users, int64(gofakeit.Uint32()))
	}
	createChatReq := &model.CreateChat{
		UserIDs: users,
	}
	createdChatRes := model.ChatModel{
		UserIDs:   users,
		ChatID:    int(chatID),
		CreatedAt: timeNow,
		UpdatedAt: sql.NullTime{
			Valid: false,
			Time:  time.Time{},
		},
	}

	type args struct {
		ctx context.Context
		req *model.CreateChat
	}

	tests := []struct {
		name           string
		args           args
		expectedResult int
		err            error

		chatRepoMock  chatRepoMockFunc
		chatCacheMock chatCacheMockFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name:           "success",
			expectedResult: int(chatID),
			args: args{
				ctx: context.Background(),
				req: createChatReq,
			},
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx, createChatReq).Return(&createdChatRes, nil)
				return mock
			},
			chatCacheMock: func(mc *minimock.Controller) cache.ChatCache {
				mock := cacheMocks.NewChatCacheMock(mc)
				mock.CreateChatMock.Expect(ctx, &createdChatRes).Return(nil)
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
				var txErr error
				createdChat, txErr := chatRepoMock.CreateChat(ctx, tt.args.req)
				if txErr != nil {
					return txErr
				}
				chatID = int64(createdChat.ID)
				return nil
			}, mc)

			serviceTest := chatService.NewService(chatRepoMock, txManagerMock, chatCacheMock)
			chat, err := serviceTest.CreateChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.expectedResult, chat)
		})
	}
}

func TestCreateMessage(t *testing.T) {
	t.Parallel()

	type chatRepoMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type chatCacheMockFunc func(mc *minimock.Controller) cache.ChatCache
	type txManagerMockFunc func(f func(context.Context) error, mc *minimock.Controller) db.TxManager

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		chatID    = gofakeit.Uint32()
		message   = gofakeit.Letter()
		authorID  = gofakeit.Int64()
		timestamp = gofakeit.Date()

		messageID        = gofakeit.Int64()
		createMessageReq = &model.CreateMessage{
			ChatID:    int(chatID),
			AuthorID:  int(authorID),
			Content:   message,
			Timestamp: timestamp,
		}
		createMessageRes = &model.MessageModel{
			ID:        int(messageID),
			ChatID:    int(chatID),
			AuthorID:  int(authorID),
			Content:   message,
			Timestamp: timestamp,
		}
	)
	type args struct {
		ctx context.Context
		req *model.CreateMessage
	}
	tests := []struct {
		name           string
		args           args
		expectedResult int
		err            error
		chatRepoMock   chatRepoMockFunc
		chatCacheMock  chatCacheMockFunc
		txManagerMock  txManagerMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: createMessageReq,
			},
			expectedResult: int(messageID),
			err:            nil,
			chatRepoMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.ChatExistsMock.Expect(ctx, int(chatID)).Return(true, nil)
				mock.CreateMessageMock.Expect(ctx, createMessageReq).Return(createMessageRes, nil)
				return mock
			},
			chatCacheMock: func(mc *minimock.Controller) cache.ChatCache {
				mock := cacheMocks.NewChatCacheMock(mc)
				mock.CreateMessageMock.Expect(ctx, createMessageRes).Return(nil)
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
				var txErr error
				createdChat, txErr := chatRepoMock.CreateMessage(ctx, createMessageReq)
				if txErr != nil {
					return txErr
				}
				chatID = uint32(createdChat.ID)
				return nil
			}, mc)

			serviceTest := chatService.NewService(chatRepoMock, txManagerMock, chatCacheMock)
			err := serviceTest.CreateMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
