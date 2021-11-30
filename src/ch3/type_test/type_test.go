package type_test

import "testing"

func Test1(t *testing.T) {
	a := 1

	var b int = 1

	b = a

	t.Log(b)

}

func Test2(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(aPtr)
	t.Logf("%T %T", a, aPtr)
}

func Test3(t *testing.T) {
	var (
		a int = 1
		b int = 2
	)
	t.Log(a)
	t.Log(b)
}
