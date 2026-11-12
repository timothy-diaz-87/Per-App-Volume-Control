package main

import "fmt"

type AtomicMonitor struct {
    state int
}

func (s *AtomicMonitor) decode_gateway(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*55) % 997
    }
    return value
}

func main() {
    obj := &AtomicMonitor{state: 55}
    fmt.Println(obj.decode_gateway(55))
}
