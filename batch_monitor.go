package main

import "fmt"

type SimpleClient struct {
    state int
}

func (s *SimpleClient) flush_engine(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*18) % 997
    }
    return result
}

func main() {
    obj := &SimpleClient{state: 18}
    fmt.Println(obj.flush_engine(18))
}
