package public

import (
	"encoding/json"
	"log"
	"net/http"
	"shopping/models"
)

type Repository interface {
	ListOfProductsForPublic() (models.Products, error)
}

type HandlerListOfProductsForPublic struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerListOfProductsForPublic {
	result := &HandlerListOfProductsForPublic{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerListOfProductsForPublic) sendResponse(writer http.ResponseWriter, listOfProducts models.Products) {
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

func (handler *HandlerListOfProductsForPublic) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	listOfProducts, err := handler.repo.ListOfProductsForPublic()
	if err != nil {
		handler.log.Printf("can not return list of products for public: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(writer, listOfProducts)
}
