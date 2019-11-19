package product

type Product struct {
	Id    string  `json:"id, omitempty"`
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
