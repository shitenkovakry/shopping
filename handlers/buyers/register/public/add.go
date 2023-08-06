package admin

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"shopping/models"
)

type NewBuyer struct {
	ID      int     `json:"id_buyer"`
	Name    string  `json:"name_buyer"`
	Email   string  `json:"email_buyer"`
	Balance float64 `json:"balance_buyer"`
	Status  string  `json:"status_buyer"`
}

type Repository interface {
	Register(newBuyer *models.Buyer) (*models.Buyer, error)
}

type HandlerRegisterBuyer struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerRegisterBuyer {
	result := &HandlerRegisterBuyer{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerRegisterBuyer) prepareRequest(request *http.Request) (*models.Buyer, error) {
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

	var newBuyerToRegistration *NewBuyer

	if err := json.Unmarshal(body, &newBuyerToRegistration); err != nil {
		handler.log.Printf("can not unmarshal body = %s: %v", string(body), err)

		return nil, err
	}

	newBuyer := &models.Buyer{
		Name:    newBuyerToRegistration.Name,
		Email:   newBuyerToRegistration.Email,
		Balance: newBuyerToRegistration.Balance,
		Status:  newBuyerToRegistration.Status,
	}

	return newBuyer, nil
}

func (handler *HandlerRegisterBuyer) sendResponse(write http.ResponseWriter, createdBuyer *models.Buyer) {
	createdBuyerMarshaled, err := json.Marshal(createdBuyer)
	if err != nil {
		handler.log.Printf("can not marshal created buyer: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := write.Write(createdBuyerMarshaled); err != nil {
		handler.log.Printf("can not send created buyer: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerRegisterBuyer) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	newBuyer, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("cannot prepare request: %v", err)
		write.WriteHeader(http.StatusBadRequest)

		return
	}

	createdBuyer, err := handler.repo.Register(newBuyer)
	if err != nil {
		handler.log.Printf("cannot create user: %v", err)
		write.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(write, createdBuyer)
}
