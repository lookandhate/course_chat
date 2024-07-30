package app

import (
	"context"
	"log"
	"net"

	"github.com/lookandhate/course_chat/internal/closer"
	"github.com/lookandhate/course_chat/pkg/chat_v1"
	_ "github.com/lookandhate/course_chat/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

// initDeps initialize dependencies.
func (a *App) initDeps(ctx context.Context) error {
	initFuncs := []func(context.Context) error{

		a.initServiceProvider, a.initGRPCServer,
	}
	for _, f := range initFuncs {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	chat_v1.RegisterChatServer(a.grpcServer, a.serviceProvider.ChatServerImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	serveAddress := a.serviceProvider.AppCfg().GPRC.Address()
	log.Printf("GRPC server is running on %s", serveAddress)

	listener, err := net.Listen("tcp", serveAddress)
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
