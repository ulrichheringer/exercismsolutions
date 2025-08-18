package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11 // Ace is worth 11 points
	case "king", "queen", "jack", "ten":
		return 10
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
	}
}
func FirstTurn(card1, card2, dealerCard string) string {
	c1 := ParseCard(card1)
	c2 := ParseCard(card2)
	dc := ParseCard(dealerCard)
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
	if c1+c2 == 21 {
		if dc == 10 || dc == 11 {
			return "S"
		}
		return "W"
	}

	if c1+c2 >= 17 && c1+c2 <= 20 {
		return "S"
	}

	// Rule 4: Sum between 12 and 16
	if c1+c2 >= 12 && c1+c2 <= 16 {
		if dc >= 7 {
			return "H"
		}
		return "S"
	}

	// Rule 5: Sum 11 or lower -> Hit
	if c1+c2 <= 11 {
		return "H"
	}

	return "S" // fallback
}
