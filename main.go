package main

import (
	"os"

	"smurk/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
