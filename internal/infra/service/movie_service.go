package service

import (
	"context"
	"io"

	"github.com/janapc/grpc-movies-go/internal/entity"
	"github.com/janapc/grpc-movies-go/internal/pb"
)

type MovieService struct {
	pb.UnimplementedMovieServiceServer
	MovieDb entity.MovieRepository
}

func NewMovieService(db entity.MovieRepository) *MovieService {
	return &MovieService{
		MovieDb: db,
	}
}

func (m *MovieService) SaveMovie(con context.Context, in *pb.SaveMovieRequest) (*pb.Movie, error) {
	movie := entity.NewMovie(in.Title, in.Description, in.ImageUrl)
	id, err := m.MovieDb.Save(movie)
	if err != nil {
		return nil, err
	}
	movieResponse := &pb.Movie{
		Id:          id,
		Title:       movie.Title,
		Description: movie.Description,
		ImageUrl:    movie.ImageUrl,
	}
	return movieResponse, nil
}

func (m *MovieService) SaveManyMovies(stream pb.MovieService_SaveManyMoviesServer) error {
	for {
		movieRequest, err := stream.Recv()
		if err == io.EOF {
			return err
		}
		if err != nil {
			return err
		}
		movie := entity.NewMovie(movieRequest.Title, movieRequest.Description, movieRequest.ImageUrl)
		id, err := m.MovieDb.Save(movie)
		if err != nil {
			return err
		}
		err = stream.Send(&pb.Movie{
			Id:          id,
			Title:       movie.Title,
			Description: movie.Description,
			ImageUrl:    movie.ImageUrl,
		})
		if err != nil {
			return nil
		}
	}
}

func (m *MovieService) DeleteManyMovies(stream pb.MovieService_DeleteManyMoviesServer) error {
	ids := &pb.IdsMovies{}
	for {
		movieRequest, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(ids)
		}
		if err != nil {
			return err
		}
		_, err = m.MovieDb.FindById(movieRequest.Id)
		if err != nil {
			return err
		}
		err = m.MovieDb.Remove(movieRequest.Id)
		if err != nil {
			return err
		}
		ids.Movies = append(ids.Movies, &pb.OnlyIdMovie{
			Id: movieRequest.Id,
		})
	}
}

func (m *MovieService) FindMovieById(con context.Context, in *pb.OnlyIdMovie) (*pb.Movie, error) {
	movie, err := m.MovieDb.FindById(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Movie{
		Id:          in.Id,
		Title:       movie.Title,
		Description: movie.Description,
		ImageUrl:    movie.ImageUrl,
	}, nil
}

func (m *MovieService) AllMovies(con context.Context, in *pb.BlankMovie) (*pb.Movies, error) {
	movies, err := m.MovieDb.FindAll()
	if err != nil {
		return nil, err
	}
	var moviesList []*pb.Movie
	for _, movie := range movies {
		moviesList = append(moviesList, &pb.Movie{
			Id:          movie.ID.Hex(),
			Title:       movie.Title,
			Description: movie.Description,
			ImageUrl:    movie.ImageUrl,
		})
	}
	return &pb.Movies{Movies: moviesList}, nil
}

func (m *MovieService) UpdateMovie(con context.Context, in *pb.UpdateMovieRequest) (*pb.Movie, error) {
	movie, err := m.MovieDb.FindById(in.Id)
	if err != nil {
		return nil, err
	}
	if in.Title != nil {
		movie.Title = *in.Title
	}
	if in.Description != nil {
		movie.Description = *in.Description
	}
	if in.ImageUrl != nil {
		movie.ImageUrl = *in.ImageUrl
	}
	err = m.MovieDb.Update(movie)
	if err != nil {
		return nil, err
	}
	movieResponse := &pb.Movie{
		Id:          movie.ID.Hex(),
		Title:       movie.Title,
		Description: movie.Description,
		ImageUrl:    movie.ImageUrl,
	}
	return movieResponse, nil
}

func (m *MovieService) DeleteMovie(con context.Context, in *pb.OnlyIdMovie) (*pb.OnlyIdMovie, error) {
	_, err := m.MovieDb.FindById(in.Id)
	if err != nil {
		return nil, err
	}
	err = m.MovieDb.Remove(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.OnlyIdMovie{
		Id: in.Id,
	}, nil
}
