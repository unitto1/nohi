package uuid4

import "github.com/google/uuid"

type Generator struct{}

func (g *Generator) Generate() string {
	return uuid.New().String()
}

func New() *Generator {
	return &Generator{}
}
