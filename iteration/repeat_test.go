package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Basic test", func(t *testing.T) {
		repeated := Repeat("a", 4)
		expected := "aaaa"

		assertCorrect(t, repeated, expected)
	})

	t.Run("Make lower case", func(t *testing.T) {
		repeated := Repeat("A", 3)
		expected := "aaa"

		assertCorrect(t, repeated, expected)
	})

	t.Run("Trim string", func(t *testing.T) {
		repeated := Repeat(".? !u ", 2)
		expected := "uu"

		assertCorrect(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	Repeat("e", 3)
	fmt.Println("eee")
	// Output: eee
}

func assertCorrect(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("Got: %q, wanted: %q", repeated, expected)
	}
}