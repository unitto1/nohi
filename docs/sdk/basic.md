# Basic Abstractions
## `adapter.Generatorer` interface
It's main abstraction for all generators. To implement this interface
your structure should have only one method: `Generate`.
Implementation details is private except for setters used for a setup.

> **NOTE:** use constructors for initialization of generators

```go
type Generatorer interface {
	Generate() string
}
```

## `adapter.BulkAdapter` structure
It's a wrapper around the generator which implements bulk generation functionality.

This structure implements given public methods:
```go
func (a *BulkAdapter) SetWorkerCount(n uint) // default to runtime.NumCPU()
func (a *BulkAdapter) Bulk(n uint, res chan string)
func (a *BulkAdapter) BulkWait(n uint) []string
// embed `Generatorer` itself too
func (g *Generator) Generate() string
```

As you probably already guessed based on signatures - the main difference between
`Bulk` and `BulkWait` it's results transport way:

* `Bulk` uses channel for sending results (it also close the channel so you can use `range`)
* `BulkWait` returns results in sync mode like a regular slice

### Example
```go
// code from cmd/uuid4.go
a := adapter.New(uuid4.New())
a.SetWorkerCount(workers)

res := make(chan string, count)
go a.Bulk(count, res)
utils.Output(res, out)
```
