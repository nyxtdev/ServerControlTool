package main

import (
	"log"
	"net"

	sso "sct.impl.host/internal/sso"
	"sct.impl.host/proto/auth"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	auth.RegisterAuthServiceServer(
		grpcServer,
		sso.NewServer(),
	)

	// UserCrudService:
	// user.RegisterUserCrudServiceServer(
	//     grpcServer,
	//     user.NewServer(),
	// )

	log.Println("Start ServerControlTool :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
