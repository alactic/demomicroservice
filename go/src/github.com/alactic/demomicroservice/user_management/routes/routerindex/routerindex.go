package routerindex

import (
	"github.com/alactic/mygoproject/routes/auth"
	"github.com/alactic/mygoproject/routes/creditcards"
	"github.com/alactic/mygoproject/routes/customers"
	"github.com/alactic/mygoproject/routes/products"
	"github.com/alactic/mygoproject/routes/receipts"
	"github.com/alactic/mygoproject/routes/staff"
	"github.com/gorilla/mux"
)

func Routerindex(router *mux.Router) {
	customers.Customers(router)
	products.Products(router)
	creditcards.Creditcards(router)
	receipts.Receipts(router)
	auth.Auth(router)
	staff.Staff(router)
}
