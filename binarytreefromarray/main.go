package main

import (
	"math"
	"fmt"
)

type BinaryTreeNode struct {
	Value interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode

}


func GetBinaryTreeFromSortedArray(array []int) *BinaryTreeNode {
	return createMinimalBinarySubTreeFromArray(IntToInterfaceSlice(array), 0, len(array)-1)
}


func IntToInterfaceSlice(array []int) []interface{} {
	result := make([]interface{}, len(array))

/*
	for i := 0; i < len(array); i++ {
		result[i] = array[i]
	}
*/

	for index, value := range array {
		result[index] = value
	}

	return result
}


func createMinimalBinarySubTreeFromArray(array []interface{}, start int, end int) *BinaryTreeNode {

	if end < start {
		return nil
	}

	var mid = int(math.Ceil(float64(end+start) / float64(2)))

	return &BinaryTreeNode{
		Value: array[mid],
		Left: createMinimalBinarySubTreeFromArray(array, start, mid-1),
		Right: createMinimalBinarySubTreeFromArray(array, mid+1, end),
	}

}


func DepthFirstSearchBinaryTree(node *BinaryTreeNode) []interface{} {
	if node == nil {
		return []interface{}{}
	}

	return append(DepthFirstSearchBinaryTree(node.Left),
			append([]interface{}{node.Value},
				DepthFirstSearchBinaryTree(node.Right)...)...)

}



func main() {

	//fmt.Printf("4.5 Ceil: %v\n", math.Ceil(4.5))
	//fmt.Printf("4.5 Floor: %v\n", math.Floor(4.5))

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tree := GetBinaryTreeFromSortedArray(expected)

	//fmt.Printf("interface slice: %v\n", IntToInterfaceSlice(expected))

	searchResult := DepthFirstSearchBinaryTree(tree)

	fmt.Printf("Depth First Search result: %v\n", searchResult)

        fmt.Printf("append([][]int{}, []int{1,2,3}) => %v\n", append([][]int{}, []int{1,2,3}) )
	fmt.Printf("append([]int{}, []int{1,2,3}...) => %v\n", append([]int{}, []int{1,2,3}...) )

}

