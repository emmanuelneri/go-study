package repository

import (
	"app/pkg/domain"
	"database/sql"
)

const insertQuery = "INSERT INTO sales_order(customer, total) VALUES($1,$2) RETURNING id"

type OrderRepository interface {
	Save(order *domain.Order) error
}

type OrderRepositoryImpl struct {
	db *sql.DB
}

func NewOrderRepositoryImpl(db *sql.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) Save(order *domain.Order) error {
	var id int
	err := r.db.QueryRow(insertQuery, order.Customer, order.Total).Scan(&id)
	order.ID = id
	return err
}
