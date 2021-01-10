package _const

import (
	"errors"
	"fmt"
	"testing"
)

func Test_error(t *testing.T) {
	type args struct {
		error error
	}
	tests := []struct {
		name    string
		args    args
		compare func(args args) bool
		want    bool
	}{
		{name: "compare error with == should be false",
			args: args{error: InvalidValue},
			compare: func(args args) bool {
				return args.error == errors.New("Invalid Param ")
			}, want: false},

		{name: "compare error message with == should be true",
			args: args{error: InvalidValue},
			compare: func(args args) bool {
				return args.error.Error() == errors.New("Invalid Param ").Error()
			}, want: true},

		{name: "compare error with errors.Is should be false because Is unwrap first argument",
			args: args{error: InvalidValue},
			compare: func(args args) bool {
				return errors.Is(args.error, errors.New("Invalid Param "))
			}, want: false},

		{name: "compare composite error with errors.Is should be true",
			args: args{error: fmt.Errorf("adding more context: %w", InvalidValue)},
			compare: func(args args) bool {
				return errors.Is(args.error, InvalidValue)
			}, want: true},

		{name: "compare error type with errors.As should be true because is same type",
			args: args{error: InvalidValue},
			compare: func(args args) bool {
				return errors.As(args.error, &InvalidValue)
			}, want: true},

		{name: "compare error type with errors.As should be false because isn`t same type",
			args: args{error: InvalidValue},
			compare: func(args args) bool {
				customError := CustomError{error: "Invalid Param "}
				return errors.As(args.error, &customError)
			}, want: false},

		{name: "compare error type type should be true because is same type",
			args: args{error: InvalidValue},
			compare: func(args args) bool {
				_, ok := args.error.(CustomError)
				return ok
			}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.compare(tt.args) != tt.want {
				t.Errorf("compare() error = %v, want %v", tt.args.error, tt.want)
			}
		})
	}
}
