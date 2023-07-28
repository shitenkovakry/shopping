package products

import (
	"encoding/json"
	"log"
	"net/http"
	"shopping/models"
)

type Repository interface {
	ListOfProductsForAdmin() (models.Products, error)
}

type HandlerListOfProductsForAdmin struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerListOfProductsForAdmin {
	result := &HandlerListOfProductsForAdmin{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerListOfProductsForAdmin) sendResponse(writer http.ResponseWriter, listOfProducts models.Products) {
	data, err := json.Marshal(listOfProducts)
	if err != nil {
		handler.log.Printf("can not marshal list of products: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(data); err != nil {
		handler.log.Printf("can not write the data to the connection as part of an HTTP reply: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerListOfProductsForAdmin) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	listOfProducts, err := handler.repo.ListOfProductsForAdmin()
	if err != nil {
		handler.log.Printf("can not return list of products for admin: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(writer, listOfProducts)
}
