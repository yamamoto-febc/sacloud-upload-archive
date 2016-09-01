package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun_tokenFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{OutStream: outStream, ErrStream: errStream}
	args := strings.Split("./sacloud-upload-archive -token", " ")

	status := cli.Run(args)
	_ = status
}

func TestRun_secretFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{OutStream: outStream, ErrStream: errStream}
	args := strings.Split("./sacloud-upload-archive -secret", " ")

	status := cli.Run(args)
	_ = status
}

func TestRun_zoneFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{OutStream: outStream, ErrStream: errStream}
	args := strings.Split("./sacloud-upload-archive -zone", " ")

	status := cli.Run(args)
	_ = status
}
