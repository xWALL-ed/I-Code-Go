package blackjack

func ParseCard(card string) int {
	switch card {
        case "ace": return 11
        case "two": return 2
        case "three": return 3
        case "four": return 4
        case "five": return 5
        case "six": return 6
        case "eight": return 8
        case "nine": return 9
        case "ten": return 10
        case "jack": return 10
        case "queen": return 10
        case "king": return 10
        case "seven": return 7
        default: return 0
    }
}

func FirstTurn(card1, card2, dealerCard string) string {
	switch{
        case card1 == "ace" && card2 == "ace" : return "P"
        case ParseCard(card1)+ParseCard(card2)==21 : 
        	if dealerCard != "ace" && dealerCard != "ten" && dealerCard != "jack" && dealerCard != "king" && dealerCard != "queen" { return "W" }
        	return "S"
        case ParseCard(card1)+ParseCard(card2)>=17 && ParseCard(card1)+ParseCard(card2)<=20 : return "S"
        case ParseCard(card1)+ParseCard(card2)>=12 && ParseCard(card1)+ParseCard(card2)<=16 :
        	if ParseCard(dealerCard)>=7 { return "H" }
        	return "S"
        case ParseCard(card1)+ParseCard(card2)<=11 : return "H"
        default : return "H"
    }
}
