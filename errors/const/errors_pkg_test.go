package _const

import (
	"github.com/pkg/errors"
	"testing"
)

func Test_error_wrap(t *testing.T) {
	type args struct {
		error error
	}
	tests := []struct {
		name    string
		args    args
		compare func(args args) bool
		want    bool
	}{
		{name: "compare with errors.Is a error created with errors.Wrap and the result should be true",
			args: args{error: errors.Wrap(InvalidValue, "Wrap error")},
			compare: func(args args) bool {
				return errors.Is(args.error, InvalidValue)
			}, want: true},

		{name: "compare cause error with errors.Is and the result should be true",
			args: args{error: errors.Cause(InvalidValue)},
			compare: func(args args) bool {
				return errors.Is(args.error, InvalidValue)
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
