package main

import (
	"crypto/tls"
	"log"

	sso "sct.impl.host/internal/sso/auth"
	"sct.impl.host/proto"
	"sct.impl.host/proto/auth"
	"sct.impl.host/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	tlsCredentials := credentials.NewTLS(&tls.Config{
		MinVersion: tls.VersionTLS13,
	})

	conn, err := grpc.NewClient(
		"server:443",
		grpc.WithTransportCredentials(tlsCredentials),
		grpc.WithPerRPCCredentials(sso.Auth{
			Token: "TEST",
		}),
	)
	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}
	defer conn.Close()

	client := proto.Client{
		Conn: conn,
		Auth: auth.NewAuthServiceClient(conn),
		User: user.NewUserCrudServiceClient(conn),
	}
	_ = client
}
