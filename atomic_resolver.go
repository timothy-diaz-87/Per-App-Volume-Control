package main

import "fmt"

type AtomicClient struct {
    state int
}

func (s *AtomicClient) parse_provider(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*31) % 997
    }
    return count
}

func main() {
    obj := &AtomicClient{state: 31}
    fmt.Println(obj.parse_provider(31))
}
