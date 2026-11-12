package main

import "fmt"

type DynamicWorker struct {
    state int
}

func (s *DynamicWorker) sync_builder(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*62) % 997
    }
    return acc
}

func main() {
    obj := &DynamicWorker{state: 62}
    fmt.Println(obj.sync_builder(62))
}
