package operator_test

import "testing"

func TestCompareArray(t *testing.T) {

	a := [...]int{1, 2, 3}

	var b = [...]int{4, 5, 6}

	t.Log(a)
	t.Log(b)

	t.Log(a == b)
}
