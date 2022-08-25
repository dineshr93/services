package main

import (
	"context"
	"log"

	pb "github.com/dineshr93/services/calculator/proto"
)

func doSqrt(c pb.CalculatorServiceClient, number float64) {
	// log.Printf("CalculatorServiceClient")
	res, err := c.Square(context.Background(), &pb.SquareRequest{
		Number: number,
	})

	if err != nil {
		log.Fatalf("Could not square the number: %v \n", err)
	}
	log.Println("Square of ", number, " is: ", res.Number)
}
