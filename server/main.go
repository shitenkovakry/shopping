package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"shopping/db"
	"syscall"
	"time"

	products "shopping/handlers/products/list/admin"

	"shopping/logger"
	productsRepo "shopping/repositories/products"

	"github.com/go-chi/chi/v5"
)

const (
	address = ":8080"
)

func main() {
	router := chi.NewRouter()
	log := logger.New()

	log.Print("we are going to start")

	dataBase := db.New()

	productsRepo := productsRepo.New(dataBase)

	handlerListOfProductsForAdmin := products.New(productsRepo, log)
	router.Method(http.MethodGet, "/list/products/admin", handlerListOfProductsForAdmin)

	server := NewServer(address, router)

	log.Printf("serving at [%s]", address)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server is error: %v", err)
		}
	}()
	<-stopChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown error: %v", err)
	}
}

func NewServer(address string, router *chi.Mux) *http.Server {
	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}
