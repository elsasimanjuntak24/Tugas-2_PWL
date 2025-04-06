package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	paymentpb "grpc/proto/payment"
)

type paymentServer struct {
	paymentpb.UnimplementedPaymentServiceServer
}

func (s *paymentServer) ProcessPayment(ctx context.Context, req *paymentpb.PaymentRequest) (*paymentpb.PaymentResponse, error) {
	log.Printf("Processing payment for Order ID: %v", req.OrderId)
	// Dummy response for simplicity
	return &paymentpb.PaymentResponse{PaymentId: "P1234", Status: "SUCCESS"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	paymentpb.RegisterPaymentServiceServer(s, &paymentServer{})
	log.Printf("Payment server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}