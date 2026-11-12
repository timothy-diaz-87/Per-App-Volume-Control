package main

import "fmt"

type HybridAdapter struct {
    state int
}

func (s *HybridAdapter) handle_loader(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*72) % 997
    }
    return total
}

func main() {
    obj := &HybridAdapter{state: 72}
    fmt.Println(obj.handle_loader(72))
}
