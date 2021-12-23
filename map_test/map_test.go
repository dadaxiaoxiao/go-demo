package map_test

import "testing"

func TestInitMap(t *testing.T) {
	// key:value
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

}
