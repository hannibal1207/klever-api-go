package main

import (
	"context"
	"fmt"
	"log"
	"teste/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewMessageServiceClient(conn)

	var msg string
	fmt.Scan(&msg)

	req := &pb.MessageInput{
		Message: msg,
	}

	res, err := client.CreateMessage(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(res.GetMessage())
}
