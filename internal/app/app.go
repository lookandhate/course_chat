package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lookandhate/course_chat/internal/interceptor"
	"github.com/lookandhate/course_chat/pkg/chat_v1"
	"github.com/lookandhate/course_platform_lib/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
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

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := a.runGRPCServer()
		if err != nil {
			log.Fatalf("Failed to run grpc server %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("Failed to run HTTP server %v", err)
		}
	}()

	wg.Wait()

	return nil
}

// initDeps initialize dependencies.
func (a *App) initDeps(ctx context.Context) error {
	initFuncs := []func(context.Context) error{

		a.initServiceProvider, a.initGRPCServer, a.initHTTPServer,
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

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := chat_v1.RegisterChatHandlerFromEndpoint(ctx, mux, a.serviceProvider.AppCfg().GPRC.Address(), opts)
	if err != nil {
		return err
	}

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.AppCfg().HTTP.Address(),
		Handler: mux,
	}

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)

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

func (a *App) runHTTPServer() error {
	serveAddress := a.serviceProvider.AppCfg().HTTP.Address()
	log.Printf("HTTP Server is running on %s", serveAddress)

	err := a.httpServer.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
