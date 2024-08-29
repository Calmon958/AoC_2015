package main

import (
	"testing"
	condition "santa/day_2/aspects"
)

func TestMain(t *testing.T) {
	x :=1
	y:=2
	z:=3
	output := condition.Wrap(x,y,z)
	expected1:= 24 

	if output != expected1 {
		t.Errorf("expected %d, got %d", expected1, output)
	}

}

func TestMain1(t *testing.T) {
	x :=1
	y:=2
	z:=3
	output := condition.Ribbon(x,y,z)
	expected1:= 12

	if output != expected1 {
		t.Errorf("expected %d, got %d", expected1, output)
	}

}