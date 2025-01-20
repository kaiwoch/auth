package main

import (
	"google.golang.org/grpc"
	"context"
	"time"
	"log"
	pb "github.com/KamigamiNoGigan/auth/pkg/user_api_v1"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := c.Create(ctx, &pb.CreateRequest{Name: "kaiwoch", Password: "q1w2e3", PasswordConfirm: "q1w2e3", Role: 1})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id)

	u, err := c.Get(ctx, &pb.GetRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(u)

	_, err = c.Delete(ctx, &pb.DeleteRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.Update(ctx, &pb.UpdateRequest{Id: 1, Name: &wrapperspb.StringValue{Value: "dandadan"}, Email: &wrapperspb.StringValue{Value: "ddd@hotmail.com"}})
	if err != nil {
		log.Fatal(err)
	}
}