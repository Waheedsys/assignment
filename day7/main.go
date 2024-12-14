package main

import (
	"encoding/json"
	"fmt"
	"github.com/waheedsys/newproject/day7/assignment8"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Post function
func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var newBook assignment8.Book
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading body: %v", err), http.StatusBadRequest)
		return
	}

	// Unmarshal the body into the Book struct
	err = json.Unmarshal(bodyBytes, &newBook)
	if err != nil {
		http.Error(w, "Failed to Unmarshal book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Save book to the database
	createdBook, err := assignment8.Post(newBook)
	if err != nil {
		http.Error(w, "Failed to post the book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back the created book as JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}

// Get function
func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Extract the book ID from the URL path (after "/book/")
		idStr := r.URL.Path[len("/book/"):]
		if idStr == "" {
			http.Error(w, "Book ID is required", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		// Retrieve book from database
		book, err := assignment8.Get(id)
		if err != nil {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		// Send back the book as JSON
		json.NewEncoder(w).Encode(book)
		return
	}

	// If method is not GET
	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)

}

// put function
func Put(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/book/update/"):]
	if idStr == "" {
		http.Error(w, "Book ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Read the updated book data
	var updatedBook assignment8.Book
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading body: %v", err), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bodyBytes, &updatedBook)
	if err != nil {
		http.Error(w, "Failed to unmarshal updated book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Save the updated book
	err = assignment8.Put(id, updatedBook)
	if err != nil {
		http.Error(w, "Failed to update the book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the updated book
	json.NewEncoder(w).Encode(updatedBook)
}

// Delete function
func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	//get id
	idStr := r.URL.Path[len("/book/delete/"):]
	if idStr == "" {
		http.Error(w, "Book ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	err = assignment8.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete the book: "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("done"))
}

func main() {
	// POST: Create a new book
	http.HandleFunc("/book", Post)

	// GET: Retrieve a book by ID (Using /book/{id} route)
	http.HandleFunc("/book/", Get)

	// PUT: Update an existing book by ID (Using /book/update/{id} route)
	http.HandleFunc("/book/update/", Put)

	//delete
	http.HandleFunc("/book/", Delete)

	// Start the server
	port := ":5000"
	fmt.Println("Server is running on port" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
