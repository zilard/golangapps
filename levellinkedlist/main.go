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

/*

type Element struct {
	Value any
}


               i2         3
              / \        / \
             i1  i4     2   5
            /   /      /   /
           i0  i3     1   4

*/





func CreateLevelLinkedList(root *BinaryTreeNode) []*list.List {
	var result []*list.List
	var resslice [][]int

	current := list.New()
	curslice := []int{}

	// PushBack inserts a new element at the back of a list
	if root != nil {
		current.PushBack(root)
		curslice = append(curslice, root.Value.(int))
	}


	for current.Len() > 0 {

		// append -> appends new elements to a slice
		result = append(result, current)
		resslice = append(resslice, curslice)

		parents := current
		current = list.New()
		curslice = []int{}

		// func (l *List) Front() *Element -  Front() returns the first element of list or nil if the list is empty.
		// type List struct {..} - List represents a doubly linked list.
		// type Element struct {..}  - Element is an element of a linked list.
		// func (e *Element) Next() *Element - Next returns the next list element or nil.

		for parent := parents.Front(); parent != nil; parent = parent.Next() {

			node := parent.Value.(*BinaryTreeNode)  // explicit type casting

			if node.Left != nil {
				current.PushBack(node.Left)
				curslice = append(curslice, node.Left.Value.(int))
			}

			if node.Right != nil {
				current.PushBack(node.Right)
                                curslice = append(curslice, node.Right.Value.(int))
			}

		}

		// --------------
		/*
		curslice := []int{}
		for c := current.Front(); c != nil; c = c.Next() {
			//fmt.Printf("%T - %T\n", c, c.Value)
                        v := c.Value.(*BinaryTreeNode).Value
			curslice = append(curslice, v.(int))
		}
		fmt.Printf("curslice %v\n", curslice)
		*/
		// ---------------

	}

	fmt.Printf("resslice: %v\n", resslice)

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

        fmt.Printf("tree: %v\n", tree)


	listsArray := CreateLevelLinkedList(tree)

	fmt.Printf("listsArray: %v\n", listsArray)


	var list1Items []int
	for item := listsArray[0].Front(); item != nil; item = item.Next() {
		list1Items = append(list1Items, item.Value.(*BinaryTreeNode).Value.(int))
	}
	fmt.Printf("list1Items: %v\n", list1Items)


        var list2Items []int
        for item := listsArray[1].Front(); item != nil; item = item.Next() {
                list2Items = append(list2Items, item.Value.(*BinaryTreeNode).Value.(int))
        }
        fmt.Printf("list2Items: %v\n", list2Items)


        var list3Items []int
        for item := listsArray[2].Front(); item != nil; item = item.Next() {
                list3Items = append(list3Items, item.Value.(*BinaryTreeNode).Value.(int))
        }
        fmt.Printf("list3Items: %v\n", list3Items)

}

