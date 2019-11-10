package main

import (
	"fmt"
	"log"
	"net"

	"github.com/OdaDaisuke/grpc-web/server/application"
	pb "github.com/OdaDaisuke/grpc-web/server/proto"
	grpc "google.golang.org/grpc"
)

func main() {
	fmt.Println(&pb.PostMessageResponse{})
	grpcSvr := grpc.NewServer()
	fmt.Println(grpcSvr)
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	app := &application.MainApplication{}
	pb.RegisterMessageServiceServer(grpcSvr, app)
	grpcSvr.Serve(lis)
}
