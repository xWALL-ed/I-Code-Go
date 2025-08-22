package techpalace

import "strings"

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, "+strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    var buildStr strings.Builder
    buildStr.WriteString(strings.Repeat("*",numStarsPerLine))
    buildStr.WriteString("\n")
    buildStr.WriteString(welcomeMsg)
    buildStr.WriteString("\n")
    buildStr.WriteString(strings.Repeat("*",numStarsPerLine))
	return buildStr.String()
}

func CleanupMessage(oldMsg string) string {
	return strings.Trim(oldMsg, "* \n")
}
