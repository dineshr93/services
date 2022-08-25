package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/dineshr93/services/calculator/proto"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	if len(os.Args) > 1 {
		addr = os.Args[1]
		// log.Println(number)

		// do something with command
	} else {
		fmt.Println("Please provide one number as an argument")
		os.Exit(1)
	}
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
