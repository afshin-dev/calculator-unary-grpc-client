package main

import (
	"calculator-unary-grpc-client/calculator"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}
	c := calculator.NewCalculateClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// this loop run until context deadline exceed
	for i := 1; i < 1000000; i++ {
		resp, err := c.Execute(ctx, &calculator.CalculateRequest{Lhs: int64(i), Rhs: int64(0), Operation: calculator.Operation_ADD})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("result is: ", resp.Result)
	}

}
