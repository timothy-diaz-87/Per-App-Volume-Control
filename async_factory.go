package main

import "fmt"

type HybridContext struct {
    state int
}

func (s *HybridContext) parse_loader(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*75) % 997
    }
    return acc
}

func main() {
    obj := &HybridContext{state: 75}
    fmt.Println(obj.parse_loader(75))
}
