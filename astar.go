package pathfinder

import (
	"container/heap"
	"fmt"
)

// AStar uses the A* algorithm to find a shortest path between two Nodes.
// It returns a Path from the source to the destination, the cost of that path,
// and an error if it was unable to find a path.
func (f Finder) AStar(from, to Node) (Path, int, error) {
	frontier := make(priorityQueue, 1)
	frontier[0] = &item{node: from, priority: 0, index: 0}
	heap.Init(&frontier)
	cameFrom := map[Node]Node{}
	costSoFar := map[Node]int{}

	cameFrom[from] = from
	costSoFar[from] = 0

	for frontier.Len() > 0 {
		i := heap.Pop(&frontier).(*item)
		cur := i.node

		if cur == to {
			curPath := Path{cur}
			c := 0
			for n := cameFrom[cur]; n != from; n = cameFrom[n] {
				c++
				curPath = append(Path{n}, curPath...)
			}
			curPath = append(Path{from}, curPath...)
			return curPath, costSoFar[cur], nil
		}

		for _, next := range f.g.Neighbors(cur) {
			nextCost := costSoFar[cur] + next.Cost
			if cost, ok := costSoFar[next.Target]; !ok || nextCost < cost {
				costSoFar[next.Target] = nextCost
				heap.Push(&frontier, &item{node: next.Target, priority: nextCost + next.Heuristic})
				cameFrom[next.Target] = cur
			}
		}
	}

	return Path{}, -1, fmt.Errorf("Unable to find path")
}
