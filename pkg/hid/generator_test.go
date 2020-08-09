package hid

import (
	"strings"
	"testing"

	"gitlab.com/unitto/nohi/pkg/adapter"
)

func TestHumanic(t *testing.T) {
	g := New()
	g.SetSeparator("_")
	g.EnableSpecial(false)
	parts := strings.Split(g.Generate(), "_")

	if len(parts) != 3 {
		t.Fail()
	}
}

const genCount = uint(1000)

func BenchmarkLoop(b *testing.B) {
	bg := adapter.New(New())

	for n := 0; n < b.N; n++ {
		for i := uint(0); i < genCount; i++ {
			bg.Generate()
		}
	}
}

func BenchmarkBulk(b *testing.B) {
	bg := adapter.New(New())

	for n := 0; n < b.N; n++ {
		r := make(chan string, genCount)
		go bg.Bulk(genCount, r)

		for range r {
		}
	}
}

func BenchmarkWait(b *testing.B) {
	bg := adapter.New(New())

	for n := 0; n < b.N; n++ {
		bg.BulkWait(genCount)
	}
}
