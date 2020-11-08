package currencyconverter_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	currencyconverter "github.com/ahmadnaufal/handson-assignments/currency_converter"
	"github.com/ahmadnaufal/handson-assignments/currency_converter/mocks"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCurrencies(t *testing.T) {
	testcases := []struct {
		title          string
		initMock       func(*mocks.CurrencyFlow)
		expectedStatus int
	}{
		{
			title: "success",
			initMock: func(m *mocks.CurrencyFlow) {
				currencies := []currencyconverter.Currency{{ID: 1, Name: "JPY"}}
				m.On("GetCurrencies", mock.Anything).Return(currencies, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			title: "error get currencies",
			initMock: func(m *mocks.CurrencyFlow) {
				m.On("GetCurrencies", mock.Anything).Return(nil, errors.New("error get currencies"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			flow := mocks.CurrencyFlow{}
			if tc.initMock != nil {
				tc.initMock(&flow)
			}

			handler := currencyconverter.NewCurrencyHandler(&flow)
			router := httprouter.New()
			router.GET("/currencies", currencyconverter.HTTP(handler.GetCurrencies))

			req, _ := http.NewRequest("GET", "/currencies", nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)

			flow.AssertExpectations(t)
		})
	}
}

func TestCreateCurrency(t *testing.T) {
	testcases := []struct {
		title          string
		body           string
		initMock       func(*mocks.CurrencyFlow)
		expectedStatus int
	}{
		{
			title: "success",
			body: `{
        "id": 1,
        "name": "JPY"
      }`,
			initMock: func(m *mocks.CurrencyFlow) {
				m.On("CreateCurrency", mock.Anything, mock.MatchedBy(func(cp *currencyconverter.Currency) bool {
					return cp.ID == int64(1) && cp.Name == "JPY"
				})).Return(nil)
			},
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
		{
			title: "error create currency",
			body: `{
        "id": 1,
        "name": "JPY"
      }`,
			initMock: func(m *mocks.CurrencyFlow) {
				m.On("CreateCurrency", mock.Anything, mock.Anything).Return(errors.New("error create currency"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			flow := mocks.CurrencyFlow{}
			if tc.initMock != nil {
				tc.initMock(&flow)
			}

			handler := currencyconverter.NewCurrencyHandler(&flow)
			router := httprouter.New()
			router.POST("/currencies", currencyconverter.HTTP(handler.CreateCurrency))

			req, _ := http.NewRequest("POST", "/currencies", strings.NewReader(tc.body))
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)

			flow.AssertExpectations(t)
		})
	}
}

func TestCreateNewConversionRate(t *testing.T) {
	testcases := []struct {
		title          string
		body           string
		initMock       func(*mocks.CurrencyFlow)
		expectedStatus int
	}{
		{
			title: "success",
			body: `{
        "currency_id_from": 1,
				"currency_id_to": 2,
				"rate": 250
      }`,
			initMock: func(m *mocks.CurrencyFlow) {
				m.On("CreateConversionRate", mock.Anything, mock.MatchedBy(func(cp *currencyconverter.ConversionRate) bool {
					return cp.CurrencyIDFrom == int64(1) && cp.CurrencyIDTo == int64(2) && cp.Rate == float64(250)
				})).Return(nil)
			},
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
		{
			title: "error create conversion rate",
			body: `{
        "currency_id_from": 1,
				"currency_id_to": 2,
				"rate": 250
      }`,
			initMock: func(m *mocks.CurrencyFlow) {
				m.On("CreateConversionRate", mock.Anything, mock.Anything).Return(errors.New("error create conversion rate"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			flow := mocks.CurrencyFlow{}
			if tc.initMock != nil {
				tc.initMock(&flow)
			}

			handler := currencyconverter.NewCurrencyHandler(&flow)
			router := httprouter.New()
			router.POST("/currencies/rates", currencyconverter.HTTP(handler.CreateNewConversionRate))

			req, _ := http.NewRequest("POST", "/currencies/rates", strings.NewReader(tc.body))
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)

			flow.AssertExpectations(t)
		})
	}
}

func TestConvertCurrency(t *testing.T) {
	testcases := []struct {
		title          string
		body           string
		initMock       func(*mocks.CurrencyFlow)
		expectedStatus int
	}{
		{
			title: "success",
			body: `{
        "currency_id_from": 1,
				"currency_id_to": 2,
				"base_amount": 250
			}`,
			initMock: func(m *mocks.CurrencyFlow) {
				m.On("ConvertCurrency", mock.Anything, mock.MatchedBy(func(cp currencyconverter.ConversionParam) bool {
					return cp.CurrencyIDFrom == int64(1) && cp.CurrencyIDTo == int64(2) && cp.BaseAmount == float64(250)
				})).Return(300.0, nil)
			},
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
		{
			title: "convert currency error",
			body: `{
        "currency_id_from": 1,
				"currency_id_to": 2,
				"base_amount": 250
			}`,
			initMock: func(m *mocks.CurrencyFlow) {
				m.On("ConvertCurrency", mock.Anything, mock.Anything).Return(0.0, errors.New("convert currency error"))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			flow := mocks.CurrencyFlow{}
			if tc.initMock != nil {
				tc.initMock(&flow)
			}

			handler := currencyconverter.NewCurrencyHandler(&flow)
			router := httprouter.New()
			router.POST("/convert", currencyconverter.HTTP(handler.ConvertCurrency))

			req, _ := http.NewRequest("POST", "/convert", strings.NewReader(tc.body))
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatus, rr.Code)

			flow.AssertExpectations(t)
		})
	}
}
