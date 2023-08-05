package public

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"shopping/models"
)

type ChangeNameOfBuyer struct {
	ID   int    `json:"id_buyer"`
	Name string `json:"new_name"`
}

type Repository interface {
	ChangeNameOfBuyer(idBuyer int, name string) (*models.Buyer, error)
}

type HandlerChangeNameBuyerForAdmin struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerChangeNameBuyerForAdmin {
	result := &HandlerChangeNameBuyerForAdmin{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerChangeNameBuyerForAdmin) prepareRequest(request *http.Request) (*models.Buyer, error) {
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

	var newNameFromAdmin *ChangeNameOfBuyer

	if err := json.Unmarshal(body, &newNameFromAdmin); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	newName := &models.Buyer{
		ID:   newNameFromAdmin.ID,
		Name: newNameFromAdmin.Name,
	}

	return newName, nil
}

func (handler *HandlerChangeNameBuyerForAdmin) sendResponse(writer http.ResponseWriter, changedName *models.Buyer) {
	changedNameBuyerMarshaled, err := json.Marshal(changedName)
	if err != nil {
		handler.log.Printf("can not marshal changed name of buyer: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(changedNameBuyerMarshaled); err != nil {
		handler.log.Printf("can not changed name: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerChangeNameBuyerForAdmin) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldChangeNameOfBuyer, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("can not prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	changedName, err := handler.repo.ChangeNameOfBuyer(shouldChangeNameOfBuyer.ID, shouldChangeNameOfBuyer.Name)
	if err != nil {
		handler.log.Printf("can not chage name: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
	handler.sendResponse(writer, changedName)
}
