package functest

import (
	"math/rand"
	"testing"
)

func randomInt() int {
	return rand.Intn(10)
}

func randomIntInt() (int, int) {
	return rand.Intn(10), rand.Intn(100)
}

func timeSpend(f func(op int) int) func(op int) int {
	return func(op1 int) int {
		return f(op1) + 100
	}
}

func f1(op int) int {
	return op
}

func Test1(t *testing.T) {
	t.Log(f1(10))
	t.Log(timeSpend(f1)(10))
}

func Sum(ops ...int) int {
	res := 0
	for _, op := range ops {
		res += op
	}
	return res
}

func Test2(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Test3(t *testing.T) {
	defer func() {
		t.Log("finally")
	}()

	panic("error")
	t.Log("main func")
}
