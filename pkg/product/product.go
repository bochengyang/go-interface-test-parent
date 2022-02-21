package product

import "fmt"

type Information interface {
	General()
	Attributes()
	Inventory()
}

type Product struct {
	Name, Description string
	Weight, Price     float64
	Stock             int
}

func NewProduct() Information {
	return &Product{}
}

func (prd Product) General() {
	fmt.Printf("\n%s", prd.Name)
	fmt.Printf("\n%s\n", prd.Description)
	fmt.Println(prd.Price)
}

func (prd Product) Attributes() {
	fmt.Println(prd.Weight)
}

func (prd Product) Inventory() {
	fmt.Println(prd.Stock)
}
