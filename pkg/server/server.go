package server

import (
	"context"
	"fmt"
	"github.com/bpalermo/conflicting-package-heights/api/service/echo"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type server struct {
	echo.UnimplementedApiGatewayServer
}

func (g *server) Echo(_ context.Context, request *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{
		Value: fmt.Sprintf("hello %s", request.Value),
	}, nil
}

func Run(grpcServerEndpoint string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a listener on TCP port
	l, err := net.Listen("tcp", grpcServerEndpoint)
	if err != nil {
		return err
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	echo.RegisterApiGatewayServer(s, &server{})
	// Serve gRPC server
	log.Printf("Serving gRPC on %s", grpcServerEndpoint)
	go func() {
		log.Fatalln(s.Serve(l))
	}()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = echo.RegisterApiGatewayHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	gwServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8081")

	return gwServer.ListenAndServe()
}
