package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Book struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// get
func (s *Store) Getbook(w http.ResponseWriter, r *http.Request) {
	idstr := strings.TrimPrefix(r.URL.Path, "/book/")
	id, err := strconv.Atoi(idstr)
	if err != nil || id <= 0 {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return

	}
	var book Book
	err = s.db.QueryRow("SELECT id, author, title FROM Book WHERE id = ?", id).Scan(&book.Id, &book.Author, &book.Title)
	if err == sql.ErrNoRows {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content_type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// post
func (s *Store) Addbook(w http.ResponseWriter, r *http.Request) {
	var book Book
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cannot read request body", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(b, &book)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	result, err := s.db.Exec("INSERT INTO Book (author, title) values (?, ?)", book.Author, book.Title)
	if err != nil {
		http.Error(w, "falied to add to book", http.StatusInternalServerError)
		return
	}
	id, _ := result.LastInsertId()
	book.Id = int(id)
	w.Header().Set("content_type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// update
func (s *Store) updatebook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/book/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}
	var updatedBook Book
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(b, &updatedBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	_, err = s.db.Exec("UPDATE Book SET author = ?, title = ? WHERE id = ?", updatedBook.Author, updatedBook.Title, id)
	if err != nil {
		http.Error(w, "failed to update the book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("book with id %d updated", id)))
}

// delete
func (s *Store) Deletebook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/book/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	_, err = s.db.Exec("DELETE FROM Book WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("Book with id %d deleted", id)))
}

func main() {

	db, err := sql.Open("mysql", "root:password@/sample")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("error connection to database %v", err)
	}

	fmt.Println("connection successful")

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	store := NewStore(db)

	http.HandleFunc("/book/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			store.Getbook(w, r)
		case http.MethodPut:
			store.updatebook(w, r)
		case http.MethodDelete:
			store.Deletebook(w, r)
		case http.MethodPost:
			store.Addbook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := ":5001"
	fmt.Println("Server is running on port" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
