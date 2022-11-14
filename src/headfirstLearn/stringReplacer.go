package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "G# Lang"
	replac := strings.NewReplacer("#", "O")
	fixed := replac.Replace(str)

	fmt.Println(fixed)
}
