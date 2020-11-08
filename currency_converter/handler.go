package currencyconverter

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CurrencyHandler struct {
	flow CurrencyFlow
}

func NewCurrencyHandler(currencyFlow CurrencyFlow) CurrencyHandler {
	return CurrencyHandler{
		flow: currencyFlow,
	}
}

func (h *CurrencyHandler) GetCurrencies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	ctx := r.Context()
	currencies, err := h.flow.GetCurrencies(ctx)
	if err != nil {
		return Error(w, err)
	}

	return OK(w, currencies)
}

func (h *CurrencyHandler) CreateCurrency(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	var currency Currency
	if err := json.NewDecoder(r.Body).Decode(&currency); err != nil {
		return Error(w, err)
	}

	ctx := r.Context()
	if err := h.flow.CreateCurrency(ctx, &currency); err != nil {
		return Error(w, err)
	}

	return Created(w, currency)
}

func (h *CurrencyHandler) CreateNewConversionRate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	var crate ConversionRate
	if err := json.NewDecoder(r.Body).Decode(&crate); err != nil {
		return Error(w, err)
	}

	ctx := r.Context()
	if err := h.flow.CreateConversionRate(ctx, &crate); err != nil {
		return Error(w, err)
	}

	return Created(w, crate)
}

func (h *CurrencyHandler) ConvertCurrency(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	var cparam ConversionParam
	if err := json.NewDecoder(r.Body).Decode(&cparam); err != nil {
		return Error(w, err)
	}

	ctx := r.Context()
	convertedAmount, err := h.flow.ConvertCurrency(ctx, cparam)
	if err != nil {
		return Error(w, err)
	}

	cparam.ConvertedAmount = convertedAmount

	return Created(w, cparam)
}
