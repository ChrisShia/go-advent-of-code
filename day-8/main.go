package main

import (
	"go-advent-of-code/utils"
	"go-advent-of-code/utils/maths"
	"strings"
)

const inputPath_ = "input/day-8.txt"

//	part b : 22103062509257

var leftRightTurns_ []int

func main() {
	file := utils.OpenFileLogFatal(inputPath_)
	defer utils.CloseFile(file)
	leftTurnOperator, rightTurnOperator := createLeftRightOperators(file)
	applyLeftRightTurnsOnStartingState(leftTurnOperator, rightTurnOperator, leftRightTurns_)
}

func applyLeftRightTurnsOnStartingNode(leftTurnOperator, rightTurnOperator maths.Matrix, startingNodeId string, leftRightTurns []int) (int, adjacency) {
	w := walker{pos: startingNodeId, firstPos: startingNodeId, visualizer: nil}
	return By{leftRightTransformFunc(leftTurnOperator, rightTurnOperator),
		func(nodeId string) bool {
			if nodeId == "ZZZ" {
				return false
			}
			return true
		}, adjacency(&w)}.apply(leftRightTurns)
}

func isStartingNode(s string) bool {
	bytes := []byte(s)
	if len(bytes) == 0 {
		return false
	}
	return byteSliceEndsInA(bytes)
}

func byteSliceEndsInA(bs []byte) bool {
	return 'A' == bs[len(bs)-1]
}

func stringEndsInZ(s string) bool {
	return 'Z' == s[len(s)-1]
}

type By struct {
	turn         func(leftOrRight int, nodeId adjacency)
	keepCounting func(inputId string) bool
	startingFrom adjacency
}

func (by By) apply(path path) (int, adjacency) {
	a := by.startingFrom
	iteration, step := 0, 0
	for !a.isEnd(by.keepCounting) {
		leftOrRight := path.step(iteration)
		by.turn(leftOrRight, a)
		step = iteration + 1
		a.visualize(step)
		iteration++
	}
	return step, a
}

type path []int

func (p path) step(i int) int {
	if i < len(p) {
		return p[i]
	}
	return p[maths.Mod(i, len(p))]
}

type adjacency interface {
	progress(m maths.Matrix)
	string() string
	isEnd(func(position string) bool) bool
	isAt(func(position string) bool) bool
	visualize(step int)
	cachedPositions(upToStep int) []int
}

type walker struct {
	pos            string
	firstPos       string
	visualizer     func(w *walker, step int)
	stepCache      []int
	cachePredicate func(pos string) bool
}

func newW(startPos string) walker {
	return newWalker(startPos, nil)
}

func newWalker(startPos string, vis func(w *walker, step int)) walker {
	return walker{pos: startPos, firstPos: startPos, visualizer: vis}
}

type team struct {
	as         []adjacency
	visualizer func(t *team, step int)
	cache      []int
}

func newT(as []adjacency) team {
	return newTeam(as, nil)
}

func newTeam(as []adjacency, vis func(t *team, step int)) team {
	return team{as: as, visualizer: vis}
}

func (w *walker) cachedPositions(upTo int) []int {
	cache := make([]int, 0)
	cache = w.stepCache
	if upTo >= cache[len(cache)-1] {
		return cache
	}
	for i, fromCache := range cache {
		if fromCache == upTo {
			return cache[:i]
		} else if fromCache > upTo {
			return cache[:i-1]
		}
	}
	return w.stepCache
}

func (t *team) cachedPositions(upTo int) []int {
	//return lastCachedMultipletUpTo(upTo)
	return nil
}

//func (t *team) lastCachedMultipletUpTo(upTo int) []int {
//	for i, a := range t.as {
//		aCache := a.cachedPositions(upTo)
//		aCache
//	}
//	return t.cache
//}

func (w *walker) updateCache(step int) {
	if w.cachePredicate(w.pos) || len(w.stepCache) <= 100 {
		w.stepCache = append(w.stepCache, step)
		return
	}
	return
}

func (w *walker) visualize(step int) {
	if w.visualizer == nil {
		return
	}
	w.visualizer(w, step)
	return
}

func (t *team) visualize(step int) {
	if t.visualizer == nil {
		return
	}
	t.visualizer(t, step)
	return
}

func (w *walker) isAt(f func(position string) bool) bool {
	return f(w.pos)
}

func (w *walker) isEnd(keepCountingFunc func(id string) bool) bool {
	if keepCountingFunc(w.pos) {
		return false
	}
	return true
}

func (w *walker) string() string {
	return w.pos
}

func (t *team) isAt(f func(position string) bool) bool {
	for _, a := range t.as {
		if a.isAt(f) {
			return true
		}
	}
	return false
}

func (t *team) isEnd(keepCountingFunc func(id string) bool) bool {
	keepCounting := false
	isEnd := true
	for _, a := range t.as {
		if keepCountingFunc(a.string()) {
			return keepCounting
		}
	}
	return isEnd
}

func (t *team) string() string {
	stringBuilder := strings.Builder{}
	for _, a := range t.as {
		stringBuilder.WriteString(a.string())
		stringBuilder.WriteByte(' ')
	}
	return stringBuilder.String()
}

func (t *team) progress(m maths.Matrix) {
	for index := range t.as {
		(t.as)[index].progress(m)
	}
	return
}

func (w *walker) progress(m maths.Matrix) {
	w.transformBy(m)
	return
}

func (w *walker) transformBy(m maths.Matrix) {
	adjacencyOfNode := m.AdjacencyMap[w.pos]
	w.pos = (*m.OrderedKeys)[adjacencyOfNode-1]
}
