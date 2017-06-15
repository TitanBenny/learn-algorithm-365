package sort

import (
	"reflect"
	"testing"
)

func Test_bubble_sort(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want []int
	}{
		// TODO: Add test cases.
		{name: "bubbleSort",
			args: []int{5, 7, 3, 8, 4, 2, 23, 254, 21, 221, 6, 24, 68, 20},
			want: []int{254, 221, 68, 24, 23, 21, 20, 8, 7, 6, 5, 4, 3, 2},
		},
	}
	for _, tt := range tests {
		if got := bubble_sort(tt.args); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. bubble_sort() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
