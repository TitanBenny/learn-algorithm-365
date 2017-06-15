package dsandaa

import (
	"log"
	"reflect"
	"testing"
)

func Test_getrandint(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		// TODO: Add test cases.
		{
			args: 0,
			want: 0,
		},
		{
			args: 5,
			want: 5,
		},
		{
			args: 10,
			want: 10,
		},
		{
			args: 100,
			want: 100,
		},
		{
			args: 1000,
			want: 1000,
		},
	}
	for _, tt := range tests {
		if got := getrandint(tt.args); !reflect.DeepEqual(len(got), tt.want) {
			t.Errorf("%q. getrandint() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_bubbleSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "冒泡排序",
			args: []int{4, 5, 6, 1, 8, 0, 2, 7},
			want: []int{8, 7, 6, 5, 4, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		if got := bubbleSort(tt.args); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. bubbleSort() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_getK(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 float64
	}{
		// TODO: Add test cases.
		{
			name: "10个数组",
			args: args{
				n: 10,
				k: 5,
			},
			want:  5,
			want1: 0.23,
		},
		{
			name: "100个数组",
			args: args{
				n: 100,
				k: 50,
			},
			want:  50,
			want1: 0.23,
		},
		{
			name: "1000个数组",
			args: args{
				n: 1000,
				k: 500,
			},
			want:  500,
			want1: 0.23,
		},
		{
			name: "10000个数组",
			args: args{
				n: 10000,
				k: 5000,
			},
			want:  5000,
			want1: 0.23,
		},
		{
			name: "10000个数组",
			args: args{
				n: 100000,
				k: 50000,
			},
			want:  50000,
			want1: 0.23,
		},
	}
	for _, tt := range tests {
		got, got1 := getK(tt.args.n, tt.args.k)
		if got != tt.want {
			t.Errorf("%q. getK() got = %v, want %v", tt.name, got, tt.want)
		}
		// if got1 != tt.want1 {
		// 	t.Errorf("%q. getK() got1 = %v, want %v", tt.name, got1, tt.want1)
		// }
		log.Printf("%s,第%d,结果%d,耗时%v", tt.name, tt.args.k, got, got1)
	}
}
