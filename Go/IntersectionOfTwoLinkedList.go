package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

func GetIntersectionNode(A, B *ListNode) *ListNode {
	for A != nil {
		for pB := B; pB != nil; pB = pB.next {
			if pB == A {
				return A
			}
		}
		A = A.next
	}

	return nil
}

func GetArrayFromInput(count int, prompt string) []int {
	if count == 0 {
		return []int{}
	}

	fmt.Print(prompt)

	v := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&v[i])
	}

	return v
}

func GetListFromArray(v []int) *ListNode {
	if len(v) == 0 {
		return nil
	}

	if len(v) == 1 {
		return &ListNode{v[0], nil}
	}

	pHead := &ListNode{v[0], nil}
	pCur := pHead
	for i := 1; i < len(v); i++ {
		pCur.next = &ListNode{v[i], nil}
		pCur = pCur.next
	}

	return pHead
}

func AddTail(pHead, pToBeAdded *ListNode) *ListNode {
	if pHead == nil {
		return pToBeAdded
	}

	pSaved := pHead
	for pHead.next != nil {
		pHead = pHead.next
	}

	pHead.next = pToBeAdded
	return pSaved
}

// func (node *ListNode) Print() {
// 	if node == nil {
// 		fmt.Println("")
// 	}

// 	fmt.Print(node.data)
// 	node = node.next
// 	for node != nil {
// 		fmt.Printf(", %d", node.data)
// 		node = node.next
// 	}
// 	fmt.Println("")
// }

func main() {
	for {
		fmt.Print("Number of nodes before intersection A and B: ")
		var c1, c2 int
		fmt.Scan(&c1, &c2)
		if c1 == 0 && c2 == 0 {
			break
		}

		vec := GetArrayFromInput(c1, "Please enter list1: ")
		headA := GetListFromArray(vec)

		vec = GetArrayFromInput(c2, "Please enter list2: ")
		headB := GetListFromArray(vec)

		fmt.Print("The number of nodes of the common list: ")
		var c3 int
		fmt.Scan(&c3)
		vec = GetArrayFromInput(c3, "Please enter the common list: ")
		headC := GetListFromArray(vec)

		headA = AddTail(headA, headC)
		//headA.Print()
		headB = AddTail(headB, headC)
		//headB.Print()

		pNode := GetIntersectionNode(headA, headB)
		fmt.Printf("The intersection node of two lists is ")
		if pNode == nil {
			fmt.Println(-1)
		} else {
			fmt.Println(pNode.data)
		}
	}
}
