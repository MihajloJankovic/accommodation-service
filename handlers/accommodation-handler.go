package handlers

import (
	"context"
	protos "github.com/MihajloJankovic/accommodation-service/protos/main"
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
	out, err := s.repo.GetById(in.GetEmail())
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return out, nil
}
func (s myAccommodationServer) SetAccommodation(ctx context.Context, in *protos.AccommodationResponse, opts ...grpc.CallOption) (*protos.Empty, error) {
	out := new(protos.AccommodationResponse)
	out.Name = in.GetName()
	out.Price = in.GetPrice()
	out.Location = in.GetLocation()
	out.Adress = in.GetAdress()

	err := s.repo.Create(out)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return new(protos.Empty), nil
}
func (s myAccommodationServer) UpdateAccommodation(ctx context.Context, in *protos.AccommodationResponse, opts ...grpc.CallOption) (*protos.Empty, error) {
	err := s.repo.Update(in)
	if err != nil {
		return nil, err
	}
	return new(protos.Empty), nil
}
