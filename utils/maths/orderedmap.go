package maths

type OrderedMap struct {
	AdjacencyMap map[string]int
	OrderedKeys  *[]string
}

func NewOrderedMap(orderedKeys *[]string) *OrderedMap {
	return &OrderedMap{make(map[string]int), orderedKeys}
}

func (om *OrderedMap) getOrder(key string) int {
	for k, v := range *om.OrderedKeys {
		if v == key {
			return k + 1
		}
	}
	return -1
}

func (om *OrderedMap) AddSingleAdjacencyForNode(nodeId string, adjNodeId string) {
	if len(nodeId) == 0 || len(adjNodeId) == 0 {
		return
	}
	nodeOrder := om.getOrder(nodeId)
	adjNodeOrder := om.getOrder(adjNodeId)
	if nodeOrder == -1 {
		*om.OrderedKeys = append(*om.OrderedKeys, nodeId)
	}
	if adjNodeId != nodeId && adjNodeOrder == -1 {
		*om.OrderedKeys = append(*om.OrderedKeys, adjNodeId)
	}
	om.AdjacencyMap[nodeId] = om.getOrder(adjNodeId)
}
