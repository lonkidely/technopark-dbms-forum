package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/stdlib"

	userHandlers "lonkidely/technopark-dbms-forum/internal/user/delivery/handlers"
	userRepository "lonkidely/technopark-dbms-forum/internal/user/repository"
	userUsecase "lonkidely/technopark-dbms-forum/internal/user/usecase"
)

const (
	PostgresURL = "user=lonkidely dbname=tech_db_forum password=lonkidely host=localhost port=5432"
)

func getPostgres() *sql.DB {
	conn, err := sql.Open("pgx", PostgresURL)
	if err != nil {
		log.Fatalln("Can't parse postgres url", err)
	}

	conn.SetMaxOpenConns(100)

	err = conn.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}

func main() {
	router := mux.NewRouter()

	postgres := getPostgres()

	// Repositories
	userRepo := userRepository.NewUserRepository(postgres)

	// Usecases
	userUse := userUsecase.NewUserUsecase(userRepo)

	// Delivery
	createUserHandler := userHandlers.NewCreateUserHandler(userUse)
	createUserHandler.Configure(router)

	getUserInfoHandler := userHandlers.NewGetUserInfoHandler(userUse)
	getUserInfoHandler.Configure(router)

	updateUserInfoHandler := userHandlers.NewUpdateUserInfoHandler(userUse)
	updateUserInfoHandler.Configure(router)

	server := http.Server{
		Addr:         ":5000",
		Handler:      router,
		ReadTimeout:  time.Duration(10) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
