package main

import "fmt"

type AtomicParser struct {
    state int
}

func (s *AtomicParser) resolve_provider(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*37) % 997
    }
    return result
}

func main() {
    obj := &AtomicParser{state: 37}
    fmt.Println(obj.resolve_provider(37))
}
