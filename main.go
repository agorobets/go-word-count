package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const (
	MAX_PARALLEL_LOADS = 5
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	re := regexp.MustCompile(`Go`)
	counter := NewCounter(re)

	dispatcher := NewDispatcher(counter, MAX_PARALLEL_LOADS)

	for scanner.Scan() {
		url := scanner.Text()
		dispatcher.StartLoadAndCount(url)
	}

	dispatcher.Wait()
	fmt.Printf("Total: %d\n", dispatcher.Total())
}
