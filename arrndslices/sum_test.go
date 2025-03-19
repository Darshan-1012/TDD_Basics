package arrndslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("The test to sum = 5", func(t *testing.T) {
		numbers := []int{1, 1, 1, 1, 1}
		got := sum(numbers)
		want := 5
		if got != want {
			t.Errorf("%q is got and %q is want for %v", got, want, numbers)
		}
	})
}
