package ds

import (
	"fmt"
	"testing"
)

func TestLinkedList_GetLinkedList(t *testing.T) {
	if GetDefaultLinkedList() == nil {
		t.Log(fmt.Sprintf("GetDefaultLinkedList should return a non nil LinkedList"))
		t.Fail()
	}
}
func TestLinkedList_GetLinkedList_ValidateLinkedList(t *testing.T) {
	if GetDefaultLinkedList().GetUnderlyingLinkedList().TailEnd != nil {
		t.Log(fmt.Sprintf("Empty LinkedLists' TailEnd should be nil"))
		t.Fail()
	}
	if GetDefaultLinkedList().GetUnderlyingLinkedList().Root.Next != nil {
		t.Log(fmt.Sprintf("Empty LinkedLists' Root->Next should be nil"))
		t.Fail()
	}
}
func TestLinkedList_ToString_Nodata(t *testing.T) {
	ll := GetDefaultLinkedList()
	if ll.ToString() != "nil" {
		t.Log(fmt.Sprintf("Empty linkedlist ToString should return an \"nil\" string, got %v", ll.ToString()))
		t.Fail()
	}
}

func TestAppendNode_Success(t *testing.T) {
	ll := GetDefaultLinkedList()
	err := ll.AppendNode(10)
	if err != nil {
		t.Log("AppendNode for a properly initialized linked should not error")
		t.Fail()
	}
	if ll.ToString() != "10->nil" {
		t.Log(fmt.Sprintf("ToString() output invalid, %v", ll.ToString()))
		t.Fail()
	}
	err = ll.AppendNode(11)
	if err != nil {
		t.Log("AppendNode for a properly initialized linked should not error")
		t.Fail()
	}
	if ll.ToString() != "10->11->nil" {
		t.Log(fmt.Sprintf("ToString() output invalid, %v", ll.ToString()))
		t.Fail()
	}
}

func TestAppendNodes_Success(t *testing.T) {
	ll := createNewLL()
	err := ll.AppendNodes(10)
	if err != nil {
		t.Log("AppendNodes for a properly initialized linked should not error")
		t.Fail()
	}
	if ll.ToString() != "10->nil" {
		t.Log(fmt.Sprintf("ToString() output invalid, %v", ll.ToString()))
		t.Fail()
	}
	err = ll.AppendNodes(11, 12, 13)
	if err != nil {
		t.Log("AppendNodes for a properly initialized linked should not error")
		t.Fail()
	}
	if ll.ToString() != "10->11->12->13->nil" {
		t.Log(fmt.Sprintf("ToString() output invalid, %v", ll.ToString()))
		t.Fail()
	}
}
