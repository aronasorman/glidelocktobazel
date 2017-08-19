package main

import (
	"fmt"
	"github.com/Masterminds/glide/cfg"
	"log"
	"os"
	"strings"
)

var goRepoTemplate = `
go_repository(
    name="%s",
    importpath="%s",
    commit="%s",
)
`

// copied from https://github.com/bazelbuild/rules_go/blob/master/go/tools/gazelle/resolve/resolve_external.go
// ImportPathToBazelRepoName converts a Go import path into a bazel repo name
// following the guidelines in http://bazel.io/docs/be/functions.html#workspace
func ImportPathToBazelRepoName(importpath string) string {
	components := strings.Split(importpath, "/")
	labels := strings.Split(components[0], ".")
	var reversed []string
	for i := range labels {
		l := labels[len(labels)-i-1]
		reversed = append(reversed, l)
	}
	repo := strings.Join(append(reversed, components[1:]...), "_")
	return strings.NewReplacer("-", "_", ".", "_").Replace(repo)
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Arguments must always be 1 -- filename.")
		return
	}

	filename := os.Args[1]

	lockfile, err := cfg.ReadLockFile(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}

	for _, imports := range lockfile.Imports {
		// do the easy parts of the go_repository first
		bazelImportpath := imports.Name
		bazelVersion := imports.Version

		bazelName := ImportPathToBazelRepoName(bazelImportpath)

		// print everything
		fmt.Printf(goRepoTemplate, bazelName, bazelImportpath, bazelVersion)
	}
}
