package main

import (
	"reflect"
	"testing"
)

const HUNDRED int = 100

func TestHand_Type(t *testing.T) {
	var tests = []struct {
		name string
		hand hand
		want int
	}{
		{"test HIGH_CARD", createHand([]Card{Card('J'), Card('A'), Card('4'), Card('3'), Card('5')}), HIGH_CARD},
		{"test ONE_PAIR", createHand([]Card{Card('3'), Card('5'), Card('5'), Card('6'), Card('J')}), ONE_PAIR},
		{"test TWO_PAIR", createHand([]Card{Card('A'), Card('A'), Card('J'), Card('J'), Card('5')}), TWO_PAIR},
		{"test THREE_OF_A_KIND", createHand([]Card{Card('9'), Card('8'), Card('4'), Card('4'), Card('4')}), THREE_OF_A_KIND},
		{"test FULL_HOUSE", createHand([]Card{Card('9'), Card('9'), Card('4'), Card('4'), Card('4')}), FULL_HOUSE},
		{"test FOUR_OF_A_KIND", createHand([]Card{Card('9'), Card('J'), Card('J'), Card('J'), Card('J')}), FOUR_OF_A_KIND},
		{"test FIVE_OF_A_KIND", createHand([]Card{Card('J'), Card('J'), Card('J'), Card('J'), Card('J')}), FIVE_OF_A_KIND},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.hand.Type()
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestHand_TypeWithJoker(t *testing.T) {
	var tests = []struct {
		name string
		hand hand
		want hand
	}{

		//{"test HIGH_CARD", createHand([]Card{Card('J'), Card('4'), Card('A'), Card('3'), Card('5')}), },
		//{"test ONE_PAIR", createHand([]Card{Card('3'), Card('5'), Card('5'), Card('6'), Card('J')}, 100), ONE_PAIR},
		//{"test TWO_PAIR", createHand([]Card{Card('A'), Card('A'), Card('J'), Card('J'), Card('5')}, 100), TWO_PAIR},
		//{"test FOUR_OF_A_KIND", createHand([]Card{Card('9'), Card('J'), Card('J'), Card('J'), Card('J')}, 100), FOUR_OF_A_KIND},
		//{"test FIVE_OF_A_KIND", createHand([]Card{Card('J'), Card('J'), Card('J'), Card('J'), Card('J')}, 100), FIVE_OF_A_KIND},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.hand.Type()
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func createHand(cards []Card) hand {
	occurrenceMap := make(map[int]int)
	for _, c := range cards {
		occurrenceMap[c.Strength()]++
	}
	return hand{cards, HUNDRED, occurrenceMap}
}
