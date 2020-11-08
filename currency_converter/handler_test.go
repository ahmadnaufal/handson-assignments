package currencyconverter_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	currencyconverter "github.com/ahmadnaufal/handson-assignments/currency_converter"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrencies(t *testing.T) {
	testcases := []struct {
		title          string
		expectedStatus int
	}{
		{
			title:          "success",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			handler := currencyconverter.NewCurrencyHandler()
			router := httprouter.New()
			router.GET("/currencies", currencyconverter.HTTP(handler.GetCurrencies))

			req, _ := http.NewRequest("GET", "/currencies", nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)
		})
	}
}

func TestCreateCurrency(t *testing.T) {
	testcases := []struct {
		title          string
		body           string
		expectedStatus int
	}{
		{
			title: "success",
			body: `{
        "id": 1,
        "name": "JPY"
      }`,
			expectedStatus: http.StatusCreated,
		},
		{
			title: "failed",
			body: `{
        "id": 1,
        "name": 123
      }`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			handler := currencyconverter.NewCurrencyHandler()
			router := httprouter.New()
			router.POST("/currencies", currencyconverter.HTTP(handler.CreateCurrency))

			req, _ := http.NewRequest("POST", "/currencies", strings.NewReader(tc.body))
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)
		})
	}
}

func TestCreateNewConversionRate(t *testing.T) {
	testcases := []struct {
		title          string
		body           string
		expectedStatus int
	}{
		{
			title: "success",
			body: `{
        "currency_id_from": 1,
				"currency_id_to": 2,
				"rate": 250
      }`,
			expectedStatus: http.StatusCreated,
		},
		{
			title: "failed",
			body: `{
        "currency_id_from": 1,
				"currency_id_to": "invalid",
				"rate": 250
      }`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			handler := currencyconverter.NewCurrencyHandler()
			router := httprouter.New()
			router.POST("/currencies/rates", currencyconverter.HTTP(handler.CreateNewConversionRate))

			req, _ := http.NewRequest("POST", "/currencies/rates", strings.NewReader(tc.body))
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)
		})
	}
}

func TestConvertCurrency(t *testing.T) {
	testcases := []struct {
		title          string
		body           string
		expectedStatus int
	}{
		{
			title: "success",
			body: `{
        "currency_id_from": 1,
				"currency_id_to": 2,
				"base_amount": 250
      }`,
			expectedStatus: http.StatusCreated,
		},
		{
			title: "failed",
			body: `{
        "currency_id_from": 1,
				"currency_id_to": "invalid",
				"base_amount": 250
      }`,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			handler := currencyconverter.NewCurrencyHandler()
			router := httprouter.New()
			router.POST("/convert", currencyconverter.HTTP(handler.ConvertCurrency))

			req, _ := http.NewRequest("POST", "/convert", strings.NewReader(tc.body))
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)
		})
	}
}
