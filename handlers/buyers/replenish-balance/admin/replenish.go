package admin

import (
	"encoding/json"
	"log"
	"net/http"
	"shopping/models"
)

type ReplenishBalanceOfBuyer struct {
	ID             int     `json:"id_buyer"`
	PriceOfProduct float64 `json:"price_product"`
}

type Repository interface {
	ReplenishBalance(idBuyer int, priceOfProduct float64) (*models.Buyer, error)
}

type HandlerReplenishBalanceBuyerForAdmin struct {
	repo Repository
	log  *log.Logger
}

func New(repo Repository, log *log.Logger) *HandlerReplenishBalanceBuyerForAdmin {
	result := &HandlerReplenishBalanceBuyerForAdmin{
		repo: repo,
		log:  log,
	}

	return result
}

func (handler *HandlerReplenishBalanceBuyerForAdmin) prepareRequest(request *http.Request) (*ReplenishBalanceOfBuyer, error) {
	var replenishData ReplenishBalanceOfBuyer

	err := json.NewDecoder(request.Body).Decode(&replenishData)
	if err != nil {
		handler.log.Printf("can not unmarshal body: %v", err)
		return nil, err
	}

	return &replenishData, nil
}

func (handler *HandlerReplenishBalanceBuyerForAdmin) sendResponse(writer http.ResponseWriter, replenishedBalance *models.Buyer) {
	replenishedBalanceOfBuyerMarshaled, err := json.Marshal(replenishedBalance)
	if err != nil {
		handler.log.Printf("can not marshal replenished balance of buyer: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(replenishedBalanceOfBuyerMarshaled); err != nil {
		handler.log.Printf("can not replenished balance: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerReplenishBalanceBuyerForAdmin) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	replenishData, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("can not prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	changedBalanceAfterShopping, err := handler.repo.ReplenishBalance(replenishData.ID, replenishData.PriceOfProduct)
	if err != nil {
		handler.log.Printf("can not change balance: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.sendResponse(writer, changedBalanceAfterShopping)
}
