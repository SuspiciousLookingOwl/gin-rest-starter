package compA

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(3, 5)
	if sum != 8 {
		t.Errorf("Sum(3,5) = %d, expected 8", sum)
	}
}
