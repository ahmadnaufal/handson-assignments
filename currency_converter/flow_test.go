package currencyconverter_test

import (
	"context"
	"testing"

	currencyconverter "github.com/ahmadnaufal/handson-assignments/currency_converter"
	"github.com/stretchr/testify/assert"
)

func TestFlowGetCurrencies(t *testing.T) {
	testcases := []struct {
		title             string
		shouldReturnError bool
	}{
		{
			title:             "success",
			shouldReturnError: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			flow := currencyconverter.NewCurrencyFlow()

			currencies, err := flow.GetCurrencies(context.Background())

			if tc.shouldReturnError {
				assert.Error(t, err)
				assert.Nil(t, currencies)
			} else {
				assert.NoError(t, err)
				assert.True(t, len(currencies) > 0)
			}
		})
	}
}

func TestFlowCreateCurrency(t *testing.T) {
	testcases := []struct {
		title             string
		input             currencyconverter.Currency
		shouldReturnError bool
	}{
		{
			title: "success",
			input: currencyconverter.Currency{
				ID:   1,
				Name: "JPY",
			},
			shouldReturnError: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			flow := currencyconverter.NewCurrencyFlow()

			err := flow.CreateCurrency(context.Background(), &tc.input)

			if tc.shouldReturnError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func TestFlowCreateConversionRate(t *testing.T) {
	testcases := []struct {
		title             string
		input             currencyconverter.ConversionRate
		shouldReturnError bool
	}{
		{
			title: "success",
			input: currencyconverter.ConversionRate{
				CurrencyIDFrom: 1,
				CurrencyIDTo:   2,
				Rate:           300.0,
			},
			shouldReturnError: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			flow := currencyconverter.NewCurrencyFlow()

			err := flow.CreateConversionRate(context.Background(), &tc.input)

			if tc.shouldReturnError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFlowConvertCurrency(t *testing.T) {
	testcases := []struct {
		title             string
		input             currencyconverter.ConversionParam
		expectedResult    float64
		shouldReturnError bool
	}{
		{
			title: "success",
			input: currencyconverter.ConversionParam{
				CurrencyIDFrom: 1,
				CurrencyIDTo:   2,
				BaseAmount:     300.0,
			},
			expectedResult:    0.0,
			shouldReturnError: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			flow := currencyconverter.NewCurrencyFlow()

			convertedAmount, err := flow.ConvertCurrency(context.Background(), tc.input)

			if tc.shouldReturnError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, convertedAmount)
			}
		})
	}
}
