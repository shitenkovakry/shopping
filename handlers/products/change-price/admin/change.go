package admin

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"shopping/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ChangePriceOfProduct struct {
	Price float64 `json:"new_price"`
}

type Repository interface {
	ChangePrice(idProduct int, price float64) (*models.Product, error)
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

	productIDParam := chi.URLParam(request, "id_product")
	productID, err := strconv.Atoi(productIDParam)

	if err != nil {
		handler.log.Printf("err = %v", err)

		return nil, err
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		handler.log.Printf("cannot read body: %v", err)

		return nil, err
	}

	var newPriceFromAdmin *ChangePriceOfProduct

	if err := json.Unmarshal(body, &newPriceFromAdmin); err != nil {
		handler.log.Printf("cannot unmarshal body=%s: %v", string(body), err)

		return nil, err
	}

	newPrice := &models.Product{
		ID:    productID,
		Price: newPriceFromAdmin.Price,
	}

	return newPrice, nil
}

func (handler *HandlerChangeProductForAdmin) sendResponse(writer http.ResponseWriter, changedPrice *models.Product) {
	changedPriceProductMarshaled, err := json.Marshal(changedPrice)
	if err != nil {
		handler.log.Printf("can not marshal changed price of product: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(changedPriceProductMarshaled); err != nil {
		handler.log.Printf("can not changed price: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *HandlerChangeProductForAdmin) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	shouldChangePriceOfProduct, err := handler.prepareRequest(request)
	if err != nil {
		handler.log.Printf("can not prepare request: %v", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	changedPrice, err := handler.repo.ChangePrice(shouldChangePriceOfProduct.ID, shouldChangePriceOfProduct.Price)
	if err != nil {
		handler.log.Printf("can not chage price: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	handler.sendResponse(writer, changedPrice)
}
