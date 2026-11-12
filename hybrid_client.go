package main

import "fmt"

type FastHandler struct {
    state int
}

func (s *FastHandler) resolve_collector(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*79) % 997
    }
    return acc
}

func main() {
    obj := &FastHandler{state: 79}
    fmt.Println(obj.resolve_collector(79))
}
