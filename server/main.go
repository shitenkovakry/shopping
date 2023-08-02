package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"shopping/db"
	"syscall"
	"time"

	products_get_admin "shopping/handlers/products/get/admin"
	products_list_admin "shopping/handlers/products/list/admin"
	products_list_public "shopping/handlers/products/list/public"

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

	handlerListOfProductsForAdmin := products_list_admin.New(productsRepo, log)
	router.Method(http.MethodGet, "/list/products/admin", handlerListOfProductsForAdmin)
	handlerListOfProductsForPublic := products_list_public.New(productsRepo, log)
	router.Method(http.MethodGet, "/list/products/public", handlerListOfProductsForPublic)
	handlerGetProductForAdmin := products_get_admin.New(productsRepo, log)
	router.Method(http.MethodGet, "/get/product/{id_product}/admin", handlerGetProductForAdmin)

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
