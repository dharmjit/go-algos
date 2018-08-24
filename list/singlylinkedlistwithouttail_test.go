package list

import (
	"testing"
)

func Test_list_addFront(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		l       *list
		args    args
		wantErr bool
	}{
		{"addFrontTest1", &list{}, args{1}, false},
		{"addFrontTest2", &list{&element{1, &element{2, nil}}}, args{3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.pushFront(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("list.addFront() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_list_popFront(t *testing.T) {
	tests := []struct {
		name string
		l    *list
		want int
	}{
		{"addFrontTest2", &list{&element{1, &element{2, nil}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.popFront(); got != tt.want {
				t.Errorf("list.popFront() = %v, want %v", got, tt.want)
			}
		})
	}
}
