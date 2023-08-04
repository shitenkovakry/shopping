package admin

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"shopping/models"
)

type ChangeNameOfProduct struct {
	ID   int    `json:"id_roduct"`
	Name string `json:"new_name"`
}

type Repository interface {
	ChangeName(idProduct int, name string) (*models.Product, error)
}

type HandlerChangeProductForAdmin struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerChangeProductForAdmin {
	result := &HandlerChangeProductForAdmin{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerChangeProductForAdmin) prepareRequest(request *http.Request) (*models.Product, error) {
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

	var newNameFromAdmin *ChangeNameOfProduct

	if err := json.Unmarshal(body, &newNameFromAdmin); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	newName := &models.Product{
		ID:   newNameFromAdmin.ID,
		Name: newNameFromAdmin.Name,
	}

	return newName, nil
}

func (handler *HandlerChangeProductForAdmin) sendResponse(writer http.ResponseWriter, changedName *models.Product) {
	changedNameProductMarshaled, err := json.Marshal(changedName)
	if err != nil {
		handler.log.Printf("can not marshal changed name of product: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(changedNameProductMarshaled); err != nil {
		handler.log.Printf("can not changed name: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerChangeProductForAdmin) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldChangeNameOfProduct, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("can not prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	changedName, err := handler.repo.ChangeName(shouldChangeNameOfProduct.ID, shouldChangeNameOfProduct.Name)
	if err != nil {
		handler.log.Printf("can not chage name: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
	handler.sendResponse(writer, changedName)
}
