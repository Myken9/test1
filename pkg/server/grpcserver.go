package server

import (
	"context"
	"test1/pkg/api"
	"test1/pkg/storage"
)

type GRPCServer struct {
	St *storage.Storage
}

func (s *GRPCServer) FindBooks(_ context.Context, req *api.Request) (*api.Books, error) {
	books, err := s.St.FindByAuthor(req.Search)
	if err != nil {
		return nil, err
	}
	return booksToApi(books), nil
}

func (s *GRPCServer) FindAuthors(_ context.Context, req *api.Request) (*api.Authors, error) {
	authors, err := s.St.FindByTitle(req.Search)
	if err != nil {
		return nil, err
	}
	return authorsToApi(authors), nil
}
