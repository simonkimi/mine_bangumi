package main

import (
	"fmt"
	"github.com/otiai10/copy"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Working directory: %s\n", wd)
	err = os.RemoveAll("cmd/server/dist")
	if err != nil {
		panic(err)
	}

	err = copy.Copy("frontend/dist", "cmd/server/dist")
	if err != nil {
		panic(err)
	}
}
