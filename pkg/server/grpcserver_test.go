package server

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
	"test1/pkg/api"
	"test1/pkg/storage"
	"testing"
)

func TestGRPCServer_FindAuthors(t *testing.T) {
	s, deferFunc := initServer()
	defer deferFunc()

	t.Run("Find successful", func(t *testing.T) {
		resp, err := s.FindAuthors(context.Background(), &api.Request{Search: "Метель"})
		assert.Nil(t, err)
		assert.Len(t, resp.Authors, 1)
		assert.Equal(t, "Пушкин", resp.Authors[0].Name)
	})

	t.Run("No results", func(t *testing.T) {
		resp, err := s.FindAuthors(context.Background(), &api.Request{Search: "белиберда"})
		assert.Nil(t, err)
		assert.Len(t, resp.Authors, 0)
	})
}

func TestGRPCServer_FindBooks(t *testing.T) {
	s, deferFunc := initServer()
	defer deferFunc()
	t.Run("Find successful", func(t *testing.T) {
		resp, err := s.FindBooks(context.Background(), &api.Request{Search: "Пушкин"})
		assert.Nil(t, err)
		assert.Len(t, resp.Books, 2)
		assert.Equal(t, "Евгений Онегин", resp.Books[0].Name)
		assert.Equal(t, "Метель", resp.Books[1].Name)
	})
	t.Run("No results", func(t *testing.T) {
		resp, err := s.FindBooks(context.Background(), &api.Request{Search: "Белиберда"})
		assert.Nil(t, err)
		assert.Len(t, resp.Books, 0)
	})
}

func initServer() (*GRPCServer, func()) {
	db, err := sql.Open("mysql", "user:pass@(localhost:3307)/db?charset=utf8")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		color.Error.Println("====== Are you started the database?")
		panic(err)
	}

	return &GRPCServer{St: storage.NewStorage(db)}, func() {
		db.Close()
	}
}
