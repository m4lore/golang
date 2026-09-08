package techpalace

import "strings"

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    s := strings.Repeat("*", numStarsPerLine)
	return s + "\n" + welcomeMsg + "\n" + s
}

func CleanupMessage(oldMsg string) string {
	msg := strings.ReplaceAll(oldMsg, "*", "")
    return strings.TrimSpace(msg)
}
