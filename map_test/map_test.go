package map_test

import "testing"

func TestInitMap(t *testing.T) {
	// key:value
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	// 新建空的字典
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	//新建切片
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))

}

// 判断字典是否存在某个键，没有返回值是0
func TestAccesNoExistStringKey(t *testing.T) {
	m1 := map[int]int{}
	// 不存的键，返回0
	t.Log(m1[1])
	m1[2] = 0
	// 键值为0
	t.Log(m1[2])
	m1[3] = 0
	// 判断是不存的键，还是键值是0
	if v, ok := m1[3]; ok {
		t.Logf("key is %d", v)
	} else {
		t.Log("key is not exist")
	}
}

// 键值对遍历
func TestTraveMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 4: 3}
	for key, value := range m1 {
		t.Log(key, value)
	}
}
