package customsort

import "testing"

func TestSortInPlace(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{"1 value", []int{1}, []int{1}},
		{"2 value", []int{2, 1}, []int{1, 2}},
		{"3 value", []int{7, 15, 1}, []int{1, 7, 15}},
		{"9 value", []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			SortInPlace(test.in)
			if len(test.in) != len(test.out) {
				t.Errorf("got %d, want %d", len(test.in), len(test.out))
			}
			for i := 0; i < len(test.in); i++ {
				if test.in[i] != test.out[i] {
					t.Errorf("got %v, want %v", test.in, test.out)
				}
			}
		})
	}

}
