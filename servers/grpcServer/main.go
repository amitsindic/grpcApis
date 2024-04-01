package main

import (
	"fmt"
	"net"

	svc "example.com/practice/controllers/grpcController"
	pb "example.com/practice/proto/sample/proto"

	"google.golang.org/grpc"
)


func main() {

	fmt.Println("Starting grpc server")
	lis, err := net.Listen("tcp", ":50051")
	fmt.Printf("server listening at %v", lis.Addr())
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAvailableServicesServer(s, &svc.Server{})

	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}

}
