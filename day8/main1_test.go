package main

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Addbook2(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO Book").WithArgs("three", "two").
		WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewStore(db)

	book := Book{
		Author: "three",
		Title:  "two",
	}
	body, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/book", io.NopCloser(bytes.NewReader(body)))
	w := httptest.NewRecorder()

	store.Addbook(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status %d, but got %d", http.StatusCreated, resp.StatusCode)
	}

	var createdBook Book
	if err := json.NewDecoder(resp.Body).Decode(&createdBook); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if createdBook.Author != "three" || createdBook.Title != "two" {
		t.Errorf("Expected book with author 'three' and title 'two', but got %+v", createdBook)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func Test_Getbook2(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "author", "title"}).
		AddRow(4, "three", "two")
	mock.ExpectQuery("SELECT id, author, title FROM Book WHERE id = ?").
		WithArgs(4).
		WillReturnRows(rows)

	store := NewStore(db)

	req := httptest.NewRequest(http.MethodGet, "/book/4", nil)
	w := httptest.NewRecorder()

	store.Getbook(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	var fetchedBook Book
	if err := json.NewDecoder(resp.Body).Decode(&fetchedBook); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if fetchedBook.Author != "three" || fetchedBook.Title != "two" {
		t.Errorf("Expected book with author 'three' and title 'two', but got %+v", fetchedBook)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func Test_Updatebook(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE Book SET author = ?, title = ? WHERE id = ?").
		WithArgs("updated author", "updated title", 7).
		WillReturnResult(sqlmock.NewResult(0, 1))

	store := NewStore(db)

	updatedBook := Book{
		Author: "updated author",
		Title:  "updated title",
	}
	body, err := json.Marshal(updatedBook)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/book/7", io.NopCloser(bytes.NewReader(body)))
	w := httptest.NewRecorder()

	store.Updatebook(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	var resultBook Book
	if err := json.NewDecoder(resp.Body).Decode(&resultBook); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if resultBook.Author != "updated author" || resultBook.Title != "updated title" {
		t.Errorf("Expected updated book, but got %+v", resultBook)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func Test_Deletebook2(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM Book WHERE id = ?").
		WithArgs(7).
		WillReturnResult(sqlmock.NewResult(0, 1))

	store := NewStore(db)

	req := httptest.NewRequest(http.MethodDelete, "/book/7", nil)
	w := httptest.NewRecorder()

	store.Deletebook(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status %d, but got %d", http.StatusNoContent, resp.StatusCode)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
