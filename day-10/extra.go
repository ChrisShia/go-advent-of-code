package main

import maths "github.com/ChrisShia/math-depot"

//func (c *circuit) withinBoundaries(gridPoint int) bool {
//	if c.contains(gridPoint) {
//		return false
//	}
//	northBoundaryPoint := gridPoint
//	southBoundaryPoint := gridPoint
//	eastBoundaryPoint := gridPoint
//	westBoundaryPoint := gridPoint
//	for !c.contains(northBoundaryPoint) {
//		northBoundaryPoint = move(northBoundaryPoint, north)
//		if northBoundaryPoint == indexOutOfGridBounds {
//			break
//		}
//	}
//	for !c.contains(southBoundaryPoint) {
//		southBoundaryPoint = move(southBoundaryPoint, south)
//		if southBoundaryPoint == indexOutOfGridBounds {
//			break
//		}
//	}
//	for !c.contains(eastBoundaryPoint) {
//		eastBoundaryPoint = move(eastBoundaryPoint, east)
//		if eastBoundaryPoint == indexOutOfGridBounds {
//			break
//		}
//	}
//	for !c.contains(westBoundaryPoint) {
//		westBoundaryPoint = move(westBoundaryPoint, west)
//		if westBoundaryPoint == indexOutOfGridBounds {
//			break
//		}
//	}
//	return
//}

func move(fromGridPoint int, toDirection direction) int {
	return pipe(fromGridPoint, toDirection)
}

func findInnerNodeAlongLine(circuitNodes []int) (int, bool) {
	found := true
	for i := 0; i < len(circuitNodes); i += 2 {
		eastAdjacentGridIndex := circuitNodes[i] + 1
		if eastAdjacentGridIndex == circuitNodes[i+1] {
			continue
		} else {
			return eastAdjacentGridIndex, found
		}
	}
	return -1, !found
}

func eastMostIndexOnLongitude(indexOnGrid int) int {
	return gridWidth_*(indexOnGrid/gridWidth_) + gridWidth_ - 1
}

func westMostIndexOnLongitude(indexOnGrid int) int {
	indexModWidth := maths.Mod(indexOnGrid+1, gridWidth_)
	return indexOnGrid - indexModWidth
}
