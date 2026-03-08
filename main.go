package main

import (
	"authentication_backend/cmd"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		cmd.Cmd()
	}
}
