package storage

import (
	"database/sql"
	"fmt"
	"test1/pkg/api"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(conn *sql.DB) *Storage {
	return &Storage{db: conn}
}

func (s *Storage) FindByAuthor(req *api.Request) (api.Books, error) {
	answer := []*api.Book{}
	rows, err := s.db.Query("SELECT title FROM author INNER JOIN book INNER JOIN books_authors ON author.author_id = book.author_id WHERE name_author = ?;", req.Search)
	if err != nil {
		fmt.Println(err)
		return api.Books{}, err
	}
	defer rows.Close()
	for rows.Next() {
		p := &api.Book{}
		err := rows.Scan(&p.Name)
		if err != nil {
			fmt.Println(err)
			return api.Books{}, err
		}
		answer = append(answer, p)
	}
	return api.Books{Books: answer}, nil
}

func (s *Storage) FindByTitle(req *api.Request) (api.Authors, error) {
	answer := []*api.Author{}
	rows, err := s.db.Query("SELECT author FROM author INNER JOIN book INNER JOIN books_authors ON author.author_id = book.author_id WHERE title = ?;", req.Search)
	if err != nil {
		fmt.Println(err)
		return api.Authors{}, err
	}
	defer rows.Close()
	for rows.Next() {
		p := &api.Author{}
		err := rows.Scan(&p.Name)
		if err != nil {
			fmt.Println(err)
			return api.Authors{}, err
		}
		answer = append(answer, p)
	}
	return api.Authors{Authors: answer}, nil
}
