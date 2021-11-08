// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: fcd9ff140d
// Version Date: 2021-07-14T06:36:40Z

// Package grpc provides a gRPC client for the Auth service.
package grpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/Presbyter/services/protobuf"
	"github.com/Presbyter/services/src/service/auth-service/svc"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.AuthServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var getcodeEndpoint endpoint.Endpoint
	{
		getcodeEndpoint = grpctransport.NewClient(
			conn,
			"protobuf.Auth",
			"GetCode",
			EncodeGRPCGetCodeRequest,
			DecodeGRPCGetCodeResponse,
			pb.GetCodeResp{},
			clientOptions...,
		).Endpoint()
	}

	var gettokenEndpoint endpoint.Endpoint
	{
		gettokenEndpoint = grpctransport.NewClient(
			conn,
			"protobuf.Auth",
			"GetToken",
			EncodeGRPCGetTokenRequest,
			DecodeGRPCGetTokenResponse,
			pb.GetTokenResp{},
			clientOptions...,
		).Endpoint()
	}

	var refreshtokenEndpoint endpoint.Endpoint
	{
		refreshtokenEndpoint = grpctransport.NewClient(
			conn,
			"protobuf.Auth",
			"RefreshToken",
			EncodeGRPCRefreshTokenRequest,
			DecodeGRPCRefreshTokenResponse,
			pb.RefreshTokenResp{},
			clientOptions...,
		).Endpoint()
	}

	var bantokenEndpoint endpoint.Endpoint
	{
		bantokenEndpoint = grpctransport.NewClient(
			conn,
			"protobuf.Auth",
			"BanToken",
			EncodeGRPCBanTokenRequest,
			DecodeGRPCBanTokenResponse,
			pb.Empty{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		GetCodeEndpoint:      getcodeEndpoint,
		GetTokenEndpoint:     gettokenEndpoint,
		RefreshTokenEndpoint: refreshtokenEndpoint,
		BanTokenEndpoint:     bantokenEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCGetCodeResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getcode reply to a user-domain getcode response. Primarily useful in a client.
func DecodeGRPCGetCodeResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetCodeResp)
	return reply, nil
}

// DecodeGRPCGetTokenResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC gettoken reply to a user-domain gettoken response. Primarily useful in a client.
func DecodeGRPCGetTokenResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetTokenResp)
	return reply, nil
}

// DecodeGRPCRefreshTokenResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC refreshtoken reply to a user-domain refreshtoken response. Primarily useful in a client.
func DecodeGRPCRefreshTokenResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.RefreshTokenResp)
	return reply, nil
}

// DecodeGRPCBanTokenResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC bantoken reply to a user-domain bantoken response. Primarily useful in a client.
func DecodeGRPCBanTokenResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.Empty)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCGetCodeRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getcode request to a gRPC getcode request. Primarily useful in a client.
func EncodeGRPCGetCodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetCodeReq)
	return req, nil
}

// EncodeGRPCGetTokenRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain gettoken request to a gRPC gettoken request. Primarily useful in a client.
func EncodeGRPCGetTokenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetTokenReq)
	return req, nil
}

// EncodeGRPCRefreshTokenRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain refreshtoken request to a gRPC refreshtoken request. Primarily useful in a client.
func EncodeGRPCRefreshTokenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.RefreshTokenReq)
	return req, nil
}

// EncodeGRPCBanTokenRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain bantoken request to a gRPC bantoken request. Primarily useful in a client.
func EncodeGRPCBanTokenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.BanTokenReq)
	return req, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}