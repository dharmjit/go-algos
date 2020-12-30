package unionfind

import "testing"

func TestUnionFind_Find(t *testing.T) {
	uf := Make([]int{1, 2, 3})
	type fields struct {
		Set []Node
	}
	type args struct {
		node int
	}
	tests := []struct {
		name string
		uf   *UnionFind
		args args
		want int
	}{
		{
			name: "SUCCESS-1",
			args: args{node: 3},
			uf:   uf,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := tt.uf
			if got := uf.Find(tt.args.node); got != tt.want {
				t.Errorf("UnionFind.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
