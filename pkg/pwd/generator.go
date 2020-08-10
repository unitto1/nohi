package pwd

import (
	"crypto/rand"

	"gitlab.com/unitto/nohi/internal/utils"
)

var (
	Special = []rune{
		33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
		58, 59, 60, 61, 62, 63, 64,
		91, 92, 93, 94, 95, 96,
		123, 124, 125, 126,
	} // 32 and 127 are skipped (space and delete).
	Digits = []rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	Upper  = []rune{
		65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77,
		78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	}
	Lower = []rune{
		97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109,
		110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
	}
)

type Generator struct {
	length      uint
	enableUpper bool
	enableDigit bool
	enableSpec  bool
}

func (g *Generator) Length() uint {
	return g.length
}

func (g *Generator) SetLength(s uint) {
	g.length = s
}

func (g *Generator) EnableUpper(f bool) {
	g.enableUpper = f
}

func (g *Generator) EnableDigit(f bool) {
	g.enableDigit = f
}

func (g *Generator) EnableSpecial(f bool) {
	g.enableSpec = f
}

func (g *Generator) Generate() string {
	choice := make([]rune, 0, len(Lower)+len(Upper)+len(Digits)+len(Special))
	choice = append(choice, Lower...)

	if g.enableDigit {
		choice = append(choice, Digits...)
	}

	if g.enableUpper {
		choice = append(choice, Upper...)
	}

	if g.enableSpec {
		choice = append(choice, Special...)
	}

	res := make([]rune, 0, g.length)
	choiceLen := uint8(len(choice))
	idxs := utils.RandInt8(rand.Reader, choiceLen, int(g.length))

	for _, idx := range idxs {
		res = append(res, choice[idx])
	}

	return string(res)
}

const defaultLength = 20

func New() *Generator {
	return &Generator{
		defaultLength,
		true,
		true,
		false,
	}
}
