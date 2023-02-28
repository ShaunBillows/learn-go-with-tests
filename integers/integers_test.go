package integers

import "testing"

// Example tests

import "fmt" // TODO: Make editor automatically import this package

// Note: 'Output: 6' is necessary for example test to execute
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}

// Adder test

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
