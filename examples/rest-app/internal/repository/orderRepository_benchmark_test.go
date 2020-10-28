package repository

import (
	"app/config"
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func newRepositoryFromContext(ctx context.Context) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db: ctx.Value(config.DB).(*sql.DB)}
}

func newRepository(db *sql.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db: db}
}

func BenchmarkRepositoryFromContext(b *testing.B) {
	db, _, _ := sqlmock.New()
	for i := 0; i < b.N; i++ {
		ctx := context.WithValue(context.Background(), config.DB, db)
		newRepositoryFromContext(ctx)
	}
}

func BenchmarkRepositoryFromDBParameter(b *testing.B) {
	db, _, _ := sqlmock.New()
	for i := 0; i < b.N; i++ {
		newRepository(db)
	}
}
