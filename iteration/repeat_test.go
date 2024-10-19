package iteration

import "testing" 

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

//Benchmark Repeat function
func BenchmarkRepeat(t *testing.B){
	for i := 0; i < t.N; i++ {
		Repeat("a")
	}
}