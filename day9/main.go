package main

import (
	"database/sql"
	"fmt"
	"github.com/Waheedsys/entities/handler"
	"github.com/Waheedsys/entities/services"
	"github.com/Waheedsys/entities/stores"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/sample")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
	} else {
		log.Println("Successfully connected to the database!")
	}

	//route ,creating route using mux
	route := mux.NewRouter()
	userstore := stores.NewDetails(db)
	userService := services.NewUserService(*userstore)
	userHandler := handler.NewUserHandler(*userService)

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
}
