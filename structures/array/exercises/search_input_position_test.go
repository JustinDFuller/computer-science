package exercises

import "testing"

func TestSearchInputPosition(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   int
	}{
		{
			input:  []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
	}

	for _, test := range tests {
		got := SearchInsert(test.input, test.target)

		if got != test.want {
			t.Errorf("searchInsert(%v, %d) = %d, want: %d", test.input, test.target, got, test.want)
		}
	}
}
