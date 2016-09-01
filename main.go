package main

import (
	"github.com/yamamoto-febc/sacloud-upload-archive/cli"
	"os"
)

func main() {
	cli := &cli.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
