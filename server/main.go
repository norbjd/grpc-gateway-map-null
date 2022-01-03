package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "github.com/norbjd/grpc-gateway-map/gen/go/example"
)

const port = 9090

type Server struct {
	pb.UnimplementedEchoerServer
}

func (s *Server) Echo(_ context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{AMap: in.AMap}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterEchoerServer(grpcServer, new(Server))
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
