package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	message := fmt.Sprintf("Welcome to my party, %s!", name)
    return message
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	message := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
    return message
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	welcomeMessage := fmt.Sprintf("Welcome to my party, %s!", name)
    tableInfo := fmt.Sprintf("You have been assigned to table %03d.", table)
    directionMessage := fmt.Sprintf("Your table is %s, exactly %.1f meters from here.", direction, distance)
    neighborInfo := fmt.Sprintf("You will be sitting next to %s.", neighbor)

    message := welcomeMessage
    message += "\n" + tableInfo + " " + directionMessage
    message += "\n" + neighborInfo
    return message
}
