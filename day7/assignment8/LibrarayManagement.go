package assignment8

import (
	"errors"
)

// Database to hold books (ID -> Book)
var Database = make(map[int]Book)

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

// Post: Create a new book and add to the database
func Post(b Book) (Book, error) {
	// Store book in the database
	Database[b.ID] = b
	return b, nil
}

// Get: Retrieve a book by ID from the database
func Get(id int) (Book, error) {
	book, exists := Database[id]
	if !exists {
		return Book{}, errors.New("book not found")
	}
	return book, nil
}

// Put: Update an existing book by ID
func Put(id int, updatedBook Book) error {
	if _, exists := Database[id]; exists {
		Database[id] = updatedBook
		return nil
	} else {
		return errors.New("book not found")
	}
}

// delete
func Delete(id int) error {
	_, exists := Database[id]
	if !exists {
		return errors.New("book not found")
	}

	// Delete the book from the database
	delete(Database, id)
	return nil
}
