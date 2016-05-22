package main

import "fmt"

type Node struct {
    val int
    next *Node
    prev *Node
}

type LinkedList struct {
   head *Node
   tail *Node
}

func NewNode(value int, next *Node, prev *Node) *Node {
    return &Node{ val: value, next: next, prev: prev }
}

func NewLinkedList() *LinkedList{
    head := new(Node)
    tail := new(Node)
    tail.prev = head
    head.next = tail

    list := LinkedList{head: head, tail: tail}
    return &list
}

func (list *LinkedList) Push(val int) *Node{
    newNode := NewNode(val, list.tail, list.tail.prev)
    list.tail.prev.next = newNode
    list.tail.prev = newNode
    return newNode
}

func (list *LinkedList) Insert(index int, x int) bool {
    if index == -1 {
        return false
    }

    prevNode := list.head.search(index-1)

    if prevNode == nil {
        return false
    }
    nextNode := prevNode.next

    newNode := NewNode(x, nextNode, prevNode)
    prevNode.next = newNode
    nextNode.prev = newNode

    return true
}

func (node *Node) search(index int) *Node {
    i := 0

    // reject head
    if node.prev == nil {
        node = node.next
    }

    // While a node exist and not tail
    for node != nil && node.next != nil{
        if i == index {
            return node
        }
        i++
        node = node.next
    }
    return nil
}

/**
*   Try to find the node on the given index
*   @param index
*   @return target.value if found or nil if not found
*/
func (list *LinkedList) get(index int) int {
    targetNode := list.head.search(index)
    if targetNode == nil {
        return -1
    }

    return targetNode.val
}

func (list *LinkedList) Delete(index int) bool {
    prevNode := list.head.search(index-1)

    if prevNode == nil {
        return false
    }

    prevNode.next = prevNode.next.next
    return true
}

func (list *LinkedList) Size() int {
    var size = 0
    current := list.head.next

    for current.next != nil {
        size++
        current = current.next
    }

    return size
}

func (list *LinkedList) print() {
    fmt.Println("Printing List Values. Size: ", list.Size())
    fmt.Println("*********************")

    current := list.head.next
    for (current != nil && current != list.tail){
        fmt.Println("value=", current.val)
        current = current.next
    }
}


func main() {
    list := NewLinkedList()
                     // idx
    list.Push(100)      //0
    list.Push(101)      //1
    list.Push(102)      //2
    list.Push(103)      //3
    list.Insert(4, 104) // 4
    list.Insert(6, 105)  // this should not be added

    list.print()

    list.Delete(4)
    list.Delete(1)
    list.print()
}