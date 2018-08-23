package search

import "testing"

func Test_binarysearchwithoutrecursion(t *testing.T) {
	type args struct {
		array []int
		v     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5, 6, 7}, 4}, 3},
		{"test2", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarysearchwithoutrecursion(tt.args.array, tt.args.v); got != tt.want {
				t.Errorf("binarysearchwithoutrecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
