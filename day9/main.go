package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Waheedsys/entities/handler"
	"github.com/Waheedsys/entities/services"
	"github.com/Waheedsys/entities/stores"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/sample")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	log.Println("Successfully connected to the database!")

	route := mux.NewRouter()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      route,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// route ,creating route using mux
	userstore := stores.NewDetails(db)
	userService := services.NewUserService(userstore)
	userHandler := handler.NewUserHandler(userService)

	route.HandleFunc("/user", userHandler.AddUser).Methods("POST")
	route.HandleFunc("/user", userHandler.GetUsers).Methods("GET")
	route.HandleFunc("/user/{name}", userHandler.GetUserByName).Methods("GET")
	route.HandleFunc("/user/{name}", userHandler.UpdateUser).Methods("PUT")
	route.HandleFunc("/user/{name}", userHandler.DeleteUser).Methods("DELETE")

	fmt.Println("Server started on http://localhost:8080")
	err = http.ListenAndServe(":8080", route)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

	log.Println("Shutting down server...")

	const timeoutDuration = 10 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()
	if err1 := server.Shutdown(ctx); err1 != nil {
		log.Fatalf("Server shutdown failed: %v", err1)
	}

	log.Println("Server gracefully stopped")
}
