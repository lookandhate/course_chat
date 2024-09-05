package app

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/lookandhate/course_chat/internal/interceptor"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lookandhate/course_chat/pkg/chat_v1"
	_ "github.com/lookandhate/course_chat/statik"
	"github.com/lookandhate/course_platform_lib/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
	swaggerServer   *http.Server
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
	wg.Add(3)

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

	go func() {
		defer wg.Done()

		err := a.runSwaggerServer()
		if err != nil {
			log.Fatalf("Failed to run swagger server %v", err)
		}
	}()

	wg.Wait()

	return nil
}

// initDeps initialize dependencies.
func (a *App) initDeps(ctx context.Context) error {
	initFuncs := []func(context.Context) error{

		a.initServiceProvider, a.initGRPCServer, a.initHTTPServer, a.initSwaggerServer,
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

func (a *App) initSwaggerServer(_ context.Context) error {
	statikFs, err := fs.New()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFs)))
	mux.HandleFunc("/api.swagger.json", serveSwaggerFile("/api.swagger.json"))

	a.swaggerServer = &http.Server{
		Addr:    a.serviceProvider.AppCfg().Swagger.Address(),
		Handler: mux,
	}

	return nil
}

func (a *App) runSwaggerServer() error {
	serveAddress := a.serviceProvider.AppCfg().Swagger.Address()
	log.Printf("Swagger Server is running on %s", serveAddress)

	err := a.swaggerServer.ListenAndServe()

	if err != nil {
		return err
	}

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

	corsMiddleware := cors.New(cors.Options{AllowedOrigins: []string{"*"}})

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.AppCfg().HTTP.Address(),
		Handler: corsMiddleware.Handler(mux),
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

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statikFs, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, err := statikFs.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(content)
	}
}
