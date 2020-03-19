package data

import (
	"encoding/json"
	"io"
	"time"
)

//Product represents a coffee
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

//Products represents list of product
type Products []*Product

//ToJSON return json encoded products
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

//FromJSON return json encoded products
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//GetProducts retrun list of all product
func GetProducts() Products {
	return productList
}

//AddProducts adds product into product list
func AddProducts(p *Product) {
	productList = append(productList, p)
}

//UpdateProduct takes id and product to upate product in product list
func UpdateProduct(id int, p *Product) {

	for i, prod := range productList {
		if prod.ID == id {
			productList[i] = p
		}
	}
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy Milky Coffee",
		Price:       3.14,
		SKU:         "1234",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee withput milk",
		Price:       1.99,
		SKU:         "fjd1233",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
}
