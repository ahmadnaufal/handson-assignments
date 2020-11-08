package currencyconverter

import "time"

type Currency struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ConversionRate struct {
	CurrencyIDFrom int64   `json:"currency_id_from"`
	CurrencyIDTo   int64   `json:"currency_id_to"`
	Rate           float64 `json:"rate"`
}

type ConversionParam struct {
	CurrencyIDFrom  int64   `json:"currency_id_from"`
	CurrencyIDTo    int64   `json:"currency_id_to"`
	BaseAmount      float64 `json:"base_amount"`
	ConvertedAmount float64 `json:"converted_amount"`
}
