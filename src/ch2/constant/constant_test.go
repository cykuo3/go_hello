package constant_test

import "testing"

const (
	A = 2 << iota
	B
	C
)

const (
	AA = 3 * iota
	BB
	CC
)

func Test_Constant(t *testing.T) {
	t.Log(AA)
	t.Log(BB)
	t.Log(CC)
}
