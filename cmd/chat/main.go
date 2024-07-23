package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lookandhate/microservice-courese/chat/internal/config"
	"github.com/lookandhate/microservice-courese/chat/internal/grpc/chat"
	repository "github.com/lookandhate/microservice-courese/chat/internal/repository/chat"
	"github.com/lookandhate/microservice-courese/chat/internal/service/chat"
	"github.com/lookandhate/microservice-courese/chat/pkg/chat_v1"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.MustLoad()
	serverHost := fmt.Sprintf("localhost:%d", cfg.GPRC.Port)

	ctx := context.Background()

	connectionPool, err := pgxpool.New(ctx, cfg.DB.GetDSN())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Printf("Serving at %v", serverHost)

	lis, err := net.Listen("tcp", serverHost)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	repo := repository.NewPostgresRepository(connectionPool)
	server := service.NewService(repo)

	s := grpc.NewServer()
	chatServer := chat.NewChatServer(server)
	chat_v1.RegisterChatServer(s, chatServer)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC server %s", err)
	}
}
