package techpalace


import (
	"fmt"
    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	message := fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
    
	return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	borderChars := strings.Repeat("*", numStarsPerLine)
    result := borderChars + "\n" + welcomeMsg + "\n" + borderChars
    
    return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	text := strings.ReplaceAll(oldMsg, "*", "")
    text = strings.TrimSpace(text)
    
    return text
}
