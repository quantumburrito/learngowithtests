package iteration

import (
	"fmt"
	"strings"
	"testing"
) 

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeatedString := Repeat("a", 10)
	fmt.Println(repeatedString)
	// Output: aaaaaaaaaa
}

//Benchmark Repeat function
func BenchmarkRepeat(t *testing.B){
	for i := 0; i < t.N; i++ {
		Repeat("a", 5)
	}
}

// Play with strings
func TestClone(t *testing.T) {
	exampleString := "ExampleString"
	want := "ExampleString"
	got := strings.Clone(exampleString)

	if want != got {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func ExampleClone() {
	fmt.Println(strings.Clone("Hello Clone"))
	// Output: Hello Clone
}
