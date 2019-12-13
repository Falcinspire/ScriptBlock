package dependency

import (
	"fmt"
	"testing"
)

func TestDependency(t *testing.T) {
	{
		graph := NewDependencyGraph()
		for i := 0; i < 6; i++ {
			AddVertex(graph)
		}
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
	}

	{
		graph := NewDependencyGraph()
		for i := 0; i < 3; i++ {
			AddVertex(graph)
		}
		AddDependency(0, 1, graph)
		AddDependency(1, 2, graph)
		AddDependency(2, 0, graph)
		order, circular := MakeDependencyOrder(graph)
		t.Log(order)
		if circular != true {
			panic(fmt.Errorf("second graph is circular"))
		}
	}
}
