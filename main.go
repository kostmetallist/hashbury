package main

import "fmt"

func main() {

	h := Init()
	keyValues := make(map[uint32]string)

	keyValues[42] = "foo"
	keyValues[33] = "bar"
	keyValues[7] = "baz"
	keyValues[2] = "qux"

	for k, v := range keyValues {
		h.Insert(k, v)
	}

	for k := range keyValues {
		fmt.Printf("Contains  %d? %t\n", k, h.Has(k))
		fmt.Printf("Value for %d is '%s'\n", k, h.Get(k))
	}

	fmt.Println(h)
}
