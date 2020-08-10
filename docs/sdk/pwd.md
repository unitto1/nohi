# Password generator
## Constructor
```go
const defaultLength = 20

func New() *Generator {
	return &Generator{
		defaultLength,
		true,
		true,
		false,
	}
}
```

## Methods
```go
func (g *Generator) Generate() string
func (g *Generator) Length() uint
func (g *Generator) SetLength(s uint)
func (g *Generator) EnableUpper(f bool)
func (g *Generator) EnableDigit(f bool)
func (g *Generator) EnableSpecial(f bool)
```
