package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectRepeat := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("got  %q but expected %q", repeated, expected)
		}
	}
	t.Run("repeat five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectRepeat(t, repeated, expected)
	})
	t.Run("repeat 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"
		assertCorrectRepeat(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	// Runs b.N times and measures execution time
	// and automatically determines "good" performance
	// To run: `go test -bench=.`
	// reported as 107.7 ns/op -> which means the fn takes on average 107.7 nanoseconds
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
