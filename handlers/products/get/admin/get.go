package admin

import (
	"encoding/json"
	"log"
	"net/http"
	"shopping/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Repository interface {
	GetProduct(idProduct int) (*models.Product, error)
}

type HandlerGetProductForAdmin struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerGetProductForAdmin {
	result := &HandlerGetProductForAdmin{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerGetProductForAdmin) prepareRequest(request *http.Request) (int, error) {
	productIDParam := chi.URLParam(request, "id_product")
	productID, err := strconv.Atoi(productIDParam)

	if err != nil {
		handler.log.Printf("err = %v", err)

		return 0, err
	}

	return productID, nil
}

func (handler *HandlerGetProductForAdmin) sendResponse(writer http.ResponseWriter, gotProduct *models.Product) {
	gotProductMarshaled, err := json.Marshal(gotProduct)
	if err != nil {
		handler.log.Printf("can not marshal got product: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(gotProductMarshaled); err != nil {
		handler.log.Printf("can not send got product: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerGetProductForAdmin) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldSendProductID, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("can not prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	gotProduct, err := handler.repo.GetProduct(shouldSendProductID)
	if err != nil {
		handler.log.Printf("can not get product: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(writer, gotProduct)
}
