package main

import (
	"fmt"
	"gitlab.com/pulsar-ai-public/technical-tests/golang-developer/graph"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: go run graph/graphmain/main.go <path to map file>")
		return
	}
	mapFile := args[1]
	graphMap := graph.ParseFile(mapFile)
	minPath := graph.FindShortestPath(graphMap)
	fmt.Println(minPath)
}
