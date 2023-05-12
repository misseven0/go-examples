package main

import (
	"fmt"
	"os"

	"github.com/misseven0/fastwalk"
)

func main() {
	fastwalk.Walk(".", func(path string, typ os.FileMode) error {
		fmt.Printf("|-%v \n", path)
		return nil
	})
}
