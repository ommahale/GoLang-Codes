package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ommahale/ecomapi/data"
	"github.com/ommahale/ecomapi/utils"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l: l}
}

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lp := data.GetInventory()
	d, err := json.Marshal(lp)
	utils.HandleHttpError(err, "Unable to parse JSON", &w)
	w.Write(d)

}
