package storage

import (
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage_FindByTitle(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer database.Close()

	storage := NewStorage(database)

	expectedQuery := "SELECT author.id, author.name FROM author"
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Дуглас Адамс")

	mock.ExpectQuery(expectedQuery).WithArgs().WillReturnRows(rows)
	authors, err := storage.FindByTitle("Автостопом по галактике")
	assert.Nil(t, err)
	assert.IsType(t, []Author{}, authors)
}

func TestStorage_FindByAuthor(t *testing.T) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer database.Close()

	storage := NewStorage(database)

	expectedQuery := "SELECT book.id, book.title FROM book"
	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow(42, "Автостопом по галактике")

	mock.ExpectQuery(expectedQuery).WithArgs().WillReturnRows(rows)
	books, err := storage.FindByAuthor("Дуглас Адамс")
	assert.Nil(t, err)
	assert.IsType(t, []Book{}, books)
}
