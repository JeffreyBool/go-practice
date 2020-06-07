/**
* Author: JeffreyBool
* Date: 2020/6/5
* Time: 22:33
* Software: GoLand
 */

package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrCouldNotOpen = errors.New("could not open")

type errNotFound struct {
	file string
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("file %q not found", e.file)
}

func IsNotFoundError(err error) bool {
	_, ok := err.(errNotFound)
	return ok
}

func FoundOpen(file string) error {
	return errNotFound{file: file}
}

func usr() {
	if err := FoundOpen(""); err != nil {
		if IsNotFoundError(err) {
			fmt.Println("errNotFound")
		} else {
			panic("unknown error")
		}
	}
}

func Open() error {
	return ErrCouldNotOpen
}

func main() {
	if err := Open(); err != nil {
		if err == ErrCouldNotOpen {
			// handle
			fmt.Println(err)
		} else {
			panic("unknown error")
		}
	}

	usr()
}
