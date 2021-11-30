package maptest_test

import (
	"testing"
)

func Test1(t *testing.T) {

	// m1 := map[string]int{"1": 1, "2": 2, "3": 3}
	// t.Log(m1)

	// m2 := map[string]int{}
	// t.Log(m2)

	// m3 := make(map[string]int, 10)
	// t.Log(m3)

	m4 := make(map[int]int, 10)
	m4[1] = 1
	m4[3] = 1
	t.Log(m4)

	// if _, ok := m4[3]; ok {
	// 	t.Log(m4[3])
	// } else {
	// 	t.Log("not exits")
	// }

	for k, v := range m4 {
		t.Log(k, v)
	}

}

func Test3(t *testing.T) {
	m1 := map[int]func(op int) int{}

	m1[0] = func(op int) int {
		return op * 2
	}

	m1[1] = func(op int) int {
		return op * op
	}

	r1 := m1[0](1)
	t.Log(r1)

	r2 := m1[1](1)
	t.Log(r2)

}

func Test4(t *testing.T) {

	s0 := map[int]bool{}
	s0[1] = true

	// n := 2
	// if s0[n] {
	// 	t.Logf("%d exist", n)
	// } else {
	// 	t.Logf("%d not exist", n)
	// }

	t.Log(len(s0))

	delete(s0, 1)

	t.Log(len(s0))

}
