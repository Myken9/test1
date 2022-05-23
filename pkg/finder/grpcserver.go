package finder

import (
	"context"
	"test1/pkg/api"
	"test1/pkg/storage"
)

type GRPCServer struct {
	St *storage.Storage
}

func (s *GRPCServer) FindBooks(_ context.Context, req *api.Request) (*api.Books, error) {
	answer, err := s.St.FindByAuthor(req)
	return &answer, err
}

func (s *GRPCServer) FindAuthors(_ context.Context, req *api.Request) (*api.Authors, error) {
	answer, err := s.St.FindByTitle(req)
	return &answer, err
}
