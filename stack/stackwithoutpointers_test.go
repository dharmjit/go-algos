package stack

import (
	"reflect"
	"testing"
)

func Test_stack_push(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		s    stack
		args args
		want stack
	}{
		// TODO: Add test cases.
		{"pushtest1", stack{1, 2, 3}, args{4}, stack{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.push(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stack.push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_pop(t *testing.T) {
	tests := []struct {
		name  string
		s     stack
		want  stack
		want1 int
	}{
		{"poptest1", stack{1, 2, 3}, stack{1, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.pop()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stack.pop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("stack.pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
