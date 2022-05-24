package server

import (
	"test1/pkg/api"
	"test1/pkg/storage"
)

func booksToApi(books []storage.Book) *api.Books {
	apiBooks := make([]*api.Book, len(books))
	for i := range books {
		apiBooks = append(apiBooks, &api.Book{
			Id:   books[i].ID,
			Name: books[i].Title,
		})
	}
	return &api.Books{
		Books: apiBooks,
	}
}

func authorsToApi(authors []storage.Author) *api.Authors {
	apiAuthors := make([]*api.Author, len(authors))
	for i := range authors {
		apiAuthors = append(apiAuthors, &api.Author{
			Id:   authors[i].ID,
			Name: authors[i].Name,
		})
	}
	return &api.Authors{
		Authors: apiAuthors,
	}
}
