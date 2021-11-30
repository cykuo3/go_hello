package loop_test

import "testing"

func TestLoop(t *testing.T) {
	// for a < 5 {
	// 	t.Log(a)
	// 	a++
	// }

	a := 5

	if a > 0 {
		t.Log(a)
	}
}
