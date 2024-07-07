package commands

import (
	"fmt"
	"strings"
)

type Command string

const (
	HELP  Command = "help"
	SMURK Command = "smurk"
)

var descriptions = map[Command]string{
	HELP:  "Help command",
	SMURK: "This is Smurk.",
}

func ReadCommand(input string) string {
	arguments := strings.Fields(input)

	switch Command(arguments[0]) {
	case SMURK:
		return descriptions[SMURK]

	case HELP:
		result := ""
		for command, description := range descriptions {
			result += fmt.Sprintf("%s\t\t\t%s\n", command, description)
		}
		return result

	default:
		return fmt.Sprintf("Smurk does not know command %q", arguments[0])
	}
}
