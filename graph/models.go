package graph

type NodeType uint8

const (
	MIDDLE NodeType = 0
	START  NodeType = 1
	END    NodeType = 2
)

type Node struct {
	Paths []*Path
	Type  NodeType
	Index uint

	// This value is intended to be used during graph resolution and is initialized to uint max value
	MinLengthFound uint
}

type Path struct {
	Length   uint
	FromNode *Node
	ToNode   *Node
}

type Map struct {
	StartNode *Node
	EndNode   *Node
	AllNodes  []*Node

	// This value is provided by the map file and may not reflect the reality (but is right for maps we provide)
	ShortestPathLength uint
}
