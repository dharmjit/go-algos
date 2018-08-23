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

func Test_node_delete(t *testing.T) {
	type args struct {
		v      int
		parent *node
	}
	tests := []struct {
		name    string
		n       *node
		args    args
		wantErr bool
	}{
		{"deletetest1", &node{2, &node{1, nil, nil}, &node{3, nil, nil}}, args{3, &node{2, &node{1, nil, nil}, &node{3, nil, nil}}}, false},
		{"deletetest1", &node{5, &node{3,
			&node{2, &node{1, nil, nil}, &node{9, nil, nil}}, &node{4, nil, nil}},
			&node{7, &node{6, nil, nil}, &node{8, nil, nil}}}, args{2, &node{3,
			&node{2, &node{1, nil, nil}, &node{9, nil, nil}}, &node{4, nil, nil}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.delete(tt.args.v, tt.args.parent); (err != nil) != tt.wantErr {
				t.Errorf("node.delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
