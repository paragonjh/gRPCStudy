package main

import (
	"gitlab.com/techschool/pcbook/pb"
	"gitlab.com/techschool/pcbook/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Client Stream gRPC: Server App
func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	imageStore := service.NewDiskImageStore("img")
	laptopServer := service.NewLaptopServer(imageStore)
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", "localhost:4343")
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
