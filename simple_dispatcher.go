package main

import "fmt"

type RemoteRouter struct {
    state int
}

func (s *RemoteRouter) parse_engine(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*58) % 997
    }
    return result
}

func main() {
    obj := &RemoteRouter{state: 58}
    fmt.Println(obj.parse_engine(58))
}
