package main

import (
	"fmt"
	"math/rand"
)

//快速排序

func quickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	base := data[0]
	l, r := 0, len(data)-1
	for i := 1; i <= r; {
		if data[i] > base {
			data[i], data[r] = data[r], data[i]
			r--
		} else {
			data[i], data[l] = data[l], data[i]
			l++
			i++
		}
	}
	quickSort(data[:l])
	quickSort(data[l+1:])
}

func main() {
	s := make([]int, 0, 8)
	for i := 0; i < 8; i++ {
		s = append(s, rand.Intn(100))
	}
	fmt.Println(s)
	quickSort(s)
	fmt.Println(s)
}
