package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc.first.try/proto"
)

var testFeature = []byte(`[{
    "location": {
        "latitude": 407838351,
        "longitude": -746143763
    },
    "name": "Patriots Path, Mendham, NJ 07945, USA"
}`)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	testFeature := pb.Feature{
		Name: "Patriots Path, Mendham, NJ 07945, USA",
		Location: &pb.Point{
			Latitude:  10,
			Longitude: 21,
		},
	}
	return &testFeature, nil
}

func newServer() *routeGuideServer {
	s := &routeGuideServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterRouteGuideServer(srv, newServer())
	srv.Serve(lis)
}
