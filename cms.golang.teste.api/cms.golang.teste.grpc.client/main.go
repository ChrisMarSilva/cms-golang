package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ChrisMarSilva/cms.golang.teste.grpc.client/proto-grpc"
	"google.golang.org/grpc"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.grpc.client
// go get -u google.golang.org/grpc
// go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
// go mod tidy

// go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
// protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto
// protoc --proto_path=src --go_opt=proto-grpc/configuracion.proto=cms.golang.teste.grpc.client/proto-grpc/configuracion.pb.go protos/buzz.proto

// protoc --go-grpc_out=./proto-grpc --go-grpc_opt=proto-grpc/configuracion.proto=cms.golang.teste.grpc.client/proto-grpc/configuracion.pb.go configuracion.proto
// protoc --go_out=. --go_opt=proto-grpc/configuracion.proto=cms.golang.teste.grpc.client/proto-grpc/ configuracion.proto
// protoc --go_out=plugins=grpc:. ~.proto
// protoc --go_out=. *.proto

// OK
// https://github.com/grpc/grpc-go/tree/master/examples
// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto-grpc/helloworld.proto

// go run main.go

func main() {
	log.Println("grpc.client.ini")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	//conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("#1 - No se puede conectar con el server :c (%v)", err)
	}

	defer conn.Close()

	//cl := pb.NewGetInfoClient(conn)
	cl := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	//ret, err := cl.ReturnInfo(ctx, &pb.RequestId{Id: "132"})
	ret, err := cl.SayHello(ctx, &pb.HelloRequest{Name: "Chris"})
	if err != nil {
		log.Fatalf("#2 - No se puede retornar la información :c (%v)", err)
	}

	//log.Printf("#3 - Respuesta del server: %s\n", ret.GetInfo())
	log.Printf("#3 - Respuesta del server: %s\n", ret.GetMessage())

	log.Println("grpc.client.fim")
}
