package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/ratelimit"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/restinbass/payment-service/cmd/app"
	"github.com/restinbass/payment-service/internal/config"
	"github.com/restinbass/payment-service/internal/interceptor"
	payment_v1 "github.com/restinbass/payment-service/pkg/proto/payment/v1"
	"github.com/restinbass/platform-libs/pkg/closer"
	"github.com/restinbass/platform-libs/pkg/grpc_healthcheck"
	"github.com/restinbass/platform-libs/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	cfg := config.Load(ctx)

	logger.Init(cfg.LoggerConfig)
	logger.Info(ctx, "logger initializaed successfully")

	closer.SetLogger(logger.Logger())
	closer.Configure(syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	var (
		repositories = app.InitRepositories(ctx, cfg.PostgresConfig)
		services     = app.InitServices(repositories)
		apis         = app.InitAPIs(services)
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GrpcServerConfig.Port()))
	if err != nil {
		logger.Fatal(ctx, "failed to listen tcp", zap.Int64("grpcPort", cfg.GrpcServerConfig.Port()), zap.Error(err))
	}
	closer.AddNamed("tcp listener", func(ctx context.Context) error {
		if cerr := lis.Close(); cerr != nil {
			logger.Error(ctx, "failed to close listener", zap.Error(cerr))
			return err
		}

		return nil
	})

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc.UnaryServerInterceptor(interceptor.LoggerInterceptor()),
			validator.UnaryServerInterceptor(),
			recovery.UnaryServerInterceptor(),
			ratelimit.UnaryServerInterceptor(interceptor.NewRateLimiter(100, 200)),
		),
	)

	grpc_healthcheck.RegisterHealthcheckService(grpcServer)
	payment_v1.RegisterPaymentServiceServer(grpcServer, apis.Payment)
	reflection.Register(grpcServer)

	go func() {
		logger.Info(ctx, "gRPC server is listening", zap.Int64("grpcPort", cfg.GrpcServerConfig.Port()))
		if err = grpcServer.Serve(lis); err != nil {
			logger.Fatal(ctx, "failed to grpcServer.Serve()", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info(ctx, "shutting down gRPC server...")
	grpcServer.GracefulStop()
	logger.Info(ctx, "gRPC server stopped")
}
