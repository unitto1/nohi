package uuid4

import (
	"testing"

	"github.com/google/uuid"

	"gitlab.com/unitto/nohi/pkg/adapter"
)

func TestUUID4(t *testing.T) {
	g := New()
	r := g.Generate()

	if _, err := uuid.Parse(r); err != nil {
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
