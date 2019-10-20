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
func main() {
	linkedList()
}
