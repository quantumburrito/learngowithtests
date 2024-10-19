package main

import (
	"reflect"
	"testing"
)

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

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected: %d, Got %d, given %v and %v", want, got, firstSlice, secondSlice)
		}

	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected %v, Got %v", want, got)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2,9}

		checkSums(t,want, got)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3,4,5})
		want := []int{0,9}
		checkSums(t,want,got)
	})
	
}