package array_test

import "testing"

func TestArryInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr1[1] = 5
	arr3 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArryTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}

	// for i := 0; i < len(arr3); i++ {
	// 	t.Log(arr3[i])
	// }
	// idx 下标，元素 值
	// for idx, e := range arr3 {
	// 	t.Log(idx, e)
	// }
	for _, e := range arr3 {
		t.Log(e)
	}
}

// 数组截取
func TestArrySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[1:len(arr3)]
	t.Log(arr3_sec)

}
