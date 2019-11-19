package receipt

import (
	"github.com/alactic/mygoproject/models/customers"
	"github.com/alactic/mygoproject/models/product"
)

type Receipt struct {
	Id       string             `json:"id, omitempty"`
	Type     string             `json:"type"`
	Customer customers.Customer `json:"customer"`
	Product  []product.Product  `json:"products"`
	Total    float32            `json:"total`
}
