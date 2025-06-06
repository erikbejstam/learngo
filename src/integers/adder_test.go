package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}
}

// Testable Examples

func ExampleAdd() {
	sum := Add(5, 6)
	fmt.Println(sum)
	// Output: 11
}