package handlers

import (
	"context"
	protos "github.com/MihajloJankovic/accommodation-service/protos/protoGen"
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
func (s MyAccommodationServer) GetOneAccommodation(ctx context.Context, in *protos.AccommodationRequestOne) (*protos.AccommodationResponse, error) {
	out, err := s.repo.GetByUuid(in.GetId())
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return out, nil
}
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
func (s MyAccommodationServer) GetAllAccommodation(_ context.Context, in *protos.Emptya) (*protos.DummyList, error) {
	out, err := s.repo.GetAll()
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
	out.Uid = in.GetUid()
	out.Name = in.GetName()
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
	filteredAccommodations, err := s.repo.FilterByPriceRange(in.MinPrice, in.MaxPrice)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	dummyList := &protos.DummyList{Dummy: filteredAccommodations}

	return dummyList, nil
}

func (s MyAccommodationServer) FilterByAmenities(_ context.Context, in *protos.AmenitiesRequest) (*protos.DummyList, error) {
	amenitiesList := in.GetAmenities()

	filteredAccommodations, err := s.repo.FilterByAmenities(amenitiesList)
	if err != nil {
		return nil, err
	}

	dummyList := &protos.DummyList{
		Dummy: filteredAccommodations,
	}

	return dummyList, nil
}

func (s MyAccommodationServer) FilterByHost(_ context.Context, in *protos.HostRequest) (*protos.DummyList, error) {
	hostEmail := in.GetHostEmail()

	filteredAccommodations, err := s.repo.FilterByHost(hostEmail)
	if err != nil {
		return nil, err
	}
	dummyList := &protos.DummyList{
		Dummy: filteredAccommodations,
	}

	return dummyList, nil
}
