package main

//testing actions
import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Fatal("That is not right")
	}
}
