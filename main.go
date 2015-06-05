package main

import (
	cmds "github.com/indraniel/srasearch/commands"
)

const version = "0.1.2"

func main() {
	cmds.Execute(version)
}
