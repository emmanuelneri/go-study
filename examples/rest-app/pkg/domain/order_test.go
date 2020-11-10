package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrder_Validate(t *testing.T) {
	t.Run("Valid order", func(t *testing.T) {
		order := Order{
			ID:       10,
			Customer: "Customer 1",
			Total:    100,
		}
		assert.Nil(t, order.CreationValidate())
	})

	t.Run("should return error with zero ID", func(t *testing.T) {
		order := Order{
			ID:       0,
			Customer: "Customer 1",
			Total:    100,
		}
		assert.Equal(t, ErrInvalidID, order.CreationValidate())
	})

	t.Run("should return error with negative ID", func(t *testing.T) {
		order := Order{
			ID:       -10,
			Customer: "Customer 1",
			Total:    100,
		}
		assert.Equal(t, ErrInvalidID, order.CreationValidate())
	})

	t.Run("should return error no customer attribute", func(t *testing.T) {
		order := Order{
			ID:    1,
			Total: 100,
		}
		assert.Equal(t, ErrCustomerRequired, order.CreationValidate())
	})

	t.Run("should return error with negative value", func(t *testing.T) {
		order := Order{
			ID:       1,
			Customer: "Customer 1",
			Total:    -1,
		}
		assert.Equal(t, ErrInvalidValue, order.CreationValidate())
	})

	t.Run("should not error with no value", func(t *testing.T) {
		order := Order{
			ID:       1,
			Customer: "Customer 1",
		}
		assert.Nil(t, order.CreationValidate())
	})
}
