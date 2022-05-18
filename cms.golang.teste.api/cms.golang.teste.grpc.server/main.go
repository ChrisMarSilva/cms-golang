package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ChrisMarSilva/cms.golang.teste.grpc.server/proto-grpc"
	"google.golang.org/grpc"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.grpc.server
// go get -u google.golang.org/grpc
// go mod tidy

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto-grpc/helloworld.proto

// go run main.go

func main() {
	log.Println("grpc.server.fim")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("#1- Fallo al levantar el servidor: %v", err)
	}

	s := grpc.NewServer()

	//pb.RegisterGetInfoServer(s, &server{})
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("#2- server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("#3- failed to serve: %v", err)
	}

	log.Println("grpc.server.fim")
}

type server struct {
	//pb.UnimplementedGetInfoServer
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
