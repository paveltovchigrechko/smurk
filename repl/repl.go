package repl

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = "smurf >> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	fmt.Fprintln(out, "This is Smurk. Type commands to do Smurks's things.")

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		} else {
			fmt.Fprintf(out, "You have entered %q\n", scanner.Text())
		}
	}

}
