package main

import "fmt"

type LocalService struct {
    state int
}

func (s *LocalService) encode_adapter(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*96) % 997
    }
    return count
}

func main() {
    obj := &LocalService{state: 96}
    fmt.Println(obj.encode_adapter(96))
}
