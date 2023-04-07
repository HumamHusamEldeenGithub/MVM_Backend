package main

import (
	"context"
	"fmt"
	"log"
	"mvm_backend/internal/app/mvm"
	v1 "mvm_backend/internal/pkg/generated/mvm-api/v1"
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/mw"
	"mvm_backend/internal/pkg/service"
	"mvm_backend/internal/pkg/store"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port : 8082 ...")

	// Start our listener, 8082 is the default gRPC port
	listener, err := net.Listen("tcp", ":8082")
	// Handle errors if any
	if err != nil {
		log.Fatalf("Unable to listen on port :8082: %v", err)
	}

	ctx := context.Background()

	// GET FROM ENV VAR
	repository := store.NewMVMRepository(ctx, "Dev", "6EHO7HJ9Zr2bG1Uw")
	jwt_manager := jwt_manager.NewAuthService("secret", 12*time.Second)
	service := service.NewMVMService(repository, jwt_manager)
	mvmServer := mvm.NewIMVMServiceServer(service)

	// Set options, here we can configure things like TLS support
	opts := []grpc.ServerOption{grpc.UnaryInterceptor(mw.TokenAuthorizer)}
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)
	// Register the service with the server
	v1.RegisterMVMServiceServer(s, mvmServer)

	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :8082")
	defer ShutDownServer(s, listener)
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
