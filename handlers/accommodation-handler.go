package handlers

import (
	"context"
	protos "github.com/MihajloJankovic/accommodation-service/protos/accommodation_service"
	"google.golang.org/grpc"
	"log"
)

type myAccommodationServer struct {
	protos.UnimplementedAccommodationServer
	logger *log.Logger
	// NoSQL: injecting product repository
	repo *AccommodationRepo
}

func NewServer(l *log.Logger, r *AccommodationRepo) *myAccommodationServer {
	return &myAccommodationServer{*new(protos.UnimplementedAccommodationServer), l, r}
}

func (s myAccommodationServer) GetAccommodation(ctx context.Context, in *protos.AccommodationRequest, opts ...grpc.CallOption) (*protos.AccommodationResponse, error) {
	return nil, nil
}
func (s myAccommodationServer) SetAccommodation(ctx context.Context, in *protos.AccommodationResponse, opts ...grpc.CallOption) (*protos.Empty, error) {
	return nil, nil
}
func (s myAccommodationServer) UpdateAccommodation(ctx context.Context, in *protos.AccommodationResponse, opts ...grpc.CallOption) (*protos.Empty, error) {
	return nil, nil
}
