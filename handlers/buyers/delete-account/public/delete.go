package public

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"shopping/errordefs"
	"shopping/models"
)

type DeleteAccount struct {
	ID int `json:"id_buyer"`
}

type Repository interface {
	DeleteAccount(idBuyer int) (*models.Buyer, error)
}

type HandlerDeleteAccountForPublic struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerDeleteAccountForPublic {
	result := &HandlerDeleteAccountForPublic{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerDeleteAccountForPublic) prepareRequest(request *http.Request) (*models.Buyer, error) {
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

	var deleteAccountFromClient *DeleteAccount

	if err := json.Unmarshal(body, &deleteAccountFromClient); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	deletedBuyer := &models.Buyer{
		ID: deleteAccountFromClient.ID,
	}

	return deletedBuyer, nil
}

func (handler *HandlerDeleteAccountForPublic) sendResponse(writer http.ResponseWriter, deletedBuyer *models.Buyer) {
	deletedAccountMarshaled, err := json.Marshal(deletedBuyer)
	if err != nil {
		handler.log.Printf("can not marshal deleted account: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(deletedAccountMarshaled); err != nil {
		handler.log.Printf("can not send to client deleted account: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerDeleteAccountForPublic) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldDeleteBuyer, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("can not prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	deletedAccount, err := handler.repo.DeleteAccount(shouldDeleteBuyer.ID)
	if errors.Is(err, errordefs.ErrNoDocuments) {
		writer.WriteHeader(http.StatusNotFound)

		return
	}

	if err != nil {
		handler.log.Printf("can not delete buyer: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(writer, deletedAccount)
}
