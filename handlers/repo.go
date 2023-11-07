package handlers

import (
	"context"
	"fmt"
	protos "github.com/MihajloJankovic/accommodation-service/protos/main"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type AccommodationRepo struct {
	logger *log.Logger
	cli    *mongo.Client
}

// NoSQL: Constructor which reads db configuration from environment
func New(ctx context.Context, logger *log.Logger) (*AccommodationRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &AccommodationRepo{
		cli:    client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (ar *AccommodationRepo) Disconnect(ctx context.Context) error {
	err := ar.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (ar *AccommodationRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := ar.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		ar.logger.Println(err)
	}
	// Print available databases
	databases, err := ar.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		ar.logger.Println(err)
	}
	fmt.Println(databases)
}
func (ar *AccommodationRepo) GetAll() (*[]protos.AccommodationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accommodationCollection := ar.getCollection()
	var accommodationsSlice []protos.AccommodationResponse

	accommodationCursor, err := accommodationCollection.Find(ctx, bson.M{})
	if err != nil {
		ar.logger.Println(err)
		return nil, err
	}
	if err = accommodationCursor.All(ctx, &accommodationsSlice); err != nil {
		ar.logger.Println(err)
		return nil, err
	}
	return &accommodationsSlice, nil
}
func (ar *AccommodationRepo) GetById(email string) (*protos.AccommodationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accommodationCollection := ar.getCollection()
	var accommodation protos.AccommodationResponse

	err := accommodationCollection.FindOne(ctx, bson.M{"email": email}).Decode(&accommodation)
	if err != nil {
		ar.logger.Println(err)
		return nil, err
	}

	return &accommodation, nil
}
func (ar *AccommodationRepo) Create(profile *protos.AccommodationResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	accommodationCollection := ar.getCollection()

	result, err := accommodationCollection.InsertOne(ctx, &profile)
	if err != nil {
		ar.logger.Println(err)
		return err
	}
	ar.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (ar *AccommodationRepo) Update(accommodation *protos.AccommodationResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	accommodationCollection := ar.getCollection()

	filter := bson.M{"email": accommodation.GetEmail()}
	update := bson.M{"$set": bson.M{
		"name":     accommodation.GetName(),
		"price":    accommodation.GetPrice(),
		"location": accommodation.GetLocation(),
		"adress":   accommodation.GetAdress(),
	}}
	result, err := accommodationCollection.UpdateOne(ctx, filter, update)
	ar.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	ar.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		ar.logger.Println(err)
		return err
	}
	return nil
}

func (ar *AccommodationRepo) getCollection() *mongo.Collection {
	accommodationDatabase := ar.cli.Database("mongoAccommodation")
	accommodationCollection := accommodationDatabase.Collection("accommodations")
	return accommodationCollection
}
