package internal

type Order struct {
	ID           int
	CustomerID   int
	ProductsCode []int
	Price        float64
	Status       string
	OrderDate    string
}
