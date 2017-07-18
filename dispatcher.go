package main

import (
	"fmt"
	"sync"
)

type Dispatcher struct {
	tokens  chan struct{}
	counter *Counter
	total   int

	wg sync.WaitGroup
	mu sync.Mutex
}

func NewDispatcher(counter *Counter, maxTokens int) *Dispatcher {
	tokens := make(chan struct{}, maxTokens)
	for i := 0; i < maxTokens; i++ {
		select {
		case tokens <- struct{}{}:
		default:
		}
	}

	return &Dispatcher{
		tokens:  tokens,
		counter: counter,
	}
}

func (d *Dispatcher) StartLoadAndCount(url string) {
	<-d.tokens

	d.wg.Add(1)
	go func() {
		defer d.wg.Done()

		if count, err := d.counter.LoadAndCount(url); err != nil {
			fmt.Printf("Error for %s: %s\n", url, err)
		} else {
			d.mu.Lock()
			d.total += count
			d.mu.Unlock()

			fmt.Printf("Count for %s: %d\n", url, count)
		}

		d.tokens <- struct{}{}
	}()
}

func (d *Dispatcher) Wait() {
	d.wg.Wait()
}

func (d *Dispatcher) Total() int {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.total
}
