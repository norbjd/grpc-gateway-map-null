package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"

	gw "github.com/norbjd/grpc-gateway-map/gen/go/example"
)

const (
	grpcServerPort = 9090
	httpServerPort = 8081
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterEchoerHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", grpcServerPort), opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf("localhost:%d", httpServerPort), mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
