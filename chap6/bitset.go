package main

import (
	"bytes"
	"fmt"
)

type BitSet struct {
	array []uint64
}

func (s *BitSet) Add(x int) {
	index, bit := x/64, uint(x%64)
	if index >= len(s.array) {
		s.array = append(s.array, 0)
	}
	s.array[index] |= 1 << bit
}

func (s *BitSet) Has(x int) bool {
	index, bit := x/64, uint(x%64)
	return index < len(s.array) && s.array[index]&(1<<bit) != 0
}

func (s *BitSet) UnionWith(t *BitSet) {
	for index, data := range t.array {
		if index < len(s.array) {
			s.array[index] |= data
		} else {
			s.array = append(s.array, data)
		}
	}
}

func (s *BitSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for index, data := range s.array {
		if data == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if data&(1<<uint(j)) != 0 {
				//not first element
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*index+j)
			}
		}

	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x BitSet
	x.Add(1)
	x.Add(2)
	fmt.Println(x.Has(1))
	fmt.Println(x.Has(2))
	fmt.Println(x.String())
	fmt.Println(&x)
	fmt.Println(x)
	fmt.Println(32 << (^uint(0) >> 63))
	fmt.Println(32 << (^uint32(0) >> 63))
	fmt.Println(32 << (^uint64(0) >> 63))
	fmt.Println(^uint(1))
	fmt.Println(^uint(0))
	fmt.Println(^1)
}
