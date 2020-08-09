package adapter

import (
	"runtime"
	"sync"
)

type Generatorer interface {
	Generate() string
}

type BulkAdapter struct {
	Generatorer
	conc uint
	wg   *sync.WaitGroup
}

func (a *BulkAdapter) SetWorkerCount(n uint) {
	a.conc = n
}

func (a *BulkAdapter) worker(n uint, res chan string) {
	defer a.wg.Done()

	for i := uint(0); i < n; i++ {
		res <- a.Generate()
	}
}

func (a *BulkAdapter) Bulk(n uint, res chan string) {
	workers := a.conc
	if n != 0 && n < workers {
		workers = n
	}

	factor := n / workers

	for i := uint(0); i < workers; i++ {
		if i == workers-1 {
			factor += n % workers
		}

		a.wg.Add(1)

		go a.worker(factor, res)
	}
	a.wg.Wait()
	close(res)
}

func (a *BulkAdapter) BulkWait(n uint) []string {
	res := make([]string, 0, n)
	rChan := make(chan string, n)

	go a.Bulk(n, rChan)

	for el := range rChan {
		res = append(res, el)
	}

	return res
}

func New(gen Generatorer) *BulkAdapter {
	return &BulkAdapter{
		gen,
		uint(runtime.NumCPU()),
		&sync.WaitGroup{},
	}
}
