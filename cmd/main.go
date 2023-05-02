package main

import (
	"L0/iternal/domain"
	"L0/iternal/repository"
	"L0/iternal/service"
	"L0/iternal/transport"
	"L0/pkg/cache"
	"L0/pkg/database"
	"L0/pkg/nats"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/nats-io/stan.go"
)

func main() {

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	db, err := postgres.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := repository.NewOrderRepo(db)

	cache := cache.NewCache()
	repository.RestoreCache(ctx, cache)

	service := service.NewOrderService(repository, cache)

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

	nc, err := nats.Connect("subscriber")
	if err != nil {
		log.Fatal("cant connect to nats", err)
	}
	
	_, err = nc.Subscribe("order", func(m *stan.Msg) {
		fmt.Print(string(m.Data))
		var order domain.Order
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Println("not valid json, cant unmarshal it")
			return
		}
		if err = service.Create(ctx, order.OrderUID, order); err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
		log.Println(err)
	}

	quit := make(chan os.Signal, 1) 
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	
	if err := server.Shutdown(ctx); err != nil {
		log.Print("failed to stop server:")
	}
}