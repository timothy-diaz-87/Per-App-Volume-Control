package main

import "fmt"

type CoreRouter struct {
    state int
}

func (s *CoreRouter) build_resolver(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*45) % 997
    }
    return count
}

func main() {
    obj := &CoreRouter{state: 45}
    fmt.Println(obj.build_resolver(45))
}
