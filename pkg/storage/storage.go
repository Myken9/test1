package storage

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(conn *sql.DB) *Storage {
	return &Storage{db: conn}
}

func (s *Storage) FindByAuthor(authorName string) ([]Book, error) {
	var books []Book
	q := `
		SELECT book.id, book.title FROM book 
        JOIN books_authors ON books_authors.book_id = book.id 
        JOIN author ON author.id = books_authors.author_id
        WHERE author.name = ?
	`
	rows, err := s.db.Query(q, authorName)
	if err != nil {
		fmt.Println(err)
		return books, err
	}
	defer rows.Close()
	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.ID, &book.Title)
		if err != nil {
			fmt.Println(err)
			return books, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *Storage) FindByTitle(bookTitle string) ([]Author, error) {
	var authors []Author
	q := `
		SELECT author.id, author.name FROM author
		JOIN books_authors ON books_authors.author_id = author.id
		JOIN book ON book.id = books_authors.book_id
		WHERE book.title = ?
	`
	rows, err := s.db.Query(q, bookTitle)
	if err != nil {
		fmt.Println(err)
		return authors, err
	}
	defer rows.Close()
	for rows.Next() {
		author := Author{}
		err := rows.Scan(&author.ID, &author.Name)
		if err != nil {
			fmt.Println(err)
			return authors, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}
