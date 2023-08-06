package admin

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"shopping/models"
)

type ChangeStatusOfProduct struct {
	ID     int    `json:"id_roduct"`
	Status string `json:"new_status"`
}

type Repository interface {
	ChangeStatus(idProduct int, status string) (*models.Product, error)
}

type HandlerChangeStatusOfProductForAdmin struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerChangeStatusOfProductForAdmin {
	result := &HandlerChangeStatusOfProductForAdmin{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerChangeStatusOfProductForAdmin) prepareRequest(request *http.Request) (*models.Product, error) {
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

	var newStatusFromAdmin *ChangeStatusOfProduct

	if err := json.Unmarshal(body, &newStatusFromAdmin); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	newStatus := &models.Product{
		ID:     newStatusFromAdmin.ID,
		Status: newStatusFromAdmin.Status,
	}

	return newStatus, nil
}

func (handler *HandlerChangeStatusOfProductForAdmin) sendResponse(writer http.ResponseWriter, changedStatus *models.Product) {
	changedStatusProductMarshaled, err := json.Marshal(changedStatus)
	if err != nil {
		handler.log.Printf("can not marshal changed status of product: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(changedStatusProductMarshaled); err != nil {
		handler.log.Printf("can not changed status: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerChangeStatusOfProductForAdmin) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldChangeStatusOfProduct, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("can not prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	changedStatus, err := handler.repo.ChangeStatus(shouldChangeStatusOfProduct.ID, shouldChangeStatusOfProduct.Status)
	if err != nil {
		handler.log.Printf("can not chage status: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
	handler.sendResponse(writer, changedStatus)
}
