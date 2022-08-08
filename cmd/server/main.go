package main

import (
	"context"
	"log"
	"net"
	"teste/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedMessageServiceServer
}

func (service *Server) CreateMessage(ctx context.Context, req *pb.MessageInput) (*pb.Message, error) {
	log.Print(req.GetMessage())

	response := &pb.Message{
		Message: req.Message,
	}
	return response, nil
}
func (service *Server) mustEmbedUnimplementedMessageServiceServer() {}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterMessageServiceServer(grpcServer, &Server{})

	port := ":5000"

	listenner, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	reflection.Register(grpcServer)
	grpc_Error := grpcServer.Serve(listenner)

	if grpc_Error != nil {
		log.Fatal(grpc_Error)
	}
}
