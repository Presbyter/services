package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

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
	httpResp, err := http.PostForm(addr, data)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}
	if val, ok := m["code"].(float64); ok && val != 200 {
		return nil, fmt.Errorf(strconv.Itoa(int(val)))
	}

	resp.Accid = m["info"].(map[string]interface{})["accid"].(string)
	resp.Token = m["info"].(map[string]interface{})["token"].(string)
	resp.Name = m["info"].(map[string]interface{})["name"].(string)
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
