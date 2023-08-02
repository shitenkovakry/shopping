package admin

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"shopping/models"
)

type NewProduct struct {
	ID     int     `json:"id_product"`
	Name   string  `json:"name_product"`
	Price  float64 `json:"price_product"`
	Status string  `json:"status_product"`
}

type Repository interface {
	Create(newProduct *models.Product) (*models.Product, error)
}

type HandlerAddProduct struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerAddProduct {
	result := &HandlerAddProduct{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerAddProduct) prepareRequest(request *http.Request) (*models.Product, error) {
	defer func() {
		if err := request.Body.Close(); err != nil {
			handler.log.Printf("can not close body: %v", err)
		}
	}()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("can not read body: %v", err)

		return nil, err
	}

	var newProductFromAdmin *NewProduct

	if err := json.Unmarshal(body, &newProductFromAdmin); err != nil {
		handler.log.Printf("can not unmarshal body = %s: %v", string(body), err)

		return nil, err
	}

	newProduct := &models.Product{
		Name:   newProductFromAdmin.Name,
		Price:  newProductFromAdmin.Price,
		Status: newProductFromAdmin.Status,
	}

	return newProduct, nil
}

func (handler *HandlerAddProduct) sendResponse(write http.ResponseWriter, createdProduct *models.Product) {
	createdProductMarshaled, err := json.Marshal(createdProduct)
	if err != nil {
		handler.log.Printf("can not marshal created product: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := write.Write(createdProductMarshaled); err != nil {
		handler.log.Printf("can not send created product: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerAddProduct) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	newProduct, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("cannot prepare request: %v", err)
		write.WriteHeader(http.StatusBadRequest)

		return
	}

	createdProduct, err := handler.repo.Create(newProduct)
	if err != nil {
		handler.log.Printf("cannot create user: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(write, createdProduct)
}
