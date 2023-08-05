package public

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"shopping/models"
)

type ChangeEmailOfBuyer struct {
	ID    int    `json:"id_buyer"`
	Email string `json:"new_email"`
}

type Repository interface {
	ChangeEmailOfBuyer(idBuyer int, email string) (*models.Buyer, error)
}

type HandlerChangeEmailBuyerForPublic struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerChangeEmailBuyerForPublic {
	result := &HandlerChangeEmailBuyerForPublic{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerChangeEmailBuyerForPublic) prepareRequest(request *http.Request) (*models.Buyer, error) {
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

	var newEmailFromClient *ChangeEmailOfBuyer

	if err := json.Unmarshal(body, &newEmailFromClient); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	newEmail := &models.Buyer{
		ID:    newEmailFromClient.ID,
		Email: newEmailFromClient.Email,
	}

	return newEmail, nil
}

func (handler *HandlerChangeEmailBuyerForPublic) sendResponse(writer http.ResponseWriter, changedEmail *models.Buyer) {
	changedEmailOfBuyerMarshaled, err := json.Marshal(changedEmail)
	if err != nil {
		handler.log.Printf("can not marshal changed email of buyer: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(changedEmailOfBuyerMarshaled); err != nil {
		handler.log.Printf("can not changed email: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerChangeEmailBuyerForPublic) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldChangeEmailOfBuyer, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("can not prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	changedEmail, err := handler.repo.ChangeEmailOfBuyer(shouldChangeEmailOfBuyer.ID, shouldChangeEmailOfBuyer.Email)
	if err != nil {
		handler.log.Printf("can not chage email: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
	handler.sendResponse(writer, changedEmail)
}
