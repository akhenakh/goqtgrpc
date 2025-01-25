package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/akhenakh/goqtgrpc/gen/go/locationsvc/v1"
	"github.com/akhenakh/goqtgrpc/locationsrv"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:55010")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	s := &locationsrv.Server{}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterLocationServiceServer(grpcServer, s)
	grpcServer.Serve(lis)
}
