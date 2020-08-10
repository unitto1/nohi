# Human-friendly ID generator
## Constructor
```go
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
```

## Methods
```go
func (g *Generator) Generate() string
func (g *Generator) Separator() string
func (g *Generator) SetSeparator(s string)
func (g *Generator) Length() uint
func (g *Generator) SetLength(s uint)
func (g *Generator) EnableUpper(f bool)
func (g *Generator) EnableDigit(f bool)
func (g *Generator) EnableSpecial(f bool)
```
