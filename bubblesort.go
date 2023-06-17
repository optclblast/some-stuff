package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var lng int
	fmt.Print("Slice size: ")
	fmt.Scan(&lng)

	rnd := randomSlice(lng)

	startTime := time.Now()

	fmt.Printf("Src slice: %v\n", rnd)

	fmt.Printf("Sorted slice: %v\n", sortTheSlice(rnd))
	runtime := time.Since(startTime)
	fmt.Printf("Runtime: %v%v%v", runtime.Minutes(), runtime.Seconds(), runtime.Milliseconds())
	fmt.Scan(&lng)
}

func sortTheSlice(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func randomNum() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000 - 1 + 1)
}

func randomSlice(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = randomNum()
		time.Sleep(2 * time.Microsecond)
	}
	return s
}
