package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/dineshr93/services/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	var number string

	if len(os.Args) > 1 {
		addr = os.Args[1]
		number = os.Args[2]
		// log.Println(number)

		// do something with command
	} else {
		fmt.Println("Please provide address and one number as an argument")
		os.Exit(1)
	}
	tls := true
	opts := []grpc.DialOption{}
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Error while loading CA trust certificate:%v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	
	// 1. get the grpc connection by grpc dial the address
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}
	
	//2. defer the close 
	defer conn.Close()
	
	//3. get the particular client from the grpc connection
	c := pb.NewCalculatorServiceClient(conn)

	//4. Client contains the rpc method defined in the proto file
	//c.Square
	
	// doSum(c)
	// doPrimes(c)
	// doAvg(c)
	// doMax(c)
	// doSqrt(c, 10)
	s, _ := strconv.ParseFloat(number, 64)
	doSqrt(c, s)
}
