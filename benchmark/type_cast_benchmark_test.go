package benchmark

import (
	"context"
	"testing"
)

// goos: darwin
// goarch: amd64
// BenchmarkContextTypeCast
// BenchmarkContextTypeCast-8   	18313848	        65.3 ns/op
// BenchmarkCustomContext
// BenchmarkCustomContext-8     	1000000000	         0.656 ns/op
// PASS

const (
	ctxKey = "key"
	value  = "test"
)

func BenchmarkContextTypeCast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx := context.WithValue(context.Background(), ctxKey, value)
		printValue(ctx.Value(ctxKey).(string))
	}
}

type CustomContext struct {
	value string
}

func BenchmarkCustomContext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx := CustomContext{value: value}
		printValue(ctx.value)
	}
}

func printValue(value string) {
	// empty
}
