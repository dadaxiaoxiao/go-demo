package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Execuatable
)

func TestConstantTry(t *testing.T) {
	t.Log(Readable, Writable)
}
