package stack

import (
	"testing"
)

func Test_pstack_push(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		s    *pstack
		args args
	}{
		{"pushtest1", &(pstack{1, 2, 3}), args{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.push(tt.args.v)
		})
	}
}

func Test_pstack_pop(t *testing.T) {
	tests := []struct {
		name string
		s    *pstack
		want int
	}{
		{"poptest1", &pstack{1, 2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.pop(); got != tt.want {
				t.Errorf("pstack.pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
