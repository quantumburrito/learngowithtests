package iteration

import (
	"fmt"
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