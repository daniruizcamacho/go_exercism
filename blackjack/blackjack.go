package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "ten" || card == "king" || card == "queen" || card == "jack":
		return 10
	case card == "nine":
		return 9
	case card == "eight":
		return 8
	case card == "seven":
		return 7
	case card == "six":
		return 6
	case card == "five":
		return 5
	case card == "four":
		return 4
	case card == "three":
		return 3
	case card == "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Parsed := ParseCard(card1)
	card2Parsed := ParseCard(card2)
	dealerCardParsed := ParseCard(dealerCard)
	sumCards := card1Parsed + card2Parsed

	switch {
	case card1Parsed == 11 && card2Parsed == 11:
		return "P"
	case sumCards == 21 && dealerCardParsed < 10:
		return "W"
	case sumCards == 21 && dealerCardParsed >= 10:
		return "S"
	case sumCards >= 17 && sumCards <= 20:
		return "S"
	case sumCards >= 12 && sumCards <= 16 && dealerCardParsed < 7:
		return "S"
	case sumCards >= 12 && sumCards <= 16 && dealerCardParsed >= 7:
		return "H"
	default:
		return "H"
	}
}
