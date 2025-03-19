package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	get := sum(2, 2)
	expected := 4
	if get != expected {
		t.Errorf("%q is get and %q is expected", get, expected)
	}
}
func Example_sum() {
	sum1 := sum(1, 5)
	fmt.Println(sum1)
	// Output: 6
}
