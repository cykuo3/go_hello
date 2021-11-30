package objtest

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test1(t *testing.T) {
	u1 := User{"xiaoming", 10}
	fmt.Printf("main address is %x\n", unsafe.Pointer(&u1.name))
	u1.login1()
	u1.login()

	// u2 := User{name: "xiaoming", age: 10}
	// t.Log(u2)

	// u3 := new(User)
	// t.Log(u3)

}

type User struct {
	name string
	age  int
}

func (u User) login() {
	fmt.Printf("address is %x\n", unsafe.Pointer(&u.name))
	fmt.Printf("%s login success\n", u.name)
}

func (u *User) login1() {
	fmt.Printf("address is %x\n", unsafe.Pointer(&u.name))
	fmt.Printf("%s login success\n", u.name)
}
