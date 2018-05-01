package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := errors.New("my error")
	fmt.Println(err)
}
