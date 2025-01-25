package locationsrv

import (
	"context"
	"log"
	"math/rand/v2"
	"time"

	pb "github.com/akhenakh/goqtgrpc/gen/go/locationsvc/v1"
	"google.golang.org/grpc"
)

type Server struct {
	StopChan chan struct{}
}

func New() *Server {
	return &Server{}
}

func (s *Server) StreamPosition(req *pb.PositionRequest, srv grpc.ServerStreamingServer[pb.PositionResponse]) error {
	log.Println("got a stream request", req.DeviceId)
	timer := time.NewTicker(time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			resp := pb.PositionResponse{
				DeviceId: req.GetDeviceId(),
				//Timestamp: timestamppb.Now(),
				Longitude: float32(rand.Float32()*180)*2 - 180,
				Latitude:  float32(rand.Float32()*180) - 90,
			}

			if err := srv.Send(&resp); err != nil {
				log.Println("error sending response")
				return err
			}
		case <-s.StopChan:
			return nil
		}
	}
}

func (s *Server) Position(ctx context.Context, req *pb.PositionRequest) (*pb.PositionResponse, error) {
	log.Println("got a unary request", req.DeviceId)

	return &pb.PositionResponse{
		DeviceId: req.GetDeviceId(),
		//Timestamp: timestamppb.Now(),
		Longitude: float32(rand.Float32()*180)*2 - 180,
		Latitude:  float32(rand.Float32()*180) - 90,
	}, nil
}
