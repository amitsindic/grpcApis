package main

import (
	"context"
	"log"

	pb "example.com/practice/proto/sample/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance.
	client := pb.NewAvailableServicesClient(conn)

	// Call the gRPC method.
	message := &pb.SampleMessage{Name: "John wick", Id: 1}
	
	response, err := client.GetUserDetail(context.Background(), message)
	if err != nil {
		log.Fatalf("Error calling GetUserDetail: %v", err)
	}

	// Print the response.
	log.Printf("Received user details: %s (ID: %d)", response.Name, response.Id)
}
