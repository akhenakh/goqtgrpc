package locationsrv

import (
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

func (s *Server) Position(req *pb.PositionRequest, srv grpc.ServerStreamingServer[pb.PositionResponse]) error {
	timer := time.NewTicker(time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			resp := pb.PositionResponse{
				Id: req.GetId(),
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
