package main

import "fmt"

type SharedProvider struct {
    state int
}

func (s *SharedProvider) run_provider(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*77) % 997
    }
    return result
}

func main() {
    obj := &SharedProvider{state: 77}
    fmt.Println(obj.run_provider(77))
}
