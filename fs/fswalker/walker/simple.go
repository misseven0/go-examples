package main

import (
	"fmt"
	"os"

	"github.com/saracen/walker"
)

func main() {
	// walk function called for every path found
	walkFn := func(pathname string, fi os.FileInfo) error {
		fmt.Printf("%s: %d bytes\n", pathname, fi.Size())
		return nil
	}

	// error function called for every error encountered
	errorCallbackOption := walker.WithErrorCallback(func(pathname string, err error) error {
		// ignore permissione errors
		if os.IsPermission(err) {
			return nil
		}
		// halt traversal on any other error
		return err
	})

	walker.Walk("/tmp", walkFn, errorCallbackOption)
}
