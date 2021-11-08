package handlers

import (
	"context"

	pb "github.com/Presbyter/services/protobuf"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.AuthServer {
	return authService{}
}

type authService struct{}

func (s authService) GetCode(ctx context.Context, in *pb.GetCodeReq) (*pb.GetCodeResp, error) {
	var resp pb.GetCodeResp
	return &resp, nil
}

func (s authService) GetToken(ctx context.Context, in *pb.GetTokenReq) (*pb.GetTokenResp, error) {
	var resp pb.GetTokenResp
	return &resp, nil
}

func (s authService) RefreshToken(ctx context.Context, in *pb.RefreshTokenReq) (*pb.RefreshTokenResp, error) {
	var resp pb.RefreshTokenResp
	return &resp, nil
}

func (s authService) BanToken(ctx context.Context, in *pb.BanTokenReq) (*pb.Empty, error) {
	var resp pb.Empty
	return &resp, nil
}
