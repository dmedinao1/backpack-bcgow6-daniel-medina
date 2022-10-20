package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Should get [1,2,3,4] when sorting [4,1,2,3]", args{[]int{4, 1, 2, 3}}, []int{1, 2, 3, 4}},
		{"Should get [1,2,3,4] when sorting [1,2,3,4]", args{[]int{4, 1, 2, 3}}, []int{1, 2, 3, 4}},
		{"Should get [] when sorting []", args{[]int{}}, []int{}},
		{"Should get nil when sorting nil", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Sort(tt.args.items), "Sort(%v)", tt.args.items)
		})
	}
}
