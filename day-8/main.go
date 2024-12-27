package main

import (
	maths "github.com/ChrisShia/math-depot"
	"go-advent-of-code/utils"
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
	w := walker{position: startingNodeId, firstPos: startingNodeId, visualizer: nil}
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
	turn           func(leftOrRight int, nodeId adjacency)
	isEndPredicate func(inputId string) bool
	startingFrom   adjacency
}

func (by By) apply(path path) (int, adjacency) {
	a := by.startingFrom
	iteration, step := 0, 0
	for !a.isEnd(by.isEndPredicate) {
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
	isAt(f func(position string) bool) bool
	visualize(step int)
}

type walker struct {
	position       string
	firstPos       string
	visualizer     func(w *walker, step int)
	stepCache      []int
	cachePredicate func(pos string) bool
}

type team struct {
	as         []adjacency
	visualizer func(t *team, step int)
	cache      []int
}

func newTeam(as []adjacency, vis func(t *team, step int)) team {
	return team{as: as, visualizer: vis}
}

func (w *walker) updateCache(step int) {
	if w.cachePredicate(w.position) || len(w.stepCache) <= 100 {
		w.stepCache = append(w.stepCache, step)
		return
	}
	return
}

func (w *walker) isAtFunc(f func(position string) bool) bool {
	return f(w.position)
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

func (t *team) isAt(f func(position string) bool) bool {
	for _, a := range t.as {
		if a.isAt(f) {
			return true
		}
	}
	return false
}

func (w *walker) isAt(f func(position string) bool) bool {
	return f(w.position)
}

func (w *walker) isEnd(isEndPredicateFunc func(id string) bool) bool {
	return isEndPredicateFunc(w.position)
}

func (w *walker) string() string {
	return w.position
}

func (t *team) isEnd(isEndPredicateFunc func(id string) bool) bool {
	keepCounting := false
	isEnd := true
	for _, a := range t.as {
		if !isEndPredicateFunc(a.string()) {
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
	adjacencyOfNode := m.AdjacencyMap[w.position]
	w.position = (*m.OrderedKeys)[adjacencyOfNode-1]
}
