package storage

import (
	"test1/pkg/api"
	"testing"
)

func (s *Storage) TestFindByAuthor(t *testing.T) {
	v, _ := s.FindByAuthor(&api.Request{Search: "Пушкин"})
	if v.Books[0].Name != "Евгений Онегин" {
		t.Error("Expected Евгений Онегин, got ", v.Books[0].Name)
	}
}

func (s *Storage) TestFindByTitle(t *testing.T) {
	v, _ := s.FindByTitle(&api.Request{Search: "Пометель"})
	if v.Authors[0].Name != "Пупкин" {
		t.Error("Expected Пупкин, got ", v.Authors[0].Name)
	}
}
