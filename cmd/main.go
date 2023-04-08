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
	"os"
	"time"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	router.POST("/login", mvmServer.LoginUser)
	router.POST("/refresh_token", mvmServer.LoginByRefreshToken)
	router.POST("/create", mvmServer.CreateUser)

	userGroup := router.Group("/user")
	userGroup.Use(mw.AuthorizeJWT(jwt_manager))
	userGroup.POST("/friends/delete", mvmServer.DeleteFriend)
	userGroup.POST("/friends/add", mvmServer.AddFriend)
	userGroup.POST("/search", mvmServer.SearchForUsers)
	userGroup.GET("/get", mvmServer.GetUserByUsername)
	userGroup.GET("/", mvmServer.GetProfile)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
