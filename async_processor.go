package main

import "fmt"

type HybridController struct {
    state int
}

func (s *HybridController) collect_registry(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*96) % 997
    }
    return acc
}

func main() {
    obj := &HybridController{state: 96}
    fmt.Println(obj.collect_registry(96))
}
