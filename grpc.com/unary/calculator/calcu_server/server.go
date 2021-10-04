package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc.om/m/v2/calculator/calculatorpb"
)

type server struct {
	calculatorpb.UnimplementedCalcuServiceServer
}

func (*server) Calcu(ctx context.Context, req *calculatorpb.CalcuRequest) (*calculatorpb.CalcuResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)
	firstNumber := req.GetCalculation().GetFirstNumber()
	secondNumber := req.GetCalculation().GetSecondNumber()
	result := firstNumber + secondNumber
	res := &calculatorpb.CalcuResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalcuServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
