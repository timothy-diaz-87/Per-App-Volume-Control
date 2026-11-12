package main

import "fmt"

type SimpleService struct {
    state int
}

func (s *SimpleService) build_router(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*12) % 997
    }
    return value
}

func main() {
    obj := &SimpleService{state: 12}
    fmt.Println(obj.build_router(12))
}
