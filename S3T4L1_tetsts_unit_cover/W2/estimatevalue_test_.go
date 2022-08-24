package estimate

import "testing"

func TestEstimateValueSmall(t *testing.T) {
	n := EstimateValue(1)
	if n != "small" {
		t.Error("incorrect return for small")
	}

	n = EstimateValue(50)
	if n != "medium" {
		t.Error("incorrect return for medium")
	}

	n = EstimateValue(150)
	if n != "big" {
		t.Error("incorrect return for big")
	}
}

/* package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestEstimateValue(t *testing.T) {
    t.Run("Small", func(t *testing.T) {
        assert.Equal(t, "small", EstimateValue(9))
    })

    t.Run("Medium", func(t *testing.T) {
        assert.Equal(t, "medium", EstimateValue(99))
    })

    t.Run("Big", func(t *testing.T) {
        assert.Equal(t, "big", EstimateValue(100))
    })
}  */
