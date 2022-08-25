package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/dineshr93/services/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	var number string
	if len(os.Args) > 1 {
		number = os.Args[1]
		// log.Println(number)

		// do something with command
	} else {
		fmt.Println("Please provide one number as an argument")
		os.Exit(1)
	}
	// doSum(c)
	// doPrimes(c)
	// doAvg(c)
	// doMax(c)
	// doSqrt(c, 10)
	s, _ := strconv.ParseFloat(number, 64)
	doSqrt(c, s)
}
