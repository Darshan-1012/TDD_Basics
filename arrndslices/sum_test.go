package arrndslices

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	t.Run("The sum []int{1,2,3} = []int{6}", func(t *testing.T) {
		got := SumAll([]int{3, 0}, []int{3, 9})
		want := []int{3, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%v is got and %v is want", got, want)
		}
	})
}
func TestSumAllTrails(t *testing.T) {

	checksums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%v is got and %v is want ", got, want)
		}
	}
	t.Run("The sum  of trials", func(t *testing.T) {
		got := SumAllTrails([]int{2, 3}, []int{0, 9})
		want := []int{3, 9}
		checksums(t, got, want)
	})

	t.Run("The sum  of empty trials", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{1, 9})
		want := []int{0, 9}
		checksums(t, got, want)
	})
}
