package main

import (
	"fmt"
)

func bsOccurrence(array []int, target int, lowindex int, highindex int) int {
	mid := int((lowindex + highindex) / 2)
	if array[mid] > target {
		return bsOccurrence(array, target, lowindex, mid)
	} else if array[mid] < target {
		return bsOccurrence(array, target, mid+1, highindex)
	} else {
		return mid
	}
}

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	target := 22
	minindex := 0
	maxindex := len(primes) - 1
	if target > primes[maxindex] || target < primes[minindex] {
		fmt.Println("the target is outside of the array")
	} else {
		fmt.Printf("Eureka !!, the target was foud at index : %d", bsOccurrence(primes, target, minindex, maxindex))
	}
}
