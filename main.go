package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the monkey programming language!\n", user)
	fmt.Printf("Start typing away!\n")
	repl.Start(os.Stdin, os.Stdout)
}
