package stringtest

import (
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	str := "中"
	// t.Log(str, len(str))
	// c := []rune(str)
	// t.Log(c, len(c))

	for _, v := range str {
		t.Logf("%[1]c %[1]x", v)
	}
}

func Test2(t *testing.T) {
	// 数字转字符串
	s := strconv.Itoa(10)
	t.Log(s)

	//字符串转数字
	if val, err := strconv.Atoi("10"); err == nil {
		t.Log(val + 10)
	}

}
