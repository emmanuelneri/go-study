package main

import (
	"math/rand"
	"time"
)

type Order struct {
	number   int
	date     time.Time
	products []Product
	total    float32
}

func NewOrder() *Order {
	return &Order{
		number:   rand.Intn(1000),
		date:     time.Now(),
		products: make([]Product, 0),
		total:    0,
	}
}

func (o *Order) addProduct(product Product, quantity int) {
	o.products = append(o.products, product)
	o.total = o.total + (product.Value * float32(quantity))
}

func (o Order) Number() int {
	return o.number
}

func (o Order) Date() time.Time {
	return o.date
}

func (o Order) Products() []Product {
	return o.products
}

func (o Order) Total() float32 {
	return o.total
}
