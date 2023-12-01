package manylogic

import (
	"fmt"
	"testing"
)

func TestOr(t *testing.T) {
	type args struct {
		bs []bool
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{[]bool{}}, false},

		{args{[]bool{false}}, false},
		{args{[]bool{true}}, true},

		{args{[]bool{false, false}}, false},
		{args{[]bool{false, true}}, true},
		{args{[]bool{true, false}}, true},
		{args{[]bool{true, true}}, true},

		{args{[]bool{false, false, false}}, false},
		{args{[]bool{false, false, true}}, true},
		{args{[]bool{false, true, false}}, true},
		{args{[]bool{false, true, true}}, true},
		{args{[]bool{true, false, false}}, true},
		{args{[]bool{true, false, true}}, true},
		{args{[]bool{true, true, false}}, true},
		{args{[]bool{true, true, true}}, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := Or(tt.args.bs...); got != tt.want {
				t.Errorf("Or(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func TestOrCb(t *testing.T) {
	intToBoolCallback := func(i int) bool { return i == 1 }

	type args struct {
		bs []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{[]int{}}, false},

		{args{[]int{0}}, false},
		{args{[]int{1}}, true},

		{args{[]int{0, 0}}, false},
		{args{[]int{0, 1}}, true},
		{args{[]int{1, 0}}, true},
		{args{[]int{1, 1}}, true},

		{args{[]int{0, 0, 0}}, false},
		{args{[]int{0, 0, 1}}, true},
		{args{[]int{0, 1, 0}}, true},
		{args{[]int{0, 1, 1}}, true},
		{args{[]int{1, 0, 0}}, true},
		{args{[]int{1, 0, 1}}, true},
		{args{[]int{1, 1, 0}}, true},
		{args{[]int{1, 1, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := OrCb(tt.args.bs, intToBoolCallback); got != tt.want {
				t.Errorf("OrCb(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
