package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
		t.Log(a)
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 1; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("基数")
		case 1, 3:
			t.Log("偶数")
		default:
			t.Log("大于3的数")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 1; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("基数")
		case i%2 == 1:
			t.Log("偶数")
		default:
			t.Log("unknow")
		}
	}
}
