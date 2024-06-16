package main

import (
	"embed"

	"github.com/kawana77b/univenv/cmd"
)

var Version string = "0.0.1"

//go:embed template
var fs embed.FS

func main() {
	cmd.Fs = fs
	cmd.Version = Version
	cmd.Execute()
}
