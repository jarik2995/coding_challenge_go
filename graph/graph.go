package graph

type Graph struct {
	NodesCount   int
	AdjacentList map[interface{}][]interface{}
}

// AddNode
func (graph *Graph) AddNode(v interface{}) {
	if graph.AdjacentList == nil {
		graph.AdjacentList = make(map[interface{}][]interface{})
	}
	graph.AdjacentList[v] = []interface{}{}
	graph.NodesCount++
}

// AddEdge
func (graph *Graph) AddEdge(n1 interface{}, n2 interface{}) {
	_, exist := graph.AdjacentList[n1]
	if exist {
		return
	}
	_, exist = graph.AdjacentList[n2]
	if exist {
		return
	}
	graph.AdjacentList[n1] = append(graph.AdjacentList[n1], n2)
	graph.AdjacentList[n2] = append(graph.AdjacentList[n2], n1)
}
