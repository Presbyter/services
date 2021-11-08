// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: fcd9ff140d
// Version Date: 2021-07-14T06:36:40Z

// Package grpc provides a gRPC client for the User service.
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
	"github.com/Presbyter/services/src/service/gateway-service/svc"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.UserServer, error) {
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
	var userinfoEndpoint endpoint.Endpoint
	{
		userinfoEndpoint = grpctransport.NewClient(
			conn,
			"protobuf.User",
			"UserInfo",
			EncodeGRPCUserInfoRequest,
			DecodeGRPCUserInfoResponse,
			pb.UserInfoResp{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		UserInfoEndpoint: userinfoEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCUserInfoResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC userinfo reply to a user-domain userinfo response. Primarily useful in a client.
func DecodeGRPCUserInfoResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.UserInfoResp)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCUserInfoRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain userinfo request to a gRPC userinfo request. Primarily useful in a client.
func EncodeGRPCUserInfoRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UserInfoReq)
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
