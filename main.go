package main

import (
	"L0/iternal/repository"
	"L0/iternal/service"
	"L0/iternal/transport"
	"L0/pkg/database"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db, err := postgres.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := repository.NewOrderRepo(db)

	service := service.NewOrderService(repository)

	handler := transport.NewHandler(service)

	server := &http.Server{
		Addr: ":8080",
		Handler: handler.Init(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))
	go func (){
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	
	quit := make(chan os.Signal, 1) 
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Print("failed to stop server:")
	}
}