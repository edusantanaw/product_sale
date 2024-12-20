package internal

type ProductRepository struct {
	product []Product
}

func (r *ProductRepository) Create(product Product) error {
	r.product = append(r.product, product)
	return nil
}

func (r *ProductRepository) Read() ([]Product, error) {
	return r.product, nil
}

func (r *ProductRepository) Update(product Product) error {
	for i, v := range r.product {
		if v.Code == product.Code {
			r.product[i] = product
		}
	}
	return nil
}

func (r *ProductRepository) Delete(code int) error {
	for i, v := range r.product {
		if v.Code == code {
			r.product = append(r.product[:i], r.product[i+1:]...)
		}
	}
	return nil
}
