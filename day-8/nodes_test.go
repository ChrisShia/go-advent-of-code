package main

import (
	"fmt"
	"testing"
)

func Test_Multiply(t *testing.T) {
	left, right := createAdjacencyMatrices()
	leftRight := left.multiply(right)
	rightLeft := right.multiply(left)
	fmt.Println(leftRight.adjacencyMap)
	fmt.Println(rightLeft.adjacencyMap)
}

func Test_ProgressWalker(t *testing.T) {
	left, _ := createAdjacencyMatrices()
	var tests = []struct {
		name                string
		nodeToBeTransformed walker
		operator            Matrix
		want                walker
	}{
		{"turn walker AAA by left op", newW("AAA"), left, newW("BBB")},
		{"turn walker BBB by left op", newW("BBB"), left, newW("AAA")},
		{"turn walker ZZZ by left op", newW("ZZZ"), left, newW("ZZZ")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.operator.transform(&tt.nodeToBeTransformed)
			if ans.string() != tt.want.string() {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func Test_ApplyLeftRightTurnsOnStartingNode(t *testing.T) {
	left, right := createAdjacencyMatrices()
	var tests = []struct {
		name                         string
		startNodeId                  string
		leftOrRightBinRepresentation []int
		want                         walker
	}{
		{"apply LR on AAA", "AAA", []int{0, 1}, newW("ZZZ")},
		{"apply LLR on AAA", "AAA", []int{0, 0, 1}, newW("BBB")},
		{"apply LLRL on AAA", "AAA", []int{0, 0, 1, 0}, newW("AAA")},
		{"apply LLRLL on AAA", "AAA", []int{0, 0, 1, 0, 0}, newW("BBB")},
		{"apply LLRLLR on AAA", "AAA", []int{0, 0, 1, 0, 0, 1}, newW("ZZZ")},
		{"apply LLRLLRRR on AAA", "AAA", []int{0, 0, 1, 0, 0, 1, 1, 1}, newW("ZZZ")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ans := applyLeftRightTurnsOnStartingNode(left, right, tt.startNodeId, tt.leftOrRightBinRepresentation)
			if ans.string() != tt.want.string() {
				t.Errorf("got %v, want %v", ans.string(), tt.want)
			}
		})
	}
}

func Test_Count(t *testing.T) {
	left, right := createAdjacencyMatrices()
	var tests = []struct {
		name                         string
		startNodeId                  string
		leftOrRightBinRepresentation []int
		want                         walker
	}{
		{"apply LR on AAA", "AAA", []int{0, 1}, newW("ZZZ")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ans := applyLeftRightTurnsOnStartingNode(left, right, tt.startNodeId, tt.leftOrRightBinRepresentation)
			if ans.string() != tt.want.string() {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func Test_ContainsFunc_ContainsNodeEndingInZ(t *testing.T) {
	tm := createNodeState()
	var tests = []struct {
		name  string
		state adjacency
		want  bool
	}{
		{"apply contains walker that ends with Z", adjacency(&tm), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.state.isAt(stringEndsInZ)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

//func Test_TeamUpdateCachePositions(t *testing.T) {
//	tm := walkersWithCachedPositions()
//	var tests = []struct {
//		name string
//		t    team
//		want []int
//	}{
//		{"", walkersWithCachedPositions(), []int{2, 4, 6}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ans := tt.t.cache
//			if !equal(ans, tt.want) {
//				t.Errorf("got %v, want %v", ans, tt.want)
//			}
//		})
//	}
//}

func equal(t1 []int, t2 []int) bool {
	for i := range t1 {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}

func walkersWithCachedPositions() team {
	as := make([]adjacency, 0)
	VBN, QGZ, KJZ := newWalkerAndCache("VBN", []int{1, 2}), newWalkerAndCache("QGZ", []int{3, 4}), newWalkerAndCache("KJZ", []int{5, 6})
	as = append(as, adjacency(&VBN))
	as = append(as, adjacency(&QGZ))
	as = append(as, adjacency(&KJZ))
	t := team{as, nil, nil}
	return t
}

func newWalkerAndCache(pos string, cached []int) walker {
	return walker{pos: pos, stepCache: cached}
}

func adjacencySet() []adjacency {
	VBN, QGZ, KJZ, LPK, MHB, JHD := newW("VBN"), newW("QGZ"), newW("KJZ"), newW("LPK"), newW("MHB"), newW("JHD")
	adjacencySlice := make([]adjacency, 0)
	adjacencySlice = append(adjacencySlice, adjacency(createCustomNodeState(VBN, QGZ, KJZ, LPK, MHB)))
	adjacencySlice = append(adjacencySlice, adjacency(createCustomNodeState(VBN, QGZ)))
	adjacencySlice = append(adjacencySlice, adjacency(createCustomNodeState(LPK, MHB, JHD)))
	adjacencySlice = append(adjacencySlice, adjacency(&MHB))
	adjacencySlice = append(adjacencySlice, adjacency(createCustomNodeState(MHB, LPK)))
	adjacencySlice = append(adjacencySlice, adjacency(&LPK))
	adjacencySlice = append(adjacencySlice, adjacency(createCustomNodeState(QGZ, KJZ, MHB)))
	adjacencySlice = append(adjacencySlice, adjacency(createCustomNodeState(LPK, VBN)))
	adjacencySlice = append(adjacencySlice, adjacency(&QGZ))
	adjacencySlice = append(adjacencySlice, adjacency(createCustomNodeState(MHB, JHD)))
	return adjacencySlice
}

func createCustomNodeState(walkers ...walker) *team {
	adjacencySlice := make([]adjacency, 0)
	t := newT(adjacencySlice)
	if walkers == nil || len(walkers) == 0 {
		return &t
	}
	for _, w := range walkers {
		adjacencySlice = append(adjacencySlice, adjacency(&w))
	}
	return &t
}

func createNodeState() team {
	AFG, BBZ, BBB, HGF, JIK := "AFG", "BBZ", "BBB", "HGF", "JIK"
	walkers := []walker{newW(AFG), newW(BBZ), newW(BBB), newW(HGF), newW(JIK)}
	adjSlice := make([]adjacency, 0)
	for _, w := range walkers {
		newW := new(walker)
		*newW = w
		adjSlice = append(adjSlice, adjacency(newW))
	}
	return newT(adjSlice)
}

func createAdjacencyMatrices() (Matrix, Matrix) {
	input := [][]string{{"AAA", "BBB", "BBB"}, {"BBB", "AAA", "ZZZ"}, {"ZZZ", "ZZZ", "ZZZ"}}
	orderedKeys := make([]string, 0)
	leftOrderedMap := newOrderedMap(&orderedKeys)
	rightOrderedMap := newOrderedMap(&orderedKeys)
	nodeSetter := nodesFromInputSetter[string](leftOrderedMap, rightOrderedMap, func(input string) string { return input })
	for _, adjacency := range input {
		nodeSetter(adjacency)
	}
	leftMatrix := Matrix{leftOrderedMap}
	rightMatrix := Matrix{rightOrderedMap}
	return leftMatrix, rightMatrix
}
