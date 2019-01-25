package main

import "fmt"

func main() {

	l1 := ListNode{
		4,
		&ListNode{
			6,
			&ListNode{
				5,
				&ListNode{},
			},
		},
	}
	l2 := ListNode{
		8,
		&ListNode{
			1,
			&ListNode{
				9,
				&ListNode{},
			},
		},
	}
	fmt.Println(AddTwoNumbers(&l1, &l2))

}

type ListNode struct {
	Value    int
	NextNode *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, currentNode ListNode
	var current int

	for ; l1.NextNode != nil; l1 = l1.NextNode {
		l2 = l2.NextNode
		currentNode.Value = l1.Value + l2.Value + current%10
		currentNode.NextNode=&ListNode{}
		current = (l1.Value + l2.Value + current) / 10
		if result.NextNode==nil {
			result= currentNode
		}else {
			result.NextNode= &currentNode
		}
	}


	return &result;
}
