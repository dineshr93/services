package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/dineshr93/services/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	if len(os.Args) > 1 {
		addr = os.Args[1]
		// log.Println(number)

		// do something with command
	} else {
		fmt.Println("Please provide address as an argument")
		// os.Exit(1)
	}

	opts := []grpc.ServerOption{}
	tsl := true
	if tsl {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))

	}
	// 1. create new grpc server
	s := grpc.NewServer(opts...)
	// 2. Register your service with the created server
	pb.RegisterCalculatorServiceServer(s, &Server{})
	// 3. tcp listener
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	log.Printf("Listening at %s\n", addr)
	// 4. connect grpc server with tcp listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
