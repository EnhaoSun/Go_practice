package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// start_node, end_node, reverse
    var start, end, preTail *ListNode
    dummyHead := ListNode{}
    dummyHead.Next = head

	i := 1
	start, end = findKGroup(head, k, &i)
    head = end.Next
	if i == k {
        // previous group head and tail
		preTail = reverse(start, end, &dummyHead)
	}
    i = 1
    leftTails := false
	for head != nil {
		start, end = findKGroup(head, k, &i)
        if end == nil {
            head = end
        } else {
            head = end.Next
        }
		if i == k {
			preTail = reverse(start, end, preTail)
		} else {
            leftTails = true
            break
        }
		i = 1
	}
    if leftTails {
        preTail.Next = start
    }
	return dummyHead.Next
}

func findKGroup(head *ListNode, k int, i *int) (*ListNode, *ListNode) {
    var endPtr *ListNode
    startPtr := head
	for head.Next != nil && (*i) < k{
		head = head.Next
		(*i)++
	}
	if (*i) == k {
		endPtr = head
	}
    return startPtr, endPtr
}

func reverse(start *ListNode, end *ListNode, pre *ListNode) (*ListNode) {
    dummyTail := start  // the tail for this group
    var dummyHead *ListNode  // the head for this group
	for start != end {
		tmp := start.Next
		start.Next = dummyHead
		dummyHead = start
		start = tmp
	}
    start.Next = dummyHead
	dummyHead = start
    pre.Next = dummyHead
    //return tail
    return dummyTail
}

func printCur(node *ListNode) {
    for node != nil {
        fmt.Print(node.Val, " -> ")
        node = node.Next
    }
    fmt.Println("")
}

func main() {
    k := 2
    head := ListNode{Val: 1}
    headPtr := &head
    for i := 2; i < 6; i++ {
        node := ListNode{Val: i}
        headPtr.Next = &node
        headPtr = &node
    }

    headPtr = &head
    /*
    for headPtr != nil {
        fmt.Println(headPtr.Val)
        headPtr = headPtr.Next
    }
    */

    res := reverseKGroup(headPtr, k)
    printCur(res)
}
