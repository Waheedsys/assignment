package main

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

// post
func Test_Addbook(t *testing.T) {
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
	store := NewStore(db)

	testcases := []struct {
		name      string
		method    string
		expected  string
		expStatus int
		id        int
		body      []byte
	}{
		{
			name:      "successfully created",
			method:    http.MethodPost,
			id:        4,
			expected:  `{"id":24,"author":"three","title":"two"}`,
			expStatus: http.StatusOK,
			body:      []byte(`{"id":24,"author":"three","title":"two"}`),
		},
	}

	for _, tc := range testcases {
		body := bytes.NewBuffer(tc.body)

		req := httptest.NewRequest(tc.method, "/book", body)
		w := httptest.NewRecorder()
		store.Addbook(w, req)
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

// get
func Test_Getbook(t *testing.T) {
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
	store := NewStore(db)

	//testcases
	testcases := []struct {
		name      string
		method    string
		expected  string
		expStatus int
		id        int
	}{
		{
			name:      "successfully get a book",
			method:    http.MethodGet,
			id:        4,
			expected:  `{"id":4,"author":"three","title":"two"}`,
			expStatus: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		url := "/book/" + strconv.Itoa(tc.id)
		req := httptest.NewRequest(tc.method, url, nil)
		w := httptest.NewRecorder()
		store.Getbook(w, req)
		resp := w.Result()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Failed to read response body: %v", err)
			return
		}
		if !strings.Contains(string(b), string(tc.expected)) {
			t.Errorf("Expected body %q, but got %q", tc.expected, string(b))
		}

	}

}

// put
func Test_UpdateBook(t *testing.T) {
	testCases := []struct {
		name      string
		method    string
		id        int
		body      string
		expected  string
		expStatus int
	}{
		{name: "successfully updated a book",
			method:    http.MethodPut,
			id:        7,
			body:      `{"id":7,"author": "updated author","title":"updated title"}`,
			expected:  `{"id":7,"author":"updated author","title":"updated title"}` + "\n",
			expStatus: http.StatusOK,
		},
	}

	db, err := sql.Open("mysql", "root:password@tcp/sample")
	if err != nil {
		t.Errorf("database is not connected %v", err)
	}
	store := NewStore(db)
	for _, tc := range testCases {
		url := fmt.Sprintf("/book/%d", tc.id)
		req := httptest.NewRequest(tc.method, url, strings.NewReader(tc.body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		store.Updatebook(w, req)

		resp := w.Result()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		if string(b) != tc.expected {
			t.Errorf("Expected body %q, but got %q", tc.expected, string(b))
		}

	}
}

// delete
func Test_Deletebook(t *testing.T) {
	testCases := []struct {
		name      string
		method    string
		id        int
		expected  string
		expStatus int
	}{
		{
			name:      "successfully deleted a book",
			method:    http.MethodDelete,
			id:        8,
			expected:  "Book with id 8 deleted",
			expStatus: http.StatusNoContent,
		},
	}

	db, err := sql.Open("mysql", "root:password@/sample")
	if err != nil {
		t.Errorf("database is not connected %v", err)
	}
	store := NewStore(db)

	for _, tc := range testCases {

		url := fmt.Sprintf("/book/%d", tc.id)
		req := httptest.NewRequest(tc.method, url, nil)
		w := httptest.NewRecorder()

		store.Deletebook(w, req)

		resp := w.Result()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		if string(b) != tc.expected {
			t.Errorf("Expected body %q, but got %q", tc.expected, string(b))
		}

		if resp.StatusCode != tc.expStatus {
			t.Errorf("Expected status %d, but got %d", tc.expStatus, resp.StatusCode)
		}

	}
}
