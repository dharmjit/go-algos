package tree

import (
	"testing"
)

func Test_node_insert(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name    string
		n       *node
		args    args
		wantErr bool
	}{
		{"inserttest1", &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, args{5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.insert(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("node.insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_node_find(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		n    *node
		args args
		want bool
	}{
		{"findtest1", &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, args{3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.find(tt.args.v); got != tt.want {
				t.Errorf("node.find() = %v, want %v", got, tt.want)
			}
		})
	}
}
