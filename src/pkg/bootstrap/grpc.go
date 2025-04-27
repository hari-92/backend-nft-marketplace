package bootstrap

import (
	"context"
	"fmt"
	"github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib/config"
	"github.com/golibs-starter/golib/log"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"net"
)

func WithGrpcServer() fx.Option {
	return fx.Options(
		golib.ProvideProps(NewGrpcServerProperties),
		fx.Provide(NewGrpcServer),
		// TODO: Register task execute
		fx.Invoke(func(lc fx.Lifecycle, server *GrpcServer) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						err := server.Start()
						if err != nil {
							log.Fatalf("Failed to start gRPC server: %v", err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					server.Stop()
					return nil
				},
			})
		}),
	)
}

type GrpcServer struct {
	Listener net.Listener
	Server   *grpc.Server
}

func (g *GrpcServer) Start() error {
	log.Info("[Boostrap GRPC] Starting gRPC server... ", g.Listener.Addr().String())
	err := g.Server.Serve(g.Listener)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return err
	}
	log.Info("[Boostrap GRPC] Listening gRPC server at ", g.Listener.Addr().String())

	return nil
}

func (g *GrpcServer) Stop() {
	g.Server.GracefulStop()
}

type GrpcServerProperties struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func (g GrpcServerProperties) Prefix() string {
	// Defined with module/config/.yml|yaml
	return "app.grpcServer"
}

func NewGrpcServer(p *GrpcServerProperties) (*GrpcServer, error) {
	fmt.Println("NewGrpcServer grpc server")
	address := fmt.Sprintf("%s:%d", p.Host, p.Port)
	log.Infof("NewGrpcServer grpc server address: %s", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("NewGrpcServer failed to listen: %v", err)
		return nil, err
	}

	server := grpc.NewServer()
	return &GrpcServer{
		Listener: listener,
		Server:   server,
	}, nil
}

func NewGrpcServerProperties(loader config.Loader) (*GrpcServerProperties, error) {
	props := GrpcServerProperties{}
	err := loader.Bind(&props)
	return &props, err
}
