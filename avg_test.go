package main

import (
	"log"
	"testing"
)

func TestAvg(t *testing.T) {
	xs := []float64{98, 93, 77, 82, 83}

	log.Println(avg(xs))

	if avg(xs) != 86.6 {
		t.Fail()
	}
}

func TestAvgEmpty(t *testing.T) {
	xs := []float64{}

	if avg(xs) != 0 {
		t.Fail()
	}
}
