package main

import (
	pb "chaitsgithub/goprojects/go-gRPC/calculator"
	context "context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// server is used to implement calculator.CalculatorServer.
type server struct {
	pb.UnimplementedCalculatorServer
}

// Add implements calculator.CalculatorServer
func (s *server) Add(ctx context.Context, in *pb.Input) (*pb.Output, error) {
	log.Printf("Received: %v %v", in.GetOperand1(), in.GetOperand2())
	result := in.GetOperand1() + in.GetOperand2()
	return &pb.Output{Result: result}, nil
}

func main() {
	// flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
