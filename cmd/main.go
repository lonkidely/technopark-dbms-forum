package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/stdlib"

	forumHandlers "lonkidely/technopark-dbms-forum/internal/forum/delivery/handlers"
	forumRepository "lonkidely/technopark-dbms-forum/internal/forum/repository"
	forumUsecase "lonkidely/technopark-dbms-forum/internal/forum/usecase"
	serviceHandlers "lonkidely/technopark-dbms-forum/internal/service/delivery/handlers"
	serviceRepository "lonkidely/technopark-dbms-forum/internal/service/repository"
	serviceUsecase "lonkidely/technopark-dbms-forum/internal/service/usecase"
	threadHandlers "lonkidely/technopark-dbms-forum/internal/thread/delivery/handlers"
	threadRepository "lonkidely/technopark-dbms-forum/internal/thread/repository"
	threadUsecase "lonkidely/technopark-dbms-forum/internal/thread/usecase"
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
	forumRepo := forumRepository.NewForumRepository(postgres)
	serviceRepo := serviceRepository.NewServiceRepository(postgres)
	threadRepo := threadRepository.NewThreadRepository(postgres)

	// Usecases
	userUse := userUsecase.NewUserUsecase(userRepo)
	forumUse := forumUsecase.NewForumUsecase(forumRepo, userRepo)
	serviceUse := serviceUsecase.NewServiceUsecase(serviceRepo)
	threadUse := threadUsecase.NewThreadUsecase(threadRepo, forumRepo, userRepo)

	// Delivery
	createUserHandler := userHandlers.NewCreateUserHandler(userUse)
	createUserHandler.Configure(router)

	getUserInfoHandler := userHandlers.NewGetUserInfoHandler(userUse)
	getUserInfoHandler.Configure(router)

	updateUserInfoHandler := userHandlers.NewUpdateUserInfoHandler(userUse)
	updateUserInfoHandler.Configure(router)

	createForumHandler := forumHandlers.NewCreateForumHandler(forumUse)
	createForumHandler.Configure(router)

	forumDetailsHandler := forumHandlers.NewForumDetailsHandler(forumUse)
	forumDetailsHandler.Configure(router)

	getForumThreadsHandler := forumHandlers.NewGetForumThreadsHandler(forumUse)
	getForumThreadsHandler.Configure(router)

	clearHandler := serviceHandlers.NewClearHandler(serviceUse)
	clearHandler.Configure(router)

	statusHandler := serviceHandlers.NewStatusHandler(serviceUse)
	statusHandler.Configure(router)

	createThreadHandler := threadHandlers.NewCreateThreadHandler(threadUse)
	createThreadHandler.Configure(router)

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
