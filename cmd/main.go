package main

import (
	"context"
	"fmt"
	"log"
	"mvm_backend/internal/app/mvm"
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/mw"
	"mvm_backend/internal/pkg/service"
	"mvm_backend/internal/pkg/store"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting MVM server  ...")

	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	ctx := context.Background()
	repository := store.NewMVMRepository(ctx, os.Getenv("MONGO_DB_USERNAME"), os.Getenv("MONGO_DB_PASSWORD"))
	jwt_manager := jwt_manager.NewAuthService(os.Getenv("JWT_SECRET"), os.Getenv("JWT_REFRESH_SECRET"), 12*time.Hour, 3600*time.Hour)
	service := service.NewMVMService(repository, jwt_manager)
	mvmServer := mvm.NewIMVMServiceServer(service)

	router := mux.NewRouter()
	SetupRouter(router, mvmServer, jwt_manager)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Create a new logging handler using the os.Stdout logger
	loggingHandler := handlers.LoggingHandler(os.Stdout, router)

	fmt.Println("Server is listening at PORT : " + port)
	if err = http.ListenAndServe(":"+port, loggingHandler); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func SetupRouter(router *mux.Router, mvmServer *mvm.MVMServiceServer, jwt_manager service.IMVMAuth) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "MVM Server") }).Methods("GET")

	router.HandleFunc("/login", mvmServer.LoginUser).Methods("POST")
	router.HandleFunc("/refresh_token", mvmServer.LoginByRefreshToken).Methods("POST")
	router.HandleFunc("/create", mvmServer.CreateUser).Methods("POST")

	router.HandleFunc("/websocket", mvmServer.HandleConnections)
	go mvmServer.HandleMessages()

	userGroup := router.PathPrefix("/user").Subrouter()
	userGroup.Use(mw.MyMiddleware(jwt_manager))
	userGroup.HandleFunc("/", mvmServer.GetProfile).Methods("GET")
	userGroup.HandleFunc("/get", mvmServer.GetUserByUsername).Methods("GET")
	userGroup.HandleFunc("/search", mvmServer.SearchForUsers).Methods("POST")

	friendsGroup := router.PathPrefix("/friends").Subrouter()
	friendsGroup.Use(mw.MyMiddleware(jwt_manager))
	friendsGroup.HandleFunc("", mvmServer.GetFriends).Methods("GET")
	friendsGroup.HandleFunc("/send", mvmServer.CreateFriendRequest).Methods("POST")
	friendsGroup.HandleFunc("/ignore", mvmServer.DeleteFriendRequest).Methods("POST")
	friendsGroup.HandleFunc("/accept", mvmServer.AddFriend).Methods("POST")
	friendsGroup.HandleFunc("/delete", mvmServer.DeleteFriend).Methods("POST")
}
