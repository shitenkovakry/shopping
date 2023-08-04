package admin

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"shopping/errordefs"
	"shopping/models"
)

type DeleteProduct struct {
	ID int `json:"id_product"`
}

type Repository interface {
	Delete(idProduct int) (*models.Product, error)
}

type HandlerDeleteProductForAdmin struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerDeleteProductForAdmin {
	result := &HandlerDeleteProductForAdmin{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerDeleteProductForAdmin) prepareRequest(request *http.Request) (*models.Product, error) {
	defer func() {
		if err := request.Body.Close(); err != nil {
			handler.log.Printf("cannot close body: %v", err)
		}
	}()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("cannot read body: %v", err)

		return nil, err
	}

	var deleteProductFromClient *DeleteProduct

	if err := json.Unmarshal(body, &deleteProductFromClient); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	deletedProduct := &models.Product{
		ID: deleteProductFromClient.ID,
	}

	return deletedProduct, nil
}

func (handler *HandlerDeleteProductForAdmin) sendResponse(writer http.ResponseWriter, deletedProduct *models.Product) {
	deletedProductMarshaled, err := json.Marshal(deletedProduct)
	if err != nil {
		handler.log.Printf("can not marshal deleted product: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(deletedProductMarshaled); err != nil {
		handler.log.Printf("can not send to client deleted product: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerDeleteProductForAdmin) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldDeleteProduct, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("can not prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	deletedProduct, err := handler.repo.Delete(shouldDeleteProduct.ID)
	if errors.Is(err, errordefs.ErrNoDocuments) {
		writer.WriteHeader(http.StatusNotFound)

		return
	}

	if err != nil {
		handler.log.Printf("can not delete product: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(writer, deletedProduct)
}
