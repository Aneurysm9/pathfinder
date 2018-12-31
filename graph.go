// Package pathfinder seeks to provide a small set of useful pathfinding algorithms with a simple, easy-to-use interface.
package pathfinder

// Node represents a node in the graph that can be traversed.
// It is important that types that implement the Node interface are comparable. They may be used as map keys.
type Node interface{}

// Edge represents a connection from one Node to another.
type Edge struct {
	Target    Node
	Cost      int
	Heuristic int
}

// Path represents a path between two Nodes.
type Path []Node

// Graph represents a set of connected nodes.
type Graph interface {
	Neighbors(Node) []Edge
}

// Finder holds a Graph and presents pathfinding operations on it.
type Finder struct {
	g Graph
}

// NewFinder returns a new Finder wrapping your Graph
func NewFinder(g Graph) *Finder {
	return &Finder{g}
}
