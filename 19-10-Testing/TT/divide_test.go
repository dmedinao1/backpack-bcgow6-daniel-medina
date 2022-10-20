package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivide(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr assert.ErrorAssertionFunc
	}{
		{"should get 2.5 when dividing 5 and 2", args{5, 2}, 2.5, assert.NoError},
		{"should get error when dividing 5 and 0", args{5, 0}, 0, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.args.a, tt.args.b)
			if !tt.wantErr(t, err, fmt.Sprintf("Divide(%v, %v)", tt.args.a, tt.args.b)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Divide(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
