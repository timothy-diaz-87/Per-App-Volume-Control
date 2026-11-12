package main

import "fmt"

type LocalParser struct {
    state int
}

func (s *LocalParser) handle_processor(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*74) % 997
    }
    return total
}

func main() {
    obj := &LocalParser{state: 74}
    fmt.Println(obj.handle_processor(74))
}
