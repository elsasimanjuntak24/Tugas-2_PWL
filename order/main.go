package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "grpc/proto/order" // Ganti dengan path yang sesuai
)

type orderServer struct {
	pb.UnimplementedOrderServiceServer
}

func (s *orderServer) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.CreateOrderResponse{OrderId: "12345", Status: "COMPLETED"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &orderServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}