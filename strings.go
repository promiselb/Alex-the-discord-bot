package discordgo

import (
	"fmt"
	"math/rand"
	"strings"
)

// help is used to get all commands and their descriptions when $help is called.
func help() string {
	var s = "Command: Description\n\n"

	for _, v := range commands {
		s += fmt.Sprintf("%s: %s\n\n", v.name, v.description)
	}
	return s
}

func sendItAsCode(s, extension string) string {
	return "```" + extension + "\n" + s + "```"
}

// ingoreWhiteSpaces is used to seperate commands when $help [command] to rendre interface with the bot dynamic.
func ingoreWhiteSpaces(s string) string {

	var sglobal string
	s = strings.Replace(s, " ", "", -1)

	for i := 0; i < len(s); i++ {
		ss := string(s[i])
		if ss == "$" && i != 0 {
			sglobal += " " + ss
		} else {
			sglobal += ss
		}
	}
	return sglobal
}

// splitArgs is used with the compare command to rendre interface with it dynamic
func splitArgs(s string) []string {

	var sglobal []string

	s = strings.Replace(s, " ", "", -1)

	for i := 0; i < len(s); i++ {
		ss := string(s[i])
		if ss == "e" {
			sglobal = append(sglobal, s[:i+1])
		}
		if ss == "<" {
			sglobal = append(sglobal, s[i:i+3])
		}
	}
	return sglobal
}

// to return true or false when $test is called.
func randomBool() bool {
	return rand.Intn(100) > 50
}
