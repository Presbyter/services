package handlers

import (
	"context"
	"net/url"

	pb "github.com/Presbyter/services/protobuf"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.ImServer {
	return ntesIMService{}
}

const (
	ntesBaseUrl = "https://api.netease.im"
)

type ntesIMService struct{}

func (s ntesIMService) CreateAccount(ctx context.Context, in *pb.CreateAccountReq) (*pb.CreateAccountResp, error) {
	var resp pb.CreateAccountResp
	addr := ntesBaseUrl + "/nimserver/user/create.action"
	data := url.Values{}
	return &resp, nil
}

func (s ntesIMService) AddGroup(ctx context.Context, in *pb.AddGroupReq) (*pb.AddGroupResp, error) {
	var resp pb.AddGroupResp
	return &resp, nil
}

func (s ntesIMService) DeleteGroup(ctx context.Context, in *pb.DeleteGroupReq) (*pb.Empty, error) {
	var resp pb.Empty
	return &resp, nil
}

func (s ntesIMService) AddFriend(ctx context.Context, in *pb.AddFriendReq) (*pb.Empty, error) {
	var resp pb.Empty
	return &resp, nil
}

func (s ntesIMService) DeleteFriend(ctx context.Context, in *pb.DeleteFriendReq) (*pb.Empty, error) {
	var resp pb.Empty
	return &resp, nil
}
