package main

import (
	"log"
	"context"

	"google.golang.org/grpc"
	pb "github.com/kuuhaku86/simple-go-rpc/dataserver"
)

func main() {
	conn, err := grpc.Dial("localhost:8600", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataServerClient(conn)

	data, err := client.GetData(context.Background(), &pb.Request{Key: "a"})
	if err != nil {
		log.Fatalf("client.GetData failed: %v", err)
	} else {
		log.Printf("Result: %d", data.Value)
	}
	
	data, err = client.GetData(context.Background(), &pb.Request{Key: "b"})
	if err != nil {
		log.Fatalf("client.GetData failed: %v", err)
	} else {
		log.Printf("Result: %d", data.Value)
	}

	data, err = client.GetData(context.Background(), &pb.Request{Key: "c"})
	if err != nil {
		log.Fatalf("client.GetData failed: %v", err)
	} else {
		log.Printf("Result: %d", data.Value)
	}
}