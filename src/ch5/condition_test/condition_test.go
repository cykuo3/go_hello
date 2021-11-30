package dondition_test

import "testing"

func Test_condition(t *testing.T) {

	if a := 1 == 1; a {
		t.Log("1==1")
	}
}

func Test_switch(t *testing.T) {
	a := 5
	switch {
	case a > 5:
		t.Log(">5")
	case a < 5:
		t.Log("<5")
	case a == 5:
		t.Log("=5")
	}

	for a := 0; a < 5; a++ {
		switch a {
		case 1, 3:
			t.Log("single")
		case 0, 2, 4:
			t.Log("double")
		}
	}
}
