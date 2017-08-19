package main

import (
	"fmt"
	"github.com/Masterminds/glide/cfg"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Arguments must always be 1 -- filename.")
		return
	}

	filename := os.Args[1]

	out, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}

	lockfile := cfg.Lockfile{}

	yaml.Unmarshal(out, &lockfile)

}
