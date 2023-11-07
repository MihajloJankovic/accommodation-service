package main

import (
	"context"
	"github.com/MihajloJankovic/accommodation-service/handlers"
	protos "github.com/MihajloJankovic/accommodation-service/protos/glavno"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	lis, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Fatal(err)
	}
	serverRegister := grpc.NewServer()

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[accommodation-glavno] ", log.LstdFlags)
	accommodationLog := log.New(os.Stdout, "[accommodation-repo-log] ", log.LstdFlags)

	accommodationRepo, err := handlers.New(timeoutContext, accommodationLog)
	if err != nil {
		logger.Fatal(err)
	}
	defer accommodationRepo.Disconnect(timeoutContext)

	// NoSQL: Checking if the connection was established
	accommodationRepo.Ping()

	//Initialize the handler and inject said logger
	service := handlers.NewServer(logger, accommodationRepo)

	protos.RegisterAccommodationServer(serverRegister, service)
	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
