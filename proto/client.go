package proto

import (
	"google.golang.org/grpc"
	"sct.impl.host/proto/auth"
	"sct.impl.host/proto/user"
)

type Client struct {
	Conn *grpc.ClientConn
	Auth auth.AuthServiceClient
	User user.UserCrudServiceClient
}
