package main

import (
	cmds "github.com/indraniel/srasearch/commands"
)

const version = "0.1.0"

func main() {
	cmds.Execute(version)
}
