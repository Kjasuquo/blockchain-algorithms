package graph

import (
	"fmt"
	"time"
)

func FindShortestPath(graphMap *Map) uint {
	// Keeping track of time taken to execute the function
	startTime := time.Now()

	// Initialize all nodes to have a minLengthFound of uint max value
	for _, node := range graphMap.AllNodes {
		node.MinLengthFound = ^uint(0)
	}

	// Start with the start node as the current node
	startNode := graphMap.StartNode

	// Set the MinLength found in the node to 0
	startNode.MinLengthFound = 0

	// Make a slice of Node with the length of 0. This is to keep track of nodes to be sorted
	queue := make([]*Node, 0)

	// This holds nodes of paths yet to be sorted
	queue = append(queue, startNode)

	for len(queue) > 0 {
		// Always work with the first node inside the queue.
		node := queue[0]
		// Using the FIFO principle, pop (take out) the node at the index of 0 (The first node). [the node being worked on]
		queue = queue[1:]

		// iterate/loop through the paths connecting the current node [being worked on] to other nodes
		// and carry out the following
		for _, path := range node.Paths {
			if path.ToNode.MinLengthFound > path.Length+node.MinLengthFound {
				path.ToNode.MinLengthFound = path.Length + node.MinLengthFound
				queue = append(queue, path.ToNode)
			}
		}
	}

	TimeDuration := time.Since(startTime)
	fmt.Println("Time Duration", TimeDuration)

	// Once all the paths have been sorted and the minimum length set for all nodes,
	// this returns the minimum length of the End node in the graph
	return graphMap.EndNode.MinLengthFound
}
