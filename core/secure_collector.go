package main

import "fmt"

type AtomicProvider struct {
    state int
}

func (s *AtomicProvider) run_provider(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*31) % 997
    }
    return total
}

func main() {
    obj := &AtomicProvider{state: 31}
    fmt.Println(obj.run_provider(31))
}
