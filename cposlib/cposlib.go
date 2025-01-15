package main

import "C"

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/akhenakh/goqtgrpc/gen/go/locationsvc/v1"
	"github.com/akhenakh/goqtgrpc/locationsrv"
)

var cs *CServer

type CServer struct {
	gServer *grpc.Server
	*locationsrv.Server
}

// Start starts the listening grpc server and
//
//	returns the tcp port where is listening to
//	0 in case of error
//
//export Start
func Start() int {
	log.Println("start cposlib")
	cs = &CServer{
		Server: locationsrv.New(),
	}

	lis, err := net.Listen("tcp", "localhost:0") // Port 0 means random available port
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return 0
	}

	addr := lis.Addr().(*net.TCPAddr)
	port := addr.Port

	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterLocationServiceServer(grpcServer, cs)
		cs.gServer = grpcServer
		grpcServer.Serve(lis)
	}()

	return port
}

// Stop end all goroutine and server
//
//export Stop
func Stop() {
	log.Println("stop cposlib")

	if cs != nil {
		cs.gServer.Stop()
	}
	cs.Server.StopChan <- struct{}{}
}

func main() {}
