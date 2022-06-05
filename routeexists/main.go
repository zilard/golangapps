package main

import (
	"container/list"
	"fmt"
)


type Graph struct {
	Nodes []GraphNode
}

type GraphNode struct {
	Value    interface{}
	Children []interface{}
}

func (graph *Graph) GetNodeFromName(name string) GraphNode {
	for _, node := range graph.Nodes {
		if node.Value == name {
			return node
		}
	}
	return GraphNode{name, []interface{}{}}
}


func (graph *Graph) RouteExists(start string, end string) bool {
	root := graph.GetNodeFromName(start)

	visitedNodes := map[interface{}]bool{}

	queue := list.New()
	queue.PushBack(root)


	// func (l *List) Back() *Element
	// Back returns the last element of list l or nil if the list is empty.

	// func (l *List) Remove(e *Element) any
	// Remove removes e from l if e is an element of list l.
	// It returns the element value e.Value. The element must not be nil.

	for queue.Len() != 0 {
		root := queue.Remove(queue.Back()).(GraphNode)
		visitedNodes[root.Value] = true

		for _, nodeName := range root.Children {
			node := graph.GetNodeFromName(nodeName.(string))
			if _, ok := visitedNodes[node.Value]; !ok {
				visitedNodes[node.Value] = true
				queue.PushBack(node)
			}

			if node.Value == end {
				return true
			}
		}


	}

	return false

}


func StringToInterfaceSlice(array []string) []interface{} {
	result := make([]interface{}, len(array))
	for index, value := range array {
		result[index] = value
	}

	return result

}


func main() {

	graph := Graph{
		[]GraphNode{
			{
				Value: "0",
				Children: StringToInterfaceSlice([]string{"1", "4", "5"}),
			},
			{
				Value: "1",
				Children: StringToInterfaceSlice([]string{"4", "3"}),
			},
			{
                                Value: "2",
                                Children: StringToInterfaceSlice([]string{"1"}),
			},
			{
                                Value: "3",
                                Children: StringToInterfaceSlice([]string{"2", "4" }),
			},
			{
                                Value: "4",
                                Children: []interface{}{},
			},
		},
	}


	fmt.Printf("route exists between 0 => 2 ?   %v\n", graph.RouteExists("0", "2"))
        fmt.Printf("route exists between 1 => 1 ?   %v\n", graph.RouteExists("1", "1"))
        fmt.Printf("route exists between 1 => 0 ?   %v\n", graph.RouteExists("1", "0"))
        fmt.Printf("route exists between 3 => 0 ?   %v\n", graph.RouteExists("3", "0"))
}



