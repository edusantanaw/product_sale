package internal

type OrderRepository struct {
	orders []Order
}

func (o *OrderRepository) Create(order Order) error {
	o.orders = append(o.orders, order)
	return nil
}

func (o *OrderRepository) Read() ([]Order, error) {
	return o.orders, nil
}

func (o *OrderRepository) Update(order Order) error {
	for i, v := range o.orders {
		if v.ID == order.ID {
			o.orders[i] = order
		}
	}
	return nil
}

func (o *OrderRepository) Delete(id int) error {
	for i, v := range o.orders {
		if v.ID == id {
			o.orders = append(o.orders[:i], o.orders[i+1:]...)
		}
	}
	return nil
}
