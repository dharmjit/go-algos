package search

import "testing"

func Test_binarysearchwithrecursion(t *testing.T) {
	type args struct {
		array  []int
		low    int
		high   int
		target int
	}
	tests := []struct {
		name      string
		args      args
		wantIndex int
	}{
		{"binarysearchtest1", args{[]int{1, 2, 3, 4, 5}, 0, 4, 3}, 2},
		{"binarysearchtest2", args{[]int{1, 2, 3, 4, 5}, 0, 4, 6}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := binarysearchwithrecursion(tt.args.array, tt.args.low, tt.args.high, tt.args.target); gotIndex != tt.wantIndex {
				t.Errorf("binarysearchwithrecursion() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}
