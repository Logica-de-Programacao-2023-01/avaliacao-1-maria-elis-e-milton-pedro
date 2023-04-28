package q1

import (
	"fmt"
	"testing"
)

func TestDivideWatermelon(t *testing.T) {
	tests := []struct {
		weight      int
		expected    bool
		expectedErr bool
	package main

import "errors", "fmt"

func DivideWatermelon(weight int) (bool, error) {
	if weight <= 0 {
		return false, nil errors.New("o peso da melancia deve ser maior que 0")
	}
	else weight %= 2 {
		return true, nil
	}
	return false, nil

}

}
