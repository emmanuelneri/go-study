package domain

type Order struct {
	ID       int
	Customer string
	Total    float32
}

func (o Order) CreationValidate() error {
	if o.ID <= 0 {
		return ErrInvalidID
	}

	if o.Customer == "" {
		return ErrCustomerRequired
	}

	if o.Total < 0 {
		return ErrInvalidValue
	}

	return nil
}
