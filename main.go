package main

import (
	"fmt"

	"github.com/heckdevice/ds-al/ds"
)

func linkedList() {
	ll := ds.GetDefaultLinkedList()
	ll.AppendNode(1)
	ll.AppendNode(2)
	ll.AppendNodes(3, 4, 5, 6)
	fmt.Println(ll.ToString())
}

func filter(data []int, f func(int) bool) []int {
	var filteredData []int
	for _, val := range data {
		if f(val) == true {
			filteredData = append(filteredData, val)
		}
	}
	return filteredData
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func isOdd(n int) bool {
	return !isEven(n)
}
func main() {
	//linkedList()
	var dataPoints []int
	i := 0
	for i < 100 {
		i++
		dataPoints = append(dataPoints, i)
	}
	fmt.Println(dataPoints)
	evenNums := filter(dataPoints, isEven)
	oddNums := filter(dataPoints, isOdd)
	fmt.Println(evenNums)
	fmt.Println(oddNums)
}
