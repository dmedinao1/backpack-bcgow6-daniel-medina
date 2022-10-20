package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubtract(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"One minus two equal to minus one", args{1, 2}, -1},
		{"Zero minus two equal to minus two", args{0, 2}, -2},
		{"Ten minus nine equal to one", args{10, 9}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Subtract(tt.args.a, tt.args.b))
		})
	}
}
