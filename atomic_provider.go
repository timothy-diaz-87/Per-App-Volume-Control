package main

import "fmt"

type AtomicProvider struct {
    state int
}

func (s *AtomicProvider) parse_handler(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*37) % 997
    }
    return acc
}

func main() {
    obj := &AtomicProvider{state: 37}
    fmt.Println(obj.parse_handler(37))
}
