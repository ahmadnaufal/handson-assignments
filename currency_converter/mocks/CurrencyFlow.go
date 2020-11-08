// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	context "context"

	currencyconverter "github.com/ahmadnaufal/handson-assignments/currency_converter"
	mock "github.com/stretchr/testify/mock"
)

// CurrencyFlow is an autogenerated mock type for the CurrencyFlow type
type CurrencyFlow struct {
	mock.Mock
}

// ConvertCurrency provides a mock function with given fields: ctx, cparam
func (_m *CurrencyFlow) ConvertCurrency(ctx context.Context, cparam currencyconverter.ConversionParam) (float64, error) {
	ret := _m.Called(ctx, cparam)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context, currencyconverter.ConversionParam) float64); ok {
		r0 = rf(ctx, cparam)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, currencyconverter.ConversionParam) error); ok {
		r1 = rf(ctx, cparam)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateConversionRate provides a mock function with given fields: ctx, crate
func (_m *CurrencyFlow) CreateConversionRate(ctx context.Context, crate *currencyconverter.ConversionRate) error {
	ret := _m.Called(ctx, crate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currencyconverter.ConversionRate) error); ok {
		r0 = rf(ctx, crate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCurrency provides a mock function with given fields: ctx, currency
func (_m *CurrencyFlow) CreateCurrency(ctx context.Context, currency *currencyconverter.Currency) error {
	ret := _m.Called(ctx, currency)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *currencyconverter.Currency) error); ok {
		r0 = rf(ctx, currency)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCurrencies provides a mock function with given fields: ctx
func (_m *CurrencyFlow) GetCurrencies(ctx context.Context) ([]currencyconverter.Currency, error) {
	ret := _m.Called(ctx)

	var r0 []currencyconverter.Currency
	if rf, ok := ret.Get(0).(func(context.Context) []currencyconverter.Currency); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]currencyconverter.Currency)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
