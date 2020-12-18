package main

import (
	"testing"
)

func TestSliceInitialization(t *testing.T) {

	t.Run("initialize with empty slice", func(t *testing.T) {
		var numbers [8]int
		for i := 0; i < 8; i++ {
			numbers[i] = i
		}

		for i, value := range numbers {
			if value != numbers[i] {
				t.Errorf("expected :%d - actual %d", value, numbers[i])
			}
		}
	})

	t.Run("initialize with empty slice", func(t *testing.T) {
		var numbers []int
		for i := 0; i < 8; i++ {
			numbers = append(numbers, i)
		}

		if len(numbers) != 8 {
			t.Errorf("expected :%d - actual %d", 8, len(numbers))
		}

		for i, value := range numbers {
			if value != numbers[i] {
				t.Errorf("expected :%d - actual %d", value, numbers[i])
			}
		}
	})

	t.Run("initialize with empty array", func(t *testing.T) {
		numbers := []int{}
		for i := 0; i < 8; i++ {
			numbers = append(numbers, i)
		}

		if len(numbers) != 8 {
			t.Errorf("expected :%d - actual %d", 8, len(numbers))
		}

		for i, value := range numbers {
			if value != numbers[i] {
				t.Errorf("expected :%d - actual %d", value, numbers[i])
			}
		}
	})

	t.Run("initialize with make", func(t *testing.T) {
		numbers := make([]int, 0)
		for i := 0; i < 8; i++ {
			numbers = append(numbers, i)
		}

		if len(numbers) != 8 {
			t.Errorf("expected :%d - actual %d", 8, len(numbers))
		}

		for i, value := range numbers {
			if value != numbers[i] {
				t.Errorf("expected :%d - actual %d", value, numbers[i])
			}
		}
	})

	t.Run("initialize with make and size defined", func(t *testing.T) {
		numbers := make([]int, 0, 8)
		for i := 0; i < 8; i++ {
			numbers = append(numbers, i)
		}

		if len(numbers) != 8 {
			t.Errorf("expected :%d - actual %d", 8, len(numbers))
		}

		if cap(numbers) != 8 {
			t.Errorf("expected :%d - actual %d", 8, cap(numbers))
		}

		for i, value := range numbers {
			if value != numbers[i] {
				t.Errorf("expected :%d - actual %d", value, numbers[i])
			}
		}
	})
}
