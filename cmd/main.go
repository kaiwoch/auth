package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "github.com/KamigamiNoGigan/auth/pkg/user_api_v1"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedUserAPIServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterUserAPIServer(s, &server{})
	
	log.Printf("server listening at %v", port)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}