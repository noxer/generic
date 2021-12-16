package main

import (
	"fmt"

	"github.com/noxer/generic"
	"github.com/noxer/generic/set"
)

func main() {
	s := make(set.Set[string])
	var i generic.MutableCollection[string] = s

	i.Put("Hello")
	i.Put("World")
	fmt.Println(i.Any())
}
