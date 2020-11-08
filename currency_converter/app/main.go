package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	currencyconverter "github.com/ahmadnaufal/handson-assignments/currency_converter"

	"github.com/joeshaw/envdecode"
	"github.com/julienschmidt/httprouter"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

type Config struct {
	Port string `env:"APP_PORT,required"`
}

func main() {
	var cfg Config
	err := envdecode.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	currencyFlow := currencyconverter.NewCurrencyFlow()
	currencyHandler := currencyconverter.NewCurrencyHandler(&currencyFlow)

	router := httprouter.New()
	router.GET("/currencies", decorate(currencyHandler.GetCurrencies))
	router.POST("/currencies", decorate(currencyHandler.CreateCurrency))
	router.POST("/currencies/rates", decorate(currencyHandler.CreateNewConversionRate))
	router.POST("/convert", decorate(currencyHandler.ConvertCurrency))

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Server available at %s\n", s.Addr)
	if serr := s.ListenAndServe(); serr != http.ErrServerClosed {
		log.Fatal(serr)
	}
}

func decorate(handler currencyconverter.Handler) httprouter.Handle {
	return currencyconverter.HTTP(handler)
}
