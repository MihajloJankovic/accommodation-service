package handlers

import (
	"context"
	"fmt"
	protos "github.com/MihajloJankovic/accommodation-service/protos/main"
	"github.com/gocql/gocql"
	"log"
	"os"
)

// New NoSQL: Constructor which reads db configuration from environment
type AccommodationRepo struct {
	logger  *log.Logger
	session *gocql.Session
}

func (ar *AccommodationRepo) CloseSession() {
	ar.session.Close()
}

// New: Constructor which reads db configuration from environment
func New(ctx context.Context, logger *log.Logger) (*AccommodationRepo, error) {
	db := os.Getenv("CASS_DB")

	// Connect to default keyspace
	cluster := gocql.NewCluster(db)
	cluster.Keyspace = "system"
	session, err := cluster.CreateSession()
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	defer session.Close()

	// Create 'student' keyspace
	err = session.Query(
		fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s
					WITH replication = {
						'class' : 'SimpleStrategy',
						'replication_factor' : %d
					}`, "accommodation", 1)).Exec()
	if err != nil {
		logger.Println(err)
	}

	// Connect to student keyspace
	cluster.Keyspace = "accommodation"
	cluster.Consistency = gocql.One
	session, err = cluster.CreateSession()
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	// Return repository with logger and DB session
	return &AccommodationRepo{
		session: session,
		logger:  logger,
	}, nil
}

// Disconnect from database
func (ar *AccommodationRepo) Disconnect(ctx context.Context) {
	ar.session.Close()
}

func (sr *AccommodationRepo) CreateTables() {
	err := sr.session.Query(
		`CREATE TABLE IF NOT EXISTS accomondations (
			uid UUID PRIMARY KEY,
			name text,
			location text,
			adress text,
			email text,
			amenities list<text>
		)`).Exec()
	if err != nil {
		sr.logger.Println(err)
	}
}

// GetAll fetches all accommodations from the database.
func (ar *AccommodationRepo) GetAll() ([]*protos.AccommodationResponse, error) {
	session := ar.session

	var accommodationsSlice []*protos.AccommodationResponse

	query := "SELECT uid, name, location, adress, email, amenities FROM accomondations ALLOW FILTERING"
	iter := session.Query(query).Iter()

	var accommodation protos.AccommodationResponse
	for iter.Scan(
		&accommodation.Uid,
		&accommodation.Name,
		&accommodation.Location,
		&accommodation.Adress,
		&accommodation.Email,
		&accommodation.Amenities,
	) {
		// Create a new instance for each row
		currentAccommodation := &protos.AccommodationResponse{
			Uid:       accommodation.Uid,
			Name:      accommodation.Name,
			Location:  accommodation.Location,
			Adress:    accommodation.Adress,
			Email:     accommodation.Email,
			Amenities: accommodation.Amenities,
		}

		// Append the new instance to the slice
		accommodationsSlice = append(accommodationsSlice, currentAccommodation)
	}

	if err := iter.Close(); err != nil {
		ar.logger.Println(err)
		return nil, err
	}

	return accommodationsSlice, nil
}

func (ar *AccommodationRepo) GetByUuid(id string) (*protos.AccommodationResponse, error) {
	query := "SELECT uid, name, location, adress, email, amenities FROM accomondations WHERE uid = ? ALLOW FILTERING"
	var acc protos.AccommodationResponse

	if err := ar.session.Query(query, id).Scan(
		&acc.Uid,
		&acc.Name,
		&acc.Location,
		&acc.Adress,
		&acc.Email,
		&acc.Amenities,
	); err != nil {
		ar.logger.Println(err)
		return nil, err
	}

	return &acc, nil
}

func (ar *AccommodationRepo) GetById(email string) ([]*protos.AccommodationResponse, error) {
	query := "SELECT uid, name, location, adress, email, amenities FROM accomondations WHERE email = ? ALLOW FILTERING"
	iter := ar.session.Query(query, email).Iter()

	var accommodationsSlice []*protos.AccommodationResponse
	var accommodation protos.AccommodationResponse
	for iter.Scan(
		&accommodation.Uid,
		&accommodation.Name,
		&accommodation.Location,
		&accommodation.Adress,
		&accommodation.Email,
		&accommodation.Amenities,
	) {
		// Create a new instance for each row
		currentAccommodation := &protos.AccommodationResponse{
			Uid:       accommodation.Uid,
			Name:      accommodation.Name,
			Location:  accommodation.Location,
			Adress:    accommodation.Adress,
			Email:     accommodation.Email,
			Amenities: accommodation.Amenities,
		}

		// Append the new instance to the slice
		accommodationsSlice = append(accommodationsSlice, currentAccommodation)
	}

	if err := iter.Close(); err != nil {
		ar.logger.Println(err)
		return nil, err
	}
	log.Println(accommodationsSlice)
	return accommodationsSlice, nil
}

func (ar *AccommodationRepo) Create(profile *protos.AccommodationResponse) error {
	query := "INSERT INTO accomondations (uid, name, location, adress, email, amenities) VALUES (?, ?, ?, ?, ?, ?)"

	err := ar.session.Query(query,
		profile.Uid,
		profile.Name,
		profile.Location,
		profile.Adress,
		profile.Email,
		profile.Amenities,
	).Exec()

	if err != nil {
		ar.logger.Println("Error inserting accommodation record:", err)
		return err
	}

	return nil
}
func (ar *AccommodationRepo) Update(accommodation *protos.AccommodationResponse) error {
	query := "UPDATE accomondations SET name = ?, location = ?, adress = ? WHERE email = ?"
	if err := ar.session.Query(query,
		accommodation.Name,
		accommodation.Location,
		accommodation.Adress,
		accommodation.Email,
	).Exec(); err != nil {
		ar.logger.Println(err)
		return err
	}

	return nil
}

func (ar *AccommodationRepo) DeleteByID(id string) error {
	query := "DELETE FROM accomondations WHERE uid = ?"
	if err := ar.session.Query(query, id).Exec(); err != nil {
		ar.logger.Println(err)
		return err
	}

	return nil
}

func (ar *AccommodationRepo) FilterByPriceRange(minPrice, maxPrice float32) ([]*protos.AccommodationResponse, error) {
	query := "SELECT uid, name, location, adress, email, amenities FROM accomondations WHERE price >= ? AND price <= ? ALLOW FILTERING"
	iter := ar.session.Query(query, minPrice, maxPrice).Iter()

	var accommodationsSlice []*protos.AccommodationResponse
	var accommodation protos.AccommodationResponse
	for iter.Scan(
		&accommodation.Uid,
		&accommodation.Name,
		&accommodation.Location,
		&accommodation.Adress,
		&accommodation.Email,
		&accommodation.Amenities,
	) {
		// Create a new instance for each row
		currentAccommodation := &protos.AccommodationResponse{
			Uid:       accommodation.Uid,
			Name:      accommodation.Name,
			Location:  accommodation.Location,
			Adress:    accommodation.Adress,
			Email:     accommodation.Email,
			Amenities: accommodation.Amenities,
		}

		// Append the new instance to the slice
		accommodationsSlice = append(accommodationsSlice, currentAccommodation)
	}

	if err := iter.Close(); err != nil {
		ar.logger.Println(err)
		return nil, err
	}

	return accommodationsSlice, nil
}

func (ar *AccommodationRepo) FilterByAmenities(amenitiesList []string) ([]*protos.AccommodationResponse, error) {
	query := "SELECT uid, name, location, adress, email, amenities FROM accomondations WHERE amenities CONTAINS ? ALLOW FILTERING"
	iter := ar.session.Query(query, amenitiesList).Iter()

	var filteredAccommodations []*protos.AccommodationResponse
	var accommodation protos.AccommodationResponse

	for iter.Scan(
		&accommodation.Uid,
		&accommodation.Name,
		&accommodation.Location,
		&accommodation.Adress,
		&accommodation.Email,
		&accommodation.Amenities,
	) {
		// Create a new instance for each row
		currentAccommodation := &protos.AccommodationResponse{
			Uid:       accommodation.Uid,
			Name:      accommodation.Name,
			Location:  accommodation.Location,
			Adress:    accommodation.Adress,
			Email:     accommodation.Email,
			Amenities: accommodation.Amenities,
		}

		// Append the new instance to the slice
		filteredAccommodations = append(filteredAccommodations, currentAccommodation)
	}
	if err := iter.Close(); err != nil {
		ar.logger.Println(err)
		return nil, err
	}

	return filteredAccommodations, nil
}

func (ar *AccommodationRepo) FilterByHost(hostEmail string) ([]*protos.AccommodationResponse, error) {
	query := "SELECT uid, name, location, adress, email, amenities FROM accomondations WHERE email = ? ALLOW FILTERING"
	iter := ar.session.Query(query, hostEmail).Iter()

	var filteredAccommodations []*protos.AccommodationResponse

	var accommodation protos.AccommodationResponse

	for iter.Scan(
		&accommodation.Uid,
		&accommodation.Name,
		&accommodation.Location,
		&accommodation.Adress,
		&accommodation.Email,
		&accommodation.Amenities,
	) {
		// Create a new instance for each row
		currentAccommodation := &protos.AccommodationResponse{
			Uid:       accommodation.Uid,
			Name:      accommodation.Name,
			Location:  accommodation.Location,
			Adress:    accommodation.Adress,
			Email:     accommodation.Email,
			Amenities: accommodation.Amenities,
		}

		// Append the new instance to the slice
		filteredAccommodations = append(filteredAccommodations, currentAccommodation)
	}
	if err := iter.Close(); err != nil {
		ar.logger.Println(err)
		return nil, err
	}

	return filteredAccommodations, nil
}
