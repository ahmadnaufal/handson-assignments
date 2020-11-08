package currencyconverter

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CurrencyHandler struct{}

func NewCurrencyHandler() CurrencyHandler {
	return CurrencyHandler{}
}

func (h *CurrencyHandler) GetCurrencies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	currencies := []Currency{
		{
			ID:   1,
			Name: "JPY",
		},
	}

	return OK(w, currencies)
}

func (h *CurrencyHandler) CreateCurrency(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	var currency Currency
	if err := json.NewDecoder(r.Body).Decode(&currency); err != nil {
		return Error(w, err)
	}

	return Created(w, currency)
}

func (h *CurrencyHandler) CreateNewConversionRate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	var crate ConversionRate
	if err := json.NewDecoder(r.Body).Decode(&crate); err != nil {
		return Error(w, err)
	}

	return Created(w, crate)
}

func (h *CurrencyHandler) ConvertCurrency(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	var cparam ConversionParam
	if err := json.NewDecoder(r.Body).Decode(&cparam); err != nil {
		return Error(w, err)
	}

	return Created(w, cparam)
}
