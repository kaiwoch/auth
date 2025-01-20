package main

import (
	pb "github.com/KamigamiNoGigan/auth/pkg/user_api_v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
	"context"
	"math/rand"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedUserAPIServer
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{Id: rand.Int63n(100)}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{Id: 1, Name: "kaiwoch", Email: "kanzartem11@mail.ru", Role: 1, CreatedAt: timestamppb.New(time.Now()), UpdatedAt: timestamppb.New(time.Now())}, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateRequest) (*emptypb.Empty, error) {
	log.Println(in.Id, in.Name, in.Email)
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	log.Println(in.Id)
	return &emptypb.Empty{}, nil
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
