package locationsrv

import (
	"context"
	"log"
	"time"

	pb "github.com/akhenakh/goqtgrpc/gen/go/locationsvc/v1"

	"github.com/akhenakh/sataas/sgp4"
	"google.golang.org/grpc"
)

const tle1 = `1 25544U 98067A   25025.00048859  .00033214  00000+0  57704-3 0  9999`
const tle2 = `2 25544  51.6377 296.2827 0003104 141.8447 313.9175 15.50506992492954`

type Server struct {
	prop     *sgp4.SGP4
	StopChan chan struct{}
}

func New() *Server {
	tle, _ := sgp4.NewTLE("ISS", tle1, tle2)
	prop, _ := sgp4.NewSGP4(tle)
	return &Server{prop: prop}
}

func (s *Server) StreamPosition(req *pb.PositionRequest, srv grpc.ServerStreamingServer[pb.PositionResponse]) error {
	log.Println("got a stream request", req.DeviceId)
	timer := time.NewTicker(time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			lat, lng, _, _ := s.prop.Position(time.Now())
			resp := pb.PositionResponse{
				DeviceId: req.GetDeviceId(),
				//Timestamp: timestamppb.Now(),
				Longitude: float32(lng),
				Latitude:  float32(lat),
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

	lat, lng, _, _ := s.prop.Position(time.Now())
	return &pb.PositionResponse{
		DeviceId: req.GetDeviceId(),
		//Timestamp: timestamppb.Now(),
		Longitude: float32(lng),
		Latitude:  float32(lat),
	}, nil
}
