package main

import (
	"reflect"
	"testing"
)

const HUNDRED int = 100

func TestHand_Type(t *testing.T) {
	normalRules := createNormalGameRules()
	var tests = []struct {
		name string
		hand *hand
		want int
	}{
		{"test HIGH_CARD", createHand([]Card{Card('J'), Card('A'), Card('4'), Card('3'), Card('5')}), HighCard},
		{"test ONE_PAIR", createHand([]Card{Card('3'), Card('5'), Card('5'), Card('6'), Card('J')}), OnePair},
		{"test TWO_PAIR", createHand([]Card{Card('A'), Card('A'), Card('J'), Card('J'), Card('5')}), TwoPair},
		{"test THREE_OF_A_KIND", createHand([]Card{Card('9'), Card('8'), Card('4'), Card('4'), Card('4')}), ThreeOfAKind},
		{"test FULL_HOUSE", createHand([]Card{Card('9'), Card('9'), Card('4'), Card('4'), Card('4')}), FullHouse},
		{"test FOUR_OF_A_KIND", createHand([]Card{Card('9'), Card('J'), Card('J'), Card('J'), Card('J')}), FourOfAKind},
		{"test FIVE_OF_A_KIND", createHand([]Card{Card('J'), Card('J'), Card('J'), Card('J'), Card('J')}), FiveOfAKind},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := normalRules.typeRule.Type(tt.hand)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestHand_TypeWithJoker(t *testing.T) {
	gameRules := createGameRulesWithTrumpCardJ()
	var tests = []struct {
		name string
		hand *hand
		want int
	}{
		{"test HIGH_CARD", createHand([]Card{Card('J'), Card('4'), Card('A'), Card('3'), Card('5')}), OnePair},
		{"test ONE_PAIR", createHand([]Card{Card('3'), Card('5'), Card('5'), Card('6'), Card('J')}), ThreeOfAKind},
		{"test TWO_PAIR", createHand([]Card{Card('A'), Card('A'), Card('J'), Card('J'), Card('5')}), FourOfAKind},
		{"test FOUR_OF_A_KIND", createHand([]Card{Card('9'), Card('J'), Card('J'), Card('J'), Card('J')}), FiveOfAKind},
		{"test FIVE_OF_A_KIND", createHand([]Card{Card('J'), Card('J'), Card('J'), Card('J'), Card('J')}), FiveOfAKind},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := gameRules.typeRule.Type(tt.hand)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestGameRule_Compare(t *testing.T) {
	gameRules := createGameRulesWithTrumpCardJ()
	var tests = []struct {
		name  string
		hand1 *hand
		hand2 *hand
		want  int
	}{
		{"test compare QQQJA with KTJJT",
			createHand([]Card{Card('Q'), Card('Q'), Card('Q'), Card('J'), Card('A')}),
			createHand([]Card{Card('K'), Card('T'), Card('J'), Card('J'), Card('T')}), -1},
		{"test compare KTJJT with QQQJA",
			createHand([]Card{Card('K'), Card('T'), Card('J'), Card('J'), Card('T')}),
			createHand([]Card{Card('Q'), Card('Q'), Card('Q'), Card('J'), Card('A')}), 1},
		{"test compare JJJJJ with QQQJJ",
			createHand([]Card{Card('J'), Card('J'), Card('J'), Card('J'), Card('J')}),
			createHand([]Card{Card('Q'), Card('Q'), Card('Q'), Card('J'), Card('J')}), -1},
		{"test compare QJJJJ with QQQJJ",
			createHand([]Card{Card('Q'), Card('J'), Card('J'), Card('J'), Card('J')}),
			createHand([]Card{Card('Q'), Card('Q'), Card('Q'), Card('J'), Card('J')}), -1},
		{"test compare KJJJJ with QQQJJ",
			createHand([]Card{Card('K'), Card('J'), Card('J'), Card('J'), Card('J')}),
			createHand([]Card{Card('Q'), Card('Q'), Card('Q'), Card('J'), Card('J')}), 1},
		{"test compare KK677 with QQQJJ",
			createHand([]Card{Card('K'), Card('J'), Card('J'), Card('J'), Card('J')}),
			createHand([]Card{Card('Q'), Card('Q'), Card('Q'), Card('J'), Card('J')}), 1},
		{"test compare KK677 with T55J5",
			createHand([]Card{Card('K'), Card('K'), Card('6'), Card('7'), Card('7')}),
			createHand([]Card{Card('T'), Card('5'), Card('5'), Card('J'), Card('5')}), -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := gameRules.compare(tt.hand1, tt.hand2)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func createNormalGameRules() *GameRules {
	typeRule := TypeRule(func(h *hand) map[Card]int { return h.cardFrequencyMap })
	strengthRule := StrengthRule(func(c *Card) int { return c.Strength() })
	return &GameRules{typeRule, strengthRule}
}

func createGameRulesWithTrumpCardJ() *GameRules {
	typeRule := TypeRule(func(h *hand) map[Card]int { return applyTrumpCard(h.cardFrequencyMap, Card('J')) })
	strengthRule := StrengthRule(func(c *Card) int { return c.StrengthJWeakest() })
	return &GameRules{typeRule, strengthRule}
}

func createHand(cards []Card) *hand {
	occurrenceMap := make(map[Card]int)
	for _, c := range cards {
		occurrenceMap[c]++
	}
	return &hand{cards, HUNDRED, occurrenceMap}
}
