package main

import (
	"github.com/pterm/pterm"
	"io"
)

func walkerVisualizer(writer io.Writer) func(w *walker, step int) {
	cachedSteps := make([]int, 0)

	return func(w *walker, step int) {
		if w.containsFunc(stringEndsInZ) && len(cachedSteps) < 100 {
			cachedSteps = append(cachedSteps, step)
		}
		pterm.Fprinto(writer, pterm.Sprintf("%s\t%s\t%v", w.pos, w.firstPos, firstFiveElementsIfPresent(cachedSteps)))
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

func teamVisualizer(writer io.Writer) func(t *team, step int) {
	return func(t *team, step int) {
		pterm.Fprinto(writer, pterm.Sprintf("Current\tStart\t%d", step))
		for _, a := range t.as {
			a.visualize(step)
		}
		return
	}
}
