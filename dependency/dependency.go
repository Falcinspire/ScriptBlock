package dependency

type DependencyGraph struct {
	N     int
	Nodes [][]int
}

func NewDependencyGraph(size int) DependencyGraph {
	graph := DependencyGraph{size, make([][]int, size)}
	for i := 0; i < size; i++ {
		graph.Nodes[i] = make([]int, 0)
	}
	return graph
}

func AddDependency(src, depends int, graph DependencyGraph) {
	graph.Nodes[src] = append(graph.Nodes[src], depends)
}

type VisitState int

const (
	UNSEEN    = 0
	RESOLVING = 1
	RESOLVED  = 2
)

func MakeDependencyOrder(graph DependencyGraph) (order []int, circular bool) {
	order = make([]int, 0)
	used := make([]VisitState, graph.N)
	for i := 0; i < graph.N; i++ {
		order, circular = dependencyRecursive(i, order, used, graph)
		if circular {
			order = make([]int, 0)
			return
		}
	}
	return
}

func dependencyRecursive(src int, order []int, used []VisitState, graph DependencyGraph) (orderNew []int, circular bool) {
	if used[src] == RESOLVED {
		return order, false
	}

	used[src] = RESOLVING
	for _, dependsOn := range graph.Nodes[src] {
		if used[dependsOn] == UNSEEN {
			order, circular = dependencyRecursive(dependsOn, order, used, graph)
			if circular {
				order = make([]int, 0)
				return
			}
		} else if used[dependsOn] == RESOLVING {
			order = make([]int, 0)
			circular = true
			return
		}
	}
	order = append(order, src)
	used[src] = RESOLVED
	return order, false
}
