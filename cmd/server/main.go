package main

import (
	"flag"
	"github.com/bpalermo/conflicting-package-heights/pkg/server"
	"log"
)

var (
	logger             *log.Logger
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func init() {
	flag.Parse()
}

func main() {
	if err := server.Run(*grpcServerEndpoint); err != nil {
		logger.Fatal(err)
	}
}
