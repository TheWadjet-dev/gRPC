package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "path/to/generated/protobuf"
)

type server struct {
	pb.UnimplementedHelloWorldServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "¡Hola, " + req.GetName() + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloWorldServer(grpcServer, &server{})

	fmt.Println("Servidor gRPC en el puerto 3000")
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
