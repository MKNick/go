package main

import (
	"fmt"
	"test/bitarray"
)

func main() {
	var x, y bitarray.IntSet
	x.Add(43)
	x.Add(54)
	x.Add(160)
	x.Add(320)
	y.Add(1)
	(&y).UnionWith(&x)
	fmt.Println("xcount:", x.Count(), "ycount:", y.Count())
	x.Clear()
	fmt.Println("y has 54:", y.Has(54))
	fmt.Println("xcount:", x.Count(), "ycount:", y.Count())
}
