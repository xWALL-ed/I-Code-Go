package techpalace

import "strings"

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, "+strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*",numStarsPerLine)+"\n"+welcomeMsg+"\n"+strings.Repeat("*",numStarsPerLine)
}

func CleanupMessage(oldMsg string) string {
	return strings.Trim(oldMsg, "* \n")
}
