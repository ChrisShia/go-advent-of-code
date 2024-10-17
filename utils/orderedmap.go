package utils

type OrderedMap[K comparable, V any] struct {
	M           map[K]V
	OrderedKeys *[]K
	OrderFunc   func(K, V)
}

func NewOrderedMap[K comparable, V any](orderedKeys *[]K) *OrderedMap[K, V] {
	return &OrderedMap[K, V]{M: make(map[K]V), OrderedKeys: orderedKeys}
}

func (om *OrderedMap[K, V]) Put(key K, v V) {
	//if key == nil || v == nil {
	//	return
	//}
	om.addKeyToOrderIfAbsent(key)
	//if adjNodeId != nodeId && adjNodeOrder == -1 {
	//	*om.OrderedKeys = append(*om.OrderedKeys, adjNodeId)
	//}
	om.M[key] = v
}

func (om *OrderedMap[K, V]) addKeyToOrderIfAbsent(key K) {
	keyOrder := om.getOrder(key)
	//adjNodeOrder := om.getOrder(adjNodeId)
	if keyOrder == -1 {
		*om.OrderedKeys = append(*om.OrderedKeys, key)
	}
}

func (om *OrderedMap[K, V]) getOrder(key K) int {
	for index, k := range *om.OrderedKeys {
		if k == key {
			return index + 1
		}
	}
	return -1
}
