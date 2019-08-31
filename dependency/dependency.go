package dependency

import (
	"errors"
)

type UnitNode struct {
	Location     string
	Dependencies []*UnitNode
}

type DependencyGraph struct {
	Nodes map[string]*UnitNode
}

func NewDependencyGraph() *DependencyGraph {
	return &DependencyGraph{make(map[string]*UnitNode)}
}

func InsertNode(location string, graph *DependencyGraph) {
	node := &UnitNode{location, make([]*UnitNode, 0)}
	graph.Nodes[location] = node
}
func AddDependency(unit, dependency string, graph *DependencyGraph) {
	unitNode := graph.Nodes[unit]
	dependencyNode := graph.Nodes[dependency]
	unitNode.Dependencies = append(unitNode.Dependencies, dependencyNode)
}
func GetNode(name string, graph *DependencyGraph) *UnitNode {
	return graph.Nodes[name]
}

type VisitState int

const (
	RESOLVING = 0
	RESOLVED  = 1
)

func MakeDependencyOrder(root string, graph *DependencyGraph) []string {
	rootNode := GetNode(root, graph)
	order := dependencyRecursive(rootNode, []string{}, make(map[string]VisitState), graph)
	return order
}

func dependencyRecursive(node *UnitNode, order []string, used map[string]VisitState, graph *DependencyGraph) (orderNew []string) {
	used[node.Location] = RESOLVING
	for _, dependsOn := range node.Dependencies {
		state, exists := used[dependsOn.Location]
		if !exists {
			order = dependencyRecursive(dependsOn, order, used, graph)
		} else if state == RESOLVING {
			panic(errors.New("Circular Reference"))
		}
	}
	order = append(order, node.Location)
	used[node.Location] = RESOLVED
	return order
}
