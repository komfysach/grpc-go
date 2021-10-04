package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"grpc.om/m/v2/calculator/calculatorpb"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalcuServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)
}

func doUnary(c calculatorpb.CalcuServiceClient) {
	req := &calculatorpb.CalcuRequest{
		Calculation: &calculatorpb.Calcu{
			FirstNumber:  3,
			SecondNumber: 10,
		},
	}
	res, err := c.Calcu(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
