package controller

import (
	"context"
	"fmt"
	pb "example.com/practice/proto/sample/proto"
	service "example.com/practice/services"
)
 
type Server struct {
	pb.UnimplementedAvailableServicesServer
}

func (s *Server) GetUserDetail(ctx context.Context, message *pb.SampleMessage) (*pb.UserDetails, error) {
	fmt.Printf("incoming request %v \n", message)
	resp := service.GetUserDetails(message.Name, message.Id)
	var id int32
	if val,ok := resp["id"].(int32); ok {
		id = val
	}else{
		id = 0
	}
	return &pb.UserDetails{Name: fmt.Sprintf("%s",resp["name"]), Id:  id }, nil
}