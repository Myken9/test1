package server

import (
	"test1/pkg/api"
	"test1/pkg/storage"
)

func booksToApi(books []storage.Book) *api.Books {
	resp := &api.Books{Books: nil}
	for i := range books {
		resp.Books = append(resp.Books, &api.Book{
			Id:   books[i].ID,
			Name: books[i].Title,
		})
	}
	return resp
}

func authorsToApi(authors []storage.Author) *api.Authors {
	resp := &api.Authors{Authors: nil}
	for i := range authors {
		resp.Authors = append(resp.Authors, &api.Author{
			Id:   authors[i].ID,
			Name: authors[i].Name,
		})
	}
	return resp
}
