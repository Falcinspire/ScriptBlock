package dependency

type DependencyGraph struct {
	Nodes [][]int
}

func NewDependencyGraph() *DependencyGraph {
	return &DependencyGraph{make([][]int, 0)}
}

func AddVertex(graph *DependencyGraph) int {
	graph.Nodes = append(graph.Nodes, make([]int, 0))
	return len(graph.Nodes) - 1
}

func AddDependency(src, depends int, graph *DependencyGraph) {
	graph.Nodes[src] = append(graph.Nodes[src], depends)
}

type VisitState int

const (
	UNSEEN    = 0
	RESOLVING = 1
	RESOLVED  = 2
)

func MakeDependencyOrder(graph *DependencyGraph) (order []int, circular bool) {
	order = make([]int, 0)
	used := make([]VisitState, len(graph.Nodes))
	for i := 0; i < len(graph.Nodes); i++ {
		order, circular = dependencyRecursive(i, order, used, graph)
		if circular {
			order = make([]int, 0)
			return
		}
	}
	return
}

func dependencyRecursive(src int, order []int, used []VisitState, graph *DependencyGraph) (orderNew []int, circular bool) {
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
