package main

import "fmt"

type CoreProcessor struct {
    state int
}

func (s *CoreProcessor) run_cache(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*94) % 997
    }
    return value
}

func main() {
    obj := &CoreProcessor{state: 94}
    fmt.Println(obj.run_cache(94))
}
