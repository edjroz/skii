package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"regexp"

	pb "github.com/edjroz/skii/types/proto"
	"google.golang.org/grpc"

	traverse "github.com/edjroz/skii/graph/compute"
	"github.com/edjroz/skii/graph/types"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedGreeterServer
	Graph *types.Graph
}

// Routes - skii.Routes returns all possible routes a skier could take given a difficulty enum of (black, red, blue)
func (s *Server) Routes(ctx context.Context, in *pb.RouteRequest) (*pb.RouteReply, error) {
	if match, _ := regexp.MatchString("(black|red|blue)", in.GetColor()); !match {
		return &pb.RouteReply{}, errors.New("Requested color does not match ('black|red|blue')")
	}
	paths := traverse.GetAllPath(s.Graph, in.GetStart(), in.GetColor())
	log.Printf("Request: (start: %s, color: %s) => %+v", in.GetStart(), in.GetColor(), paths)
	routes := []*pb.RouteReply_Node{}
	for _, path := range paths {
		route := &pb.RouteReply_Node{Node: []string{}}
		route.Node = append(route.Node, path...)
		routes = append(routes, route)
	}
	return &pb.RouteReply{Route: routes}, nil
}

func Start(g *types.Graph, port string) {
	fmt.Println(port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{Graph: g})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
