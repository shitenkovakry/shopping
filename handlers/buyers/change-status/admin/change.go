package admin

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"shopping/models"
)

type ChangeStatusOfBuyer struct {
	ID     int    `json:"id_buyer"`
	Status string `json:"new_status"`
}

type Repository interface {
	ChangeStatuslOfBuyer(idBuyer int, status string) (*models.Buyer, error)
}

type HandlerChangeStatusBuyerForPublic struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerChangeStatusBuyerForPublic {
	result := &HandlerChangeStatusBuyerForPublic{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerChangeStatusBuyerForPublic) prepareRequest(request *http.Request) (*models.Buyer, error) {
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

	var newStatusFromAdmin *ChangeStatusOfBuyer

	if err := json.Unmarshal(body, &newStatusFromAdmin); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	newStatus := &models.Buyer{
		ID:     newStatusFromAdmin.ID,
		Status: newStatusFromAdmin.Status,
	}

	return newStatus, nil
}

func (handler *HandlerChangeStatusBuyerForPublic) sendResponse(writer http.ResponseWriter, changedStatus *models.Buyer) {
	changedStatusOfBuyerMarshaled, err := json.Marshal(changedStatus)
	if err != nil {
		handler.log.Printf("can not marshal changed status of buyer: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(changedStatusOfBuyerMarshaled); err != nil {
		handler.log.Printf("can not changed status: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerChangeStatusBuyerForPublic) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldChangeStatusOfBuyer, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("can not prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	changedStatus, err := handler.repo.ChangeStatuslOfBuyer(shouldChangeStatusOfBuyer.ID, shouldChangeStatusOfBuyer.Status)
	if err != nil {
		handler.log.Printf("can not chage status: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
	handler.sendResponse(writer, changedStatus)
}
