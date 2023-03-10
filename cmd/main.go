package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"

	forumHandlers "lonkidely/technopark-dbms-forum/internal/forum/delivery/handlers"
	forumRepository "lonkidely/technopark-dbms-forum/internal/forum/repository"
	forumUsecase "lonkidely/technopark-dbms-forum/internal/forum/usecase"
	postHandlers "lonkidely/technopark-dbms-forum/internal/post/delivery/handlers"
	postRepository "lonkidely/technopark-dbms-forum/internal/post/repository"
	postUsecase "lonkidely/technopark-dbms-forum/internal/post/usecase"
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

func getPostgres() *pgx.ConnPool {
	conn, err := pgx.ParseConnectionString(PostgresURL)
	if err != nil {
		log.Fatalln("Can't parse postgres url", err)
	}

	conn.PreferSimpleProtocol = true

	cfg := pgx.ConnPoolConfig{
		ConnConfig:     conn,
		MaxConnections: 200,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	}

	pool, err := pgx.NewConnPool(cfg)
	if err != nil {
		log.Fatalln("error creating pool", err)
	}

	return pool
}

func main() {
	router := mux.NewRouter()

	postgres := getPostgres()

	// Repositories
	userRepo := userRepository.NewUserRepository(postgres)
	forumRepo := forumRepository.NewForumRepository(postgres)
	serviceRepo := serviceRepository.NewServiceRepository(postgres)
	threadRepo := threadRepository.NewThreadRepository(postgres)
	postRepo := postRepository.NewPostRepository(postgres)

	// Usecases
	userUse := userUsecase.NewUserUsecase(userRepo)
	forumUse := forumUsecase.NewForumUsecase(forumRepo, userRepo)
	serviceUse := serviceUsecase.NewServiceUsecase(serviceRepo)
	threadUse := threadUsecase.NewThreadUsecase(threadRepo, forumRepo, userRepo)
	postUse := postUsecase.NewPostUsecase(postRepo)

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

	getForumUsersHandler := forumHandlers.NewGetForumUsersHandler(forumUse)
	getForumUsersHandler.Configure(router)

	clearHandler := serviceHandlers.NewClearHandler(serviceUse)
	clearHandler.Configure(router)

	statusHandler := serviceHandlers.NewStatusHandler(serviceUse)
	statusHandler.Configure(router)

	createThreadHandler := threadHandlers.NewCreateThreadHandler(threadUse)
	createThreadHandler.Configure(router)

	getThreadDetailsHandler := threadHandlers.NewGetThreadDetailsHandler(threadUse)
	getThreadDetailsHandler.Configure(router)

	updateThreadDetailsHandler := threadHandlers.NewUpdateThreadDetailsHandler(threadUse)
	updateThreadDetailsHandler.Configure(router)

	voteThreadHandler := threadHandlers.NewVoteThreadHandler(threadUse)
	voteThreadHandler.Configure(router)

	getThreadPostsHandler := threadHandlers.NewGetThreadPostsHandler(threadUse)
	getThreadPostsHandler.Configure(router)

	createPostsHandler := threadHandlers.NewCreatePostsHandler(threadUse)
	createPostsHandler.Configure(router)

	getPostDetailsHanlder := postHandlers.NewGetPostDetailsHandler(postUse)
	getPostDetailsHanlder.Configure(router)

	updatePostHandler := postHandlers.NewUpdatePostHandler(postUse)
	updatePostHandler.Configure(router)

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
