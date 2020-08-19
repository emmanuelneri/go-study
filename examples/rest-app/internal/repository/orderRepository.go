package repository

import (
	"app/config"
	"app/pkg/model"
	"context"
	"database/sql"
)

const insertQuery = "INSERT INTO sales_order(customer, total) VALUES($1,$2) RETURNING id"

type OrderRepository interface {
	Save(order *model.Order) error
}

type OrderRepositoryImpl struct {
	db *sql.DB
}

func NewOrderRepositoryImpl(ctx context.Context) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db: ctx.Value(config.DB).(*sql.DB)}
}

func (r *OrderRepositoryImpl) Save(order *model.Order) error {
	var id int
	err := r.db.QueryRow(insertQuery, order.Customer, order.Total).Scan(&id)
	order.ID = id
	return err
}
