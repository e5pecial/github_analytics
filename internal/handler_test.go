package internal

import (
	"reflect"
	"testing"
)

func Test_sortByValue(t *testing.T) {
	type args struct {
		mapping map[int64]int
	}
	tests := []struct {
		name string
		args args
		want []Counter
	}{
		{"correct", args{map[int64]int{1: 1, 2: 0, 3: 2, 4: 4}}, []Counter{{4, 4}, {3, 2}, {1, 1}, {2, 0}}},
		{"one", args{map[int64]int{1: 1}}, []Counter{{1, 1}}},
		{"empty", args{map[int64]int{}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByValue(tt.args.mapping); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
