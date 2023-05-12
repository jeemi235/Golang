package master

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(2, 2)
	if sum != 4 {
		t.Errorf("Expected %v, got %v", 4, sum)
	}
}
