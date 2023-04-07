package main

import (
	"context"
	"fmt"
	"log"
	"mvm_backend/internal/app/mvm"
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/service"
	"mvm_backend/internal/pkg/store"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting MVM server  ...")

	ctx := context.Background()

	// GET FROM ENV VAR
	repository := store.NewMVMRepository(ctx, "Dev", "6EHO7HJ9Zr2bG1Uw")
	jwt_manager := jwt_manager.NewAuthService("secret", 12*time.Hour, 3600*time.Hour)
	service := service.NewMVMService(repository, jwt_manager)
	mvmServer := mvm.NewIMVMServiceServer(service)

	router := gin.Default()

	router.POST("/login", mvmServer.LoginUser)
	router.POST("/user", mvmServer.CreateUser)
	router.GET("/user", mvmServer.GetUser)

	router.Run()

}

func ShutDownServer(s *grpc.Server, listener net.Listener) {
	// Right way to stop the server using a SHUTDOWN HOOK
	// Create a channel to receive OS signals
	c := make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Closing MongoDB connection")
	// db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}
