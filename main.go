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
)
`

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

		// now for the slightly involved part -- generating the bazel repo name.
		// For an example, it transforms the following string like this:
		// github.com/Masterminds/vcs -> com_github_masterminds_vcs
		// it takes the domain, then inverts it so that the TLD goes first.

		// first, let's divide it by the slash (/), separating the domain and the urlpath
		// we should have exactly three divisions
		splits := strings.SplitAfterN(imports.Name, "/", 2)

		if len(splits) != 2 {
			panic(fmt.Sprintf("One of the repo names is invalid :%v", imports.Name))
		}

		fulldomain, urlpath := splits[0], splits[1]

		// just slugify the urlpath
		slugURLPath := strings.Replace(urlpath, ".", "_", -1)
		slugURLPath = strings.Replace(slugURLPath, "/", "_", -1)
		slugURLPath = strings.ToLower(slugURLPath)

		// further split the domain into the TLD and the actual domain
		domainsplit := strings.Split(fulldomain, ".")
		domain, tld := domainsplit[0], domainsplit[1]
		// remove the slash from the tld that Split leaves behind
		tld = strings.Replace(tld, "/", "", -1)

		// put everything together
		bazelName := fmt.Sprintf("%s_%s_%s", tld, domain, slugURLPath)

		// print everything
		fmt.Printf(goRepoTemplate, bazelName, bazelImportpath, bazelVersion)
	}
}
