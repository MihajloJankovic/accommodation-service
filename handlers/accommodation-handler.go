package handlers

import (
	"context"
	protos "github.com/MihajloJankovic/accommodation-service/protos/glavno"
	"log"
)

type MyAccommodationServer struct {
	protos.UnimplementedAccommodationServer
	logger *log.Logger
	// NoSQL: injecting product repository
	repo *AccommodationRepo
}

func NewServer(l *log.Logger, r *AccommodationRepo) *MyAccommodationServer {
	return &MyAccommodationServer{*new(protos.UnimplementedAccommodationServer), l, r}
}

//GetAccommodation(context.Context, *AccommodationRequest) (*AccommodationResponse, error)
//SetAccommodation(context.Context, *AccommodationResponse) (*Empty, error)
//UpdateAccommodation(context.Context, *AccommodationResponse) (*Empty, error)

func (s MyAccommodationServer) GetAccommodation(_ context.Context, in *protos.AccommodationRequest) (*protos.DummyList, error) {
	out, err := s.repo.GetById(in.GetEmail())
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	ss := new(protos.DummyList)
	ss.Dummy = out
	return ss, nil
}
func (s MyAccommodationServer) SetAccommodation(_ context.Context, in *protos.AccommodationResponse) (*protos.Emptya, error) {
	out := new(protos.AccommodationResponse)
	out.Name = in.GetName()
	out.Price = in.GetPrice()
	out.Location = in.GetLocation()
	out.Adress = in.GetAdress()
	out.Email = in.GetEmail()

	err := s.repo.Create(out)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return new(protos.Emptya), nil
}
func (s MyAccommodationServer) UpdateAccommodation(_ context.Context, in *protos.AccommodationResponse) (*protos.Emptya, error) {
	err := s.repo.Update(in)
	if err != nil {
		return nil, err
	}
	return new(protos.Emptya), nil
}

func (s MyAccommodationServer) FilterByPriceRange(_ context.Context, in *protos.PriceRangeRequest) (*protos.DummyList, error) {
	// Implementiraj filtriranje smeštaja po opsegu cene
	// ...
}

func (s MyAccommodationServer) FilterByAmenities(_ context.Context, in *protos.AmenitiesRequest) (*protos.DummyList, error) {
	// Implementiraj filtriranje smeštaja po pogodnostima
	// ...
}

func (s MyAccommodationServer) FilterByHost(_ context.Context, in *protos.HostRequest) (*protos.DummyList, error) {
	// Implementiraj filtriranje smeštaja po host-u
	// ...
}
