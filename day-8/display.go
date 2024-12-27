package main

import (
	"github.com/pterm/pterm"
	"io"
)

func walkerVisualizerWithEmbeddedCache(writer io.Writer) func(w *walker, step int) {
	cachedSteps := make([]int, 0)
	return func(w *walker, step int) {
		if w.isAtFunc(stringEndsInZ) && len(cachedSteps) < 100 {
			cachedSteps = append(cachedSteps, step)
		}
		pterm.Fprinto(writer, pterm.Sprintf("%s\t%s\t%v", w.position, w.firstPos, firstFiveElementsIfPresent(cachedSteps)))
	}
}

func firstFiveElementsIfPresent(arr []int) []int {
	return arr[:arrayDisplayLimit(arr)]
}

func arrayDisplayLimit(arr []int) int {
	limit := 15
	if len(arr) <= limit {
		return len(arr)
	}
	return limit
}

func teamVisualizer(writer ...io.Writer) func(t *team, step int) {
	return func(t *team, step int) {
		pterm.Fprinto(writer[0], pterm.Sprintf("Current\tStart\t%d", step))
		//lcm := 1
		for _, a := range t.as {
			//lcm = Lcm(a.)
			a.visualize(step)
		}
		return
	}
}
