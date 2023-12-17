package main

import (
	"context"
	"log"
	"net"

	"github.com/janapc/grpc-movies-go/internal/infra/database"
	"github.com/janapc/grpc-movies-go/internal/infra/service"
	"github.com/janapc/grpc-movies-go/internal/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientMongoDB := database.Connection()
	defer func() {
		if err := clientMongoDB.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	movieCollection := clientMongoDB.Database("grpc-movies").Collection("movies")
	movieDB := database.NewMovieDatabase(movieCollection)
	movieService := service.NewMovieService(movieDB)

	grpcServer := grpc.NewServer()
	pb.RegisterMovieServiceServer(grpcServer, movieService)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
