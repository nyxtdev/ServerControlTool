package sso

import (
	"context"

	pb "sct.impl.host/proto/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	api pb.AuthServiceClient
}

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		api: pb.NewAuthServiceClient(conn),
	}
}

func (c *Client) Login(ctx context.Context, username string, password string) (string, error) {
	res, err := c.api.Login(ctx, &pb.LoginRequest{Username: username, Password: password})
	if err != nil {
		return "", err
	}

	return res.GetAccessToken(), nil
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.GetUsername() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "username and password are required")
	}

	// TODO: БД

	token := "generated-token"

	return &pb.LoginResponse{
		AccessToken: token,
	}, nil
}
