package main

import (
	"github.com/thmhoag/urlbyter/pkg/cli"
	"os"
)

var (
	semver, commit, built, goversion string
)

func main() {

	cli.Semver = semver
	cli.Commit = commit
	cli.Built = built
	cli.GoVersion = goversion

	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}