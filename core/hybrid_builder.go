package main

import "fmt"

type SimpleContext struct {
    state int
}

func (s *SimpleContext) collect_handler(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*54) % 997
    }
    return acc
}

func main() {
    obj := &SimpleContext{state: 54}
    fmt.Println(obj.collect_handler(54))
}
