package main

import "testing"

func TestSum(t *testing.T) {
	
	t.Run("Collection of any size", func(t *testing.T){
		numbers := []int{1,2,3}
		got := Sum(numbers) 
		want := 6

		if got != want {
			t.Errorf("Expected %d, got %d, given %v", want, got, numbers)
		}
	})
	
}

func TestSumAll(t *testing.T){
	t.Run("sum two slices, return two slices of sums", func(t *testing.T) {
		firstSlice := []int{1,2}
		secondSlice := []int{0,9}
		
		got := SumAll(firstSlice, secondSlice)
		want := []int{3,9}

		if want != got {
			t.Errorf("Expected: %d, Got %d, given %v and %v", want, got, firstSlice, secondSlice)
		}

	})
}