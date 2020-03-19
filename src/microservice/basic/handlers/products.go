package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/kunalbanjara/repo/data"
)

//Products struct represents products handler
type Products struct {
	l *log.Logger
}

//NewProducts creates Product struct with log.logger injected
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

//ServeHTTP retruns hello handler
func (h *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		h.addProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		regEx := regexp.MustCompile(`/([0-9]+)`)
		g := regEx.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid Url1", http.StatusBadRequest)
			return
		}

		idStr := g[0][1]
		id, _ := strconv.Atoi(idStr)
		h.updateProduct(id, rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (h *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error while marshalling the product list", http.StatusInternalServerError)
	}
}

func (h *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	data.AddProducts(prod)
}

func (h *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	data.UpdateProduct(id, prod)
}
