package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main () {
	err := testfn()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func testfn () error {
	return errors.New ("testerror")
}
