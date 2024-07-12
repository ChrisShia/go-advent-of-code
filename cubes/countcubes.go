package main

import (
	"bytes"
	"fmt"
	"go-advent-of-code/dictionary"
	"go-advent-of-code/utils"
	"unicode"
)

func main() {
	input := utils.ReadFile("/Users/christos/Practise-code/Go/go-advent/cubes/configurations")
	//idSum := applyGameProcessor(input, summer())
	minimumSetSum := applyGameProcessor(input, minimumSetPowerSummer())
	fmt.Printf("The sum of ids is : %d", minimumSetSum)
}

func applyGameProcessor(gameSequence []byte, gameProcessor func(game []byte) int) int {
	var sum int
	var found = true
	var game []byte
	var remainingSeq = gameSequence
	for found {
		game, remainingSeq, found = bytes.Cut(remainingSeq, []byte{'\n'})
		sum = gameProcessor(game)
	}
	return sum
}

func minimumSetPowerSummer() func(game []byte) int {
	minimumSetPowerSum := 0
	return func(game []byte) int {
		var power = PowerOfMinSet(game)
		minimumSetPowerSum += power
		return minimumSetPowerSum
	}
}

func PowerOfMinSet(game []byte) int {
	var colorScoreMap = map[byte]int{
		'r': 0,
		'g': 0,
		'b': 0,
	}
	_, subSets, _ := bytes.Cut(game, []byte{':'})
	var foundSubSet = true
	var subSet []byte
	var remainingSubSets = subSets
	for foundSubSet {
		subSet, remainingSubSets, foundSubSet = bytes.Cut(remainingSubSets, []byte{';'})
		red, green, blue := parseSubSet(subSet)
		if colorScoreMap['r'] < red {
			colorScoreMap['r'] = red
		}
		if colorScoreMap['g'] < green {
			colorScoreMap['g'] = green
		}
		if colorScoreMap['b'] < blue {
			colorScoreMap['b'] = blue
		}
	}
	return colorScoreMap['r'] * colorScoreMap['g'] * colorScoreMap['b']
}

func possibleGameSummer() func(game []byte) int {
	eligibleGamesIdSum := 0
	return func(game []byte) int {
		isPossible, id := assessEligibilityOfGame(game)
		if isPossible {
			eligibleGamesIdSum += id
		}
		return eligibleGamesIdSum
	}
}

func assessEligibilityOfGame(game []byte) (bool, int) {
	before, after, _ := bytes.Cut(game, []byte{':'})
	return isGamePossible(after), extractGameId(before)
}

func isGamePossible(subSets []byte) bool {
	var foundSubSet = true
	var subSet []byte
	var remainingSubSets = subSets
	for foundSubSet {
		subSet, remainingSubSets, foundSubSet = bytes.Cut(remainingSubSets, []byte{';'})
		if isSubSetImpossible(subSet) {
			return false
		}
	}
	return true
}

func isSubSetImpossible(subSet []byte) bool {
	var red, green, blue int = parseSubSet(subSet)
	return red > 12 || green > 13 || blue > 14
}

func parseSubSet(subSet []byte) (int, int, int) {
	var colorScoreMap = map[byte]int{
		'r': 0,
		'g': 0,
		'b': 0,
	}
	var foundColor = true
	var colorScore []byte
	var remainingColorScore = subSet
	for foundColor {
		colorScore, remainingColorScore, foundColor = bytes.Cut(remainingColorScore, []byte{','})
		colorScore = bytes.TrimSpace(colorScore)
		score, color, _ := bytes.Cut(colorScore, []byte(" "))
		colorScoreMap[color[0]] = dictionary.BytesToInt(score)
	}
	return colorScoreMap['r'], colorScoreMap['g'], colorScoreMap['b']
}

func extractGameId(bs []byte) int {
	id := bytes.TrimFunc(bs, isNotNumber)
	return dictionary.BytesToInt(id)
}

func isNotNumber(r rune) bool {
	return !unicode.IsNumber(r)
}
