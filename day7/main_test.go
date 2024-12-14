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
		ID:     0,
		Name:   "two",
		Author: "three",
	}

	createBookBody, err := json.Marshal(createBook)
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
			name:      "success with name",
			body:      createBookBody,
			expected:  createBookBody,
			method:    http.MethodPost,
			expStatus: http.StatusCreated,
		},
		{
			name:      "success with different name",
			expected:  createBookBody,
			method:    http.MethodPost,
			expStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		body := bytes.NewBuffer(tc.body)

		req := httptest.NewRequest(tc.method, "/book", body)
		w := httptest.NewRecorder()

		Post(w, req)

		resp := w.Result()
		b, _ := io.ReadAll(resp.Body)

		if string(b) != string(tc.expected) {
			t.Errorf("Expected %s, but got %s", tc.expected, string(b))
		}

		if w.Code != tc.expStatus {
			t.Errorf("Expected %d, but got %d", tc.expStatus, w.Code)
		}
	}
}

func Test_Get(t *testing.T) {
	createbook := assignment8.Book{
		ID:     3,
		Name:   "two",
		Author: "three",
	}

	createBookBody, err := json.Marshal(createbook)
	if err != nil {
		t.Fatalf("error while getting create book body. err : %v", err)
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
	}

	for _, tc := range testcases {
		assignment8.Database = map[int]assignment8.Book{3: createbook}

		req := httptest.NewRequest(tc.method, "/book/"+tc.id, nil)
		w := httptest.NewRecorder()

		Get(w, req)

		resp := w.Result()
		b, _ := io.ReadAll(resp.Body)

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

	updateBookBody, err := json.Marshal(updateBook)
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
		//{
		//	name:      "success with different name",
		//	expected:  updateBookBody,
		//	method:    http.MethodPut,
		//	expStatus: http.StatusOK,
		//},
	}

	for _, tc := range testCases {
		assignment8.Database = map[int]assignment8.Book{3: updateBook}

		body := bytes.NewBuffer(tc.body)

		req := httptest.NewRequest(tc.method, "/book/update/"+tc.id, body)
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
		//{
		//	name:      "success with different name",
		//	expected:  updateBookBody,
		//	method:    http.MethodPut,
		//	expStatus: http.StatusOK,
		//},
	}

	for _, tc := range testCases {
		assignment8.Database = map[int]assignment8.Book{3: book}

		req := httptest.NewRequest(tc.method, "/book/delete/"+tc.id, http.NoBody)
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
