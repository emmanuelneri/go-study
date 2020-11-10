package container

import (
	"app/internal/repository"
	"app/internal/service"
	"database/sql"
)

type Container struct {
	OrderService service.OrderService
}

func Start(db *sql.DB) *Container {
	orderRepository := repository.NewOrderRepositoryImpl(db)
	orderService := service.OrderService(orderRepository)

	return &Container{OrderService: orderService}
}
