package dependency

import (
	"fmt"
	"testing"
)

func TestDependency(t *testing.T) {
	graph := NewDependencyGraph(6)
	AddDependency(0, 1, graph)
	AddDependency(1, 2, graph)
	AddDependency(0, 3, graph)
	AddDependency(3, 4, graph)
	AddDependency(4, 5, graph)
	order, circular := MakeDependencyOrder(graph)
	t.Log(order)
	if circular != false {
		panic(fmt.Errorf("first graph is not circular"))
	}

	graph2 := NewDependencyGraph(3)
	AddDependency(0, 1, graph2)
	AddDependency(1, 2, graph2)
	AddDependency(2, 0, graph2)
	order2, circular2 := MakeDependencyOrder(graph2)
	t.Log(order2)
	if circular2 != true {
		panic(fmt.Errorf("second graph is circular"))
	}
}
