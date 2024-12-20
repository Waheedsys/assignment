package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/waheedsys/newproject/day7/assignment8"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Post(t *testing.T) {
	createBook := assignment8.Book{
		ID:     1,
		Name:   "three",
		Author: "two",
	}

	fmt.Println(createBook)

	createBook = assignment8.Book{
		ID:     1,
		Name:   "hello",
		Author: "world",
	}

	createBookBody, err := json.Marshal(createBook)
	if err != nil {
		t.Fatalf("error while getting create book body. err : %v", err)
	}
	//invalidJSONBody := []byte(`{"Name": "test"}`)
	testCases := []struct {
		name      string
		method    string
		expected  []byte
		expStatus int
		id        string
		body      []byte
	}{
		{
			name:      "success with name",
			body:      createBookBody,
			expected:  createBookBody,
			method:    http.MethodPost,
			expStatus: http.StatusCreated,
		},
		{
			name:      "success with different name",
			expected:  createBookBody,
			body:      createBookBody,
			method:    http.MethodPost,
			expStatus: http.StatusCreated,
		},
		{ // with different method
			name:      "success with name",
			body:      createBookBody,
			expected:  []byte("Invalid method"),
			method:    http.MethodGet,
			expStatus: http.StatusMethodNotAllowed,
		},
		//{
		//	name:      "invalid JSON format",
		//	body:      invalidJSONBody,                                       // Malformed JSON body
		//	expected:  []byte("Failed to Unmarshal book: invalid character"), // Expected error message
		//	method:    http.MethodPost,
		//	expStatus: http.StatusInternalServerError,
		//},
		//{
		//	name:      "missing fields in body",
		//	body:      []byte(`{"Name": "Incomplete Book"}`),                             // Missing Author field
		//	expected:  []byte("Failed to Unmarshal book: missing required field Author"), // Expected error message
		//	method:    http.MethodPost,
		//	expStatus: http.StatusBadRequest,
		//},
	}

	for _, tc := range testCases {
		body := bytes.NewBuffer(tc.body)

		req := httptest.NewRequest(tc.method, "/book", body)
		w := httptest.NewRecorder()

		Post(w, req)

		resp := w.Result()
		b, _ := io.ReadAll(resp.Body)

		if !strings.Contains(string(b), string(tc.expected)) {
			t.Errorf("Expected %s, but got %s", tc.expected, string(b))
		}

		if w.Code != tc.expStatus {
			t.Errorf("Expected %d, but got %d", tc.expStatus, w.Code)
		}
	}
}

func Test_Get(t *testing.T) {
	book1 := assignment8.Book{
		ID:     3,
		Name:   "two",
		Author: "three",
	}

	book2 := assignment8.Book{
		ID:     2,
		Name:   "hiii",
		Author: "run",
	}
	createBookBody, err := json.Marshal(book1)
	if err != nil {
		t.Fatalf("error while getting create book body. err : %v", err)
	}
	book2Body, err := json.Marshal(book2)
	if err != nil {
		t.Fatalf("error while getting book2 body. err : %v", err)
	}

	testcases := []struct {
		name      string
		method    string
		expected  []byte
		expStatus int
		id        string
	}{
		{
			name:      "success with book ID",
			expected:  createBookBody,
			method:    http.MethodGet,
			expStatus: http.StatusOK,
			id:        "3",
		},
		{
			name:      "success with second book ID",
			expected:  book2Body,
			method:    http.MethodGet,
			expStatus: http.StatusOK,
			id:        "2",
		},
		{
			name:      "book not found",
			expected:  []byte("Book not found"),
			method:    http.MethodGet,
			expStatus: http.StatusNotFound,
			id:        "100", // A non-existing ID
		},
		{
			name:      "invalid method",
			expected:  []byte("Invalid method"),
			method:    http.MethodPost, // Incorrect method
			expStatus: http.StatusMethodNotAllowed,
			id:        "3",
		},
		{
			name:      "invalid ID format",
			expected:  []byte("Invalid ID format"),
			method:    http.MethodGet,
			expStatus: http.StatusBadRequest,
			id:        "abc", // Non-numeric ID
		},
	}

	for _, tc := range testcases {
		assignment8.Database = map[int]assignment8.Book{
			3: book1,
			2: book2,
		}
		req := httptest.NewRequest(tc.method, "/book/"+tc.id, nil)
		w := httptest.NewRecorder()

		Get(w, req)

		resp := w.Result()
		fmt.Println(resp)

		b, _ := io.ReadAll(resp.Body)
		fmt.Println(resp)

		fmt.Println(string(b))
		fmt.Println(string(tc.expected))

		if !strings.Contains(string(b), string(tc.expected)) {
			t.Errorf("Expected %s, but got %s", tc.expected, string(b))
		}

		if w.Code != tc.expStatus {
			t.Errorf("Expected %d, but got %d", tc.expStatus, w.Code)
		}
	}
}

func TestPut(t *testing.T) {
	updateBook := assignment8.Book{
		ID:     3,
		Name:   "two",
		Author: "three",
	}
	book2 := assignment8.Book{
		ID:     3,
		Name:   "two",
		Author: "three",
	}

	updateBookBody, err := json.Marshal(updateBook)
	if err != nil {
		t.Fatalf("error while getting create book body. err : %v", err)
	}

	Book2Body, err := json.Marshal(book2)
	if err != nil {
		t.Fatalf("error while getting create book body. err : %v", err)
	}

	testCases := []struct {
		name      string
		method    string
		expected  []byte
		expStatus int
		id        string
		body      []byte
	}{
		{
			"success with name",
			http.MethodPut,
			updateBookBody,
			http.StatusOK,
			"3",
			updateBookBody,
		},
		{
			name:      "success with different name",
			expected:  Book2Body,
			method:    http.MethodPut,
			expStatus: http.StatusOK,
			id:        "4",
			body:      Book2Body,
		},
		{
			name:      "book not found",
			expected:  []byte("Failed to update the book: book not found"),
			method:    http.MethodPut,
			expStatus: http.StatusNotFound,
			id:        "100", // A non-existing book ID
			body:      updateBookBody,
		},
		{
			name:      "invalid method",
			expected:  []byte("Invalid method"),
			method:    http.MethodGet, // Incorrect method
			expStatus: http.StatusMethodNotAllowed,
			id:        "3",
			body:      updateBookBody,
		},
		{
			name:      "invalid ID format",
			expected:  []byte("Invalid ID format"),
			method:    http.MethodPut,
			expStatus: http.StatusBadRequest,
			id:        "abc", // Invalid ID format
			body:      updateBookBody,
		},
	}

	for _, tc := range testCases {
		assignment8.Database = map[int]assignment8.Book{
			3: updateBook,
			4: book2,
		}
		body := bytes.NewBuffer(tc.body)

		req := httptest.NewRequest(tc.method, "/book/"+tc.id, body)
		w := httptest.NewRecorder()

		Put(w, req)

		resp := w.Result()
		b, _ := io.ReadAll(resp.Body)

		if !strings.Contains(string(b), string(tc.expected)) {
			t.Errorf("Expected %s, but got %s", tc.expected, string(b))
		}

		if w.Code != tc.expStatus {
			t.Errorf("Expected %d, but got %d", tc.expStatus, w.Code)
		}
	}
}

func TestDelete(t *testing.T) {
	book := assignment8.Book{
		ID:     3,
		Name:   "two",
		Author: "three",
	}
	book2 := assignment8.Book{
		ID:     4,
		Name:   "jjj",
		Author: "xxxxx",
	}

	testCases := []struct {
		name      string
		method    string
		expected  []byte
		expStatus int
		id        string
	}{
		{
			"success with name",
			http.MethodDelete,
			[]byte("done"),
			http.StatusNoContent,
			"3",
		},
		{
			name:      "success with different name",
			expected:  []byte("done"),
			method:    http.MethodDelete,
			expStatus: http.StatusNoContent,
			id:        "4",
		},
		{
			name:      "book not found",
			expected:  []byte("Failed to delete the book: book not found"),
			method:    http.MethodDelete,
			expStatus: http.StatusNotFound,
			id:        "100", // Non-existing book
		},
		{
			name:      "invalid method",
			expected:  []byte("Invalid method"),
			method:    http.MethodGet, // Incorrect method
			expStatus: http.StatusMethodNotAllowed,
			id:        "3",
		},
		{
			name:      "invalid ID format",
			expected:  []byte("Invalid ID format"),
			method:    http.MethodDelete,
			expStatus: http.StatusBadRequest,
			id:        "abc", // Invalid ID
		},
	}

	for _, tc := range testCases {
		assignment8.Database = map[int]assignment8.Book{
			3: book,
			4: book2,
		}

		req := httptest.NewRequest(tc.method, "/book/"+tc.id, http.NoBody)
		w := httptest.NewRecorder()

		Delete(w, req)

		resp := w.Result()
		b, _ := io.ReadAll(resp.Body)

		if !strings.Contains(string(b), string(tc.expected)) {
			t.Errorf("Expected %s, but got %s", tc.expected, string(b))
		}

		if w.Code != tc.expStatus {
			t.Errorf("Expected %d, but got %d", tc.expStatus, w.Code)
		}
	}
}
