package ds

import (
	"fmt"
	"sync"

	"github.com/go-openapi/errors"
)

type LinkedListFeatures interface {
	ToString() string
	AppendNode(data interface{}) error
	AppendNodes(data ...interface{}) error
	GetUnderlyingLinkedList() *LinkedList
}
type Node struct {
	Data interface{}
	Next *Node
}
type LinkedList struct {
	Root    *Node
	TailEnd *Node
}

var (
	linkedList *LinkedList
	once       sync.Once
)

func GetDefaultLinkedList() LinkedListFeatures {
	if linkedList == nil {
		initializeLinkedList()
	}
	return linkedList
}

func initializeLinkedList() {
	once.Do(func() {
		linkedList = createNewLL().GetUnderlyingLinkedList()
	})
}

func createNewLL() LinkedListFeatures {
	ll := new(LinkedList)
	ll.Root = new(Node)
	ll.Root.Data = nil
	ll.TailEnd = nil
	ll.Root.Next = nil
	return ll
}

func (ll *LinkedList) ToString() string {
	response := ""
	if ll.Root == nil {
		return response
	}
	curNode := ll.Root.Next
	for {
		if curNode != nil {
			response += fmt.Sprintf("%v", curNode.Data) + "->"
			curNode = curNode.Next
		} else {
			return response + "nil"
		}
	}
}

func createNewNode(data interface{}) *Node {
	node := new(Node)
	node.Next = nil
	node.Data = data
	return node
}
func (ll *LinkedList) AppendNode(data interface{}) error {
	if ll.Root == nil {
		return errors.New(503, "Invalid LinkedList, root node is nil")
	}
	newNode := createNewNode(data)
	if ll.TailEnd == nil {
		ll.TailEnd = newNode
		ll.Root.Next = newNode
	} else {
		ll.TailEnd.Next = newNode
		ll.TailEnd = newNode
	}
	return nil
}
func (ll *LinkedList) AppendNodes(nodesData ...interface{}) error {
	for _, data := range nodesData {
		err := ll.AppendNode(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ll *LinkedList) GetUnderlyingLinkedList() *LinkedList {
	return ll
}
