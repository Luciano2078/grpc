package services

import (
	"context"
	"fmt"
	"github.com/Luciano2078/grpc/pb/pb"

)


type UserService struct {
	pb.UnimplementedUserServiceServer

}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return &pb.User {
		id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}