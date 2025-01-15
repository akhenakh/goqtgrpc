package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:0") // Port 0 means random available port
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	addr := lis.Addr().(*net.TCPAddr)
	port := addr.Port

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
