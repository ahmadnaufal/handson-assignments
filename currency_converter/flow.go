package currencyconverter

import "context"

type CurrencyFlow interface {
	GetCurrencies(ctx context.Context) ([]Currency, error)
	CreateCurrency(ctx context.Context, currency *Currency) error
	CreateConversionRate(ctx context.Context, crate *ConversionRate) error
	ConvertCurrency(ctx context.Context, cparam ConversionParam) (float64, error)
}

type CurrencyProvider struct{}

func NewCurrencyFlow() CurrencyProvider {
	return CurrencyProvider{}
}

func (p *CurrencyProvider) GetCurrencies(ctx context.Context) ([]Currency, error) {
	currencies := []Currency{
		{
			ID:   1,
			Name: "JPY",
		},
	}

	return currencies, nil
}

func (p *CurrencyProvider) CreateCurrency(ctx context.Context, currency *Currency) error {
	return nil
}

func (p *CurrencyProvider) CreateConversionRate(ctx context.Context, crate *ConversionRate) error {
	return nil
}

func (p *CurrencyProvider) ConvertCurrency(ctx context.Context, cparam ConversionParam) (float64, error) {
	return 0.0, nil
}
