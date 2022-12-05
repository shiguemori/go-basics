package main

import (
	"fmt"
	"sort"
	"sync"
)

// divide a slice into two slices
func divideSlice(list []int) (left []int, right []int) {
	mid := len(list) / 2
	left = list[:mid]
	right = list[mid:]
	return
}

func findMaxSum(numbers []int) int {
	wg := sync.WaitGroup{}
	if len(numbers) > 1000 {
		left, right := divideSlice(numbers)
		wg.Add(2)
		go func() {
			sort.Ints(left)
			wg.Done()
		}()
		go func() {
			sort.Ints(right)
			wg.Done()
		}()

		num1 := left[len(left)-1]
		num2 := right[len(right)-1]
		num3 := left[len(left)-2]
		num4 := right[len(right)-2]

		newSlice := []int{num1, num2, num3, num4}
		wg.Wait()
		return findMaxSum(newSlice)
	}
	sort.Ints(numbers)
	return numbers[len(numbers)-1] + numbers[len(numbers)-2]
}

func main() {
	list := []int{5, 9, 7, 11}
	fmt.Println(findMaxSum(list))
}
