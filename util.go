package util

import "container/list"

type IntegerStruct struct {
	Val     int
	NotNull bool
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Element struct {
	Node  *TreeNode
	Index int
}

func LevelOrder(nums []IntegerStruct) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	head := &TreeNode{
		Val:   nums[0].Val,
		Left:  nil,
		Right: nil,
	}
	queue := list.New()
	queue.PushBack(Element{
		Node:  head,
		Index: 0,
	})

	for queue.Len() > 0 {
		frontEl := queue.Front().Value.(Element)
		queue.Remove(queue.Front())

		// left child
		if (2*frontEl.Index+1) < n && nums[2*frontEl.Index+1].NotNull {
			leftChild := Element{
				Index: 2*frontEl.Index + 1,
				Node: &TreeNode{
					Val:   nums[2*frontEl.Index+1].Val,
					Left:  nil,
					Right: nil,
				},
			}
			frontEl.Node.Left = leftChild.Node
			queue.PushBack(leftChild)
		}

		// right child
		if (2*frontEl.Index+2) < n && nums[2*frontEl.Index+2].NotNull {
			rightChild := Element{
				Index: 2*frontEl.Index + 2,
				Node: &TreeNode{
					Val:   nums[2*frontEl.Index+2].Val,
					Left:  nil,
					Right: nil,
				},
			}
			frontEl.Node.Right = rightChild.Node
			queue.PushBack(rightChild)
		}
	}
	return head
}
