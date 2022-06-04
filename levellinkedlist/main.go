package main

import (
	"container/list"
	"math"
	"fmt"
)


type BinaryTreeNode struct {
        Value interface{}
        Left *BinaryTreeNode
        Right *BinaryTreeNode
}




func CreateLevelLinkedList(root *BinaryTreeNode) []*list.List {
	var result []*list.List

	current := list.New()

	// PushBack inserts a new element at the back of a list
	if root != nil {
		current.PushBack(root)
	}


	for current.Len() > 0 {

		// append -> appends new elements to a slice
		result = append(result, current)

		parents := current
		current = list.New()

		// Front returns the first element of list or nil if the list is empty.
		for parent := parents.Front(); parent != nil; parent = parent.Next() {

			node := parent.Value.(*BinaryTreeNode)  // explicit type casting

			if node.Left != nil {
				current.PushBack(node.Left)
			}

			if node.Right != nil {
				current.PushBack(node.Right)
			}

		}

	}

	return result

}




func GetBinaryTreeFromSortedArray(array []int) *BinaryTreeNode {
	return createMinimalBinarySubTreeFromArray(IntToInterfaceSlice(array), 0, len(array)-1)
}




func IntToInterfaceSlice(array []int) []interface{} {

	// creating a slice with make
	// make([]int, 5)
	// make([]int, 0, 5)  length 0    capacity 5

	result := make([]interface{}, len(array))
	for index, value := range array {
		result[index] = value
	}
	return result
}

/*

   [1,2,3,4,5]
    0.1.2.3.4

   mid = (0+4)/2 = 2
   
   Ceil of a number is the least integer value greater than or equal to that number.

   &BinaryTreeNode{ Value: array[2]
                    Left:   0..mid-1  0..1 => BinaryTreeNode{array[1],0..0,2..1} => BinaryTreeNode{array[0]..}
                    Right:  mid+1..4  3..4 => BinaryTreeNode{array[4],3..3,5..4} => BinaryTreeNode{array[3]..}


               i2         3
              / \        / \
             i1  i4     2   5
            /   /      /   /
           i0  i3     1   4

*/


func createMinimalBinarySubTreeFromArray(array []interface{},
					 start int,
					 end int) *BinaryTreeNode {
	if end < start {
		return nil
	}

	var mid = int(math.Ceil(float64(end+start) / float64(2)))

	return &BinaryTreeNode{
		Value: array[mid],
		Left:  createMinimalBinarySubTreeFromArray(array, start, mid-1),
		Right: createMinimalBinarySubTreeFromArray(array, mid+1, end),
	}
}



func main() {

	expected := []int{1, 2, 3, 4, 5}

	tree := GetBinaryTreeFromSortedArray(expected)


	//listsArray := CreateLevelLinkedList(tree)


	fmt.Printf("tree: %v\n", tree)



}



