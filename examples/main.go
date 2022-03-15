package main

import (
	"fmt"

	"github.com/polynomialspace/buildinfo"
)

func main() {
	fmt.Println(buildinfo.Info["vcs.revision"])
	fmt.Println(buildinfo.VCSRevision)

	fmt.Println(buildinfo.Info["vcs.revision"][:7])
	fmt.Println(buildinfo.VCSRev)
}
