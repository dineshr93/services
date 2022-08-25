package main

import (
	"context"
	"log"
	"math"

	pb "github.com/dineshr93/services/calculator/proto"
)

func (s *Server) Square(ctx context.Context, req *pb.SquareRequest) (*pb.SquareResponse, error) {
	log.Println("Computing square function in server for", req)
	return &pb.SquareResponse{
		Number: math.Pow(req.Number, 2),
	}, nil

}
