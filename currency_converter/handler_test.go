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
