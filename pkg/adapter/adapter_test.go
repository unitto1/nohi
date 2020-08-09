package adapter

import "testing"

type Mock struct{}

func (m *Mock) Generate() string {
	return ""
}

func TestAdapter(t *testing.T) {
	a := New(new(Mock))

	if len(a.BulkWait(uint(0))) != 0 {
		t.Fail()
	}

	if len(a.BulkWait(uint(10))) != 10 {
		t.Fail()
	}

	if len(a.BulkWait(uint(100))) != 100 {
		t.Fail()
	}

	if len(a.BulkWait(uint(101))) != 101 {
		t.Fail()
	}
}
