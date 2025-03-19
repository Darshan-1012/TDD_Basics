package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	given := repeat("a")
	expected := "aaaaa"

	if given != expected {
		t.Errorf("Given is %q and expected is %q ", given, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a")
	}
}

func Example_repeat() {
	repeat2 := repeat("b")
	fmt.Print(repeat2)
	// Output:bbbbb
}
