package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	shippingpb "grpc/proto/shipping"
)

type shippingServer struct {
	shippingpb.UnimplementedShippingServiceServer
}

func (s *shippingServer) StartShipping(ctx context.Context, req *shippingpb.ShippingRequest) (*shippingpb.ShippingResponse, error) {
	log.Printf("Starting shipping for Order ID: %v", req.OrderId)
	// Dummy response for simplicity
	return &shippingpb.ShippingResponse{ShipmentId: "S1234", Status: "SHIPPED"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	shippingpb.RegisterShippingServiceServer(s, &shippingServer{})
	log.Printf("Shipping server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}