package main

import (
	"fmt"
	"luffy/src/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Please type in commands\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
