package slice_test

import "testing"

func Test1(t *testing.T) {

	var sli []int

	sli = append(sli, 1)

	t.Log(len(sli), cap(sli))

	sli1 := make([]int, 3, 5)
	t.Log(len(sli1), cap(sli1))

}

func Test2(t *testing.T) {
	numArr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	t.Log("numArr:", numArr)

	sli1 := numArr[3:5]
	t.Log("sli1:", len(sli1), cap(sli1))

	sli2 := numArr[2:8]
	t.Log("sli2:", len(sli2), cap(sli2))

	sli1[0] = 100
	t.Log(sli2)
}
