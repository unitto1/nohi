package hid

import (
	"crypto/rand"
	"strings"

	"gitlab.com/unitto/nohi/internal/utils"
	"gitlab.com/unitto/nohi/pkg/pwd"
)

type Generator struct {
	pwd.Generator
	sep string
}

func (g *Generator) Separator() string {
	return g.sep
}

func (g *Generator) SetSeparator(s string) {
	g.sep = s
}

func (g *Generator) Generate() string {
	res := make([]string, 0, 3)
	leftLen := uint8(len(left))
	rightLen := uint8(len(right))

begin:
	leftIdx := utils.RandInt8(rand.Reader, leftLen, 1)

	rightIdx := utils.RandInt8(rand.Reader, rightLen, 1)
	l, r := left[leftIdx[0]], right[rightIdx[0]]

	// from original algorithm
	// Steve Wozniak is not boring
	if l == "boring" && r == "wozniak" {
		goto begin
	}

	res = append(res, l)
	res = append(res, r)

	if g.Length() > 0 {
		res = append(res, g.Generator.Generate())
	}

	return strings.Join(res, g.sep)
}

const defaultSuffixLength = 5

func New() *Generator {
	g := pwd.New()
	g.SetLength(defaultSuffixLength)
	g.EnableUpper(false)
	g.EnableDigit(true)
	g.EnableSpecial(false)

	return &Generator{
		*g,
		"_",
	}
}
