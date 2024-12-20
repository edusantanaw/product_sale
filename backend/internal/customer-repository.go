package internal

type CustomerRepository struct {
	customers []Customer
}

func (c *CustomerRepository) Create(customer Customer) error {
	c.customers = append(c.customers, customer)
	return nil
}

func (c *CustomerRepository) Read() ([]Customer, error) {
	return c.customers, nil
}

func (c *CustomerRepository) Update(customer Customer) error {
	for i, v := range c.customers {
		if v.ID == customer.ID {
			c.customers[i] = customer
		}
	}
	return nil
}

func (c *CustomerRepository) Delete(id int) error {
	for i, v := range c.customers {
		if v.ID == id {
			c.customers = append(c.customers[:i], c.customers[i+1:]...)
		}
	}
	return nil
}
