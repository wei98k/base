package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("a", "b")
	addEdge("c", "d")
	addEdge("a", "d")
	addEdge("d", "a")
	fmt.Println(hasEdge("a", "b"))
	fmt.Println(hasEdge("c", "d"))
	fmt.Println(hasEdge("a", "d"))
	fmt.Println(hasEdge("d", "a"))
	fmt.Println(hasEdge("x", "b"))
	fmt.Println(hasEdge("c", "d"))
	fmt.Println(hasEdge("x", "d"))
	fmt.Println(hasEdge("d", "x"))
}

func addEdge(from, to string) {
	edge := graph[from]
	if edge == nil {
		edge = make(map[string]bool)
		graph[from] = edge
	}
	edge[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
