package handlers

import (
	"context"

	pb "github.com/Presbyter/services/protobuf"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.UserServer {
	return userService{}
}

type userService struct{}

func (s userService) UserInfo(ctx context.Context, in *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	var resp pb.UserInfoResp
	return &resp, nil
}
