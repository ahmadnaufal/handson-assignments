package currencyconverter

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler func(http.ResponseWriter, *http.Request, httprouter.Params) error

func HTTP(handle Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := handle(w, r, p)

		if err != nil {
			log.Printf("Error: %s", err.Error())
		} else {
			log.Println("OK")
		}
	}
}
