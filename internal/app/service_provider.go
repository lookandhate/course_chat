package app

import (
	"context"
	"fmt"
	"log"

	chatServer "github.com/lookandhate/course_chat/internal/api/chat"
	"github.com/lookandhate/course_platform_lib/pkg/closer"
	"github.com/lookandhate/course_platform_lib/pkg/db"
	"github.com/lookandhate/course_platform_lib/pkg/db/pg"
	"github.com/lookandhate/course_platform_lib/pkg/db/transaction"

	"github.com/lookandhate/course_chat/internal/config"
	"github.com/lookandhate/course_chat/internal/repository"
	chatRepo "github.com/lookandhate/course_chat/internal/repository/chat"
	"github.com/lookandhate/course_chat/internal/service"
	chatService "github.com/lookandhate/course_chat/internal/service/chat"
)

// serviceProvider is a DI container.
type serviceProvider struct {
	appCfg *config.AppConfig

	dbClient           db.Client
	transactionManager db.TxManager

	chatRepository repository.ChatRepository

	chatService    service.ChatService
	chatServerImpl *chatServer.Server
}

// newServiceProvider creates plain serviceProvider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// AppCfg returns config.AppConfig.
func (s *serviceProvider) AppCfg() *config.AppConfig {
	if s.appCfg == nil {
		s.appCfg = config.MustLoad()
	}

	return s.appCfg
}

// ChatRepository creates(if not exist) and returns repository.ChatRepository instance.
func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepo.NewPostgresRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

// ChatService creates and returns service.ChatService.
func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.ChatRepository(ctx), s.TxManager(ctx))
	}

	return s.chatService
}

// ChatServerImpl returns GRPC implementation of the server.
func (s *serviceProvider) ChatServerImpl(ctx context.Context) *chatServer.Server {
	if s.chatServerImpl == nil {
		s.chatServerImpl = chatServer.NewChatServer(s.ChatService(ctx))
	}
	return s.chatServerImpl
}

// DBClient returns db.Client object.
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.AppCfg().DB.GetDSN())
		fmt.Printf("dsn: %s\n%v", s.AppCfg().DB.GetDSN(), cl)
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// TxManager returns transaction manager.
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.transactionManager == nil {
		s.transactionManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.transactionManager
}
