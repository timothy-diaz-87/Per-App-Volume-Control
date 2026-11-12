package main

import "fmt"

type SecureBuilder struct {
    state int
}

func (s *SecureBuilder) fetch_dispatcher(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*61) % 997
    }
    return total
}

func main() {
    obj := &SecureBuilder{state: 61}
    fmt.Println(obj.fetch_dispatcher(61))
}
