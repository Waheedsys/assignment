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

func Test_Getbook(t *testing.T) {

	book := Book{
		Id:     3,
		Title:  "two",
		Author: "three",
	}
	createBookBody, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("error while getting create book body. err : %v", err)
	}
	//testcases
	testcases := []struct {
		name      string
		method    string
		expected  []byte
		expStatus int
		id        string
	}{
		{name: "success with book ID",
			expected:  createBookBody,
			method:    http.MethodGet,
			expStatus: http.StatusOK,
			id:        "3"},
	}

}
