package list

import (
	"testing"
)

func Test_list_pushFront(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		l       *list
		args    args
		wantErr bool
	}{
		{"pushFrontTest1", &list{}, args{1}, false},
		{"pushFrontTest2", &list{&element{1, &element{2, nil}}, 2}, args{3}, false},
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
		{"popFrontTest2", &list{&element{1, &element{2, nil}}, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.popFront(); got != tt.want {
				t.Errorf("list.popFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_pushBack(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name    string
		l       *list
		args    args
		wantErr bool
	}{
		{"pushBackTest1", &list{&element{1, &element{2, nil}}, 2}, args{3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.pushBack(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("list.pushBack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_list_popBack(t *testing.T) {
	tests := []struct {
		name string
		l    *list
		want int
	}{
		{"popBackTest2", &list{&element{1, &element{2, nil}}, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.popBack(); got != tt.want {
				t.Errorf("list.popBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_len(t *testing.T) {
	tests := []struct {
		name string
		l    *list
		want int
	}{
		{"lentest1", &list{&element{1, &element{2, nil}}, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.len(); got != tt.want {
				t.Errorf("list.len() = %v, want %v", got, tt.want)
			}
		})
	}
}
