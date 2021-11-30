package arr_test

import "testing"

func Test1(t *testing.T) {

	var arr [2]int
	t.Log(arr)

	arr1 := [2]int{1, 2}
	t.Log(arr1)

	arr2 := [...][2]int{{2, 2}, {1, 1}}

	t.Log(arr2)

	var arr4 = [2]int{1, 2}

	t.Log(arr4)

}

func Test2(t *testing.T) {
	arr1 := [...]int{1, 2, 3, 4, 5}

	// for i := 0; i < len(arr1); i++ {
	// 	t.Log(arr1[i])
	// }
	// for _, e := range arr1 {
	// 	t.Log(e)
	// }

	t.Log(arr1[2:3])
}
