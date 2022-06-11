package main

import (
	"fmt"
	"sort"
)

func main() {
	//define map
	states := make(map[string]string)
	fmt.Println(states)
	// popluate key, values
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"

	fmt.Println(states)
	// get value by key from map
	california := states["CA"]

	fmt.Println(california)

	//delete item in map
	delete(states, "OR")
	// add item
	states["NY"] = "New York"

	fmt.Println(states)

	// loop over map with range
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	//sorting
	// slice of map keys
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	// sorting keys
	sort.Strings(keys)
	fmt.Println(keys)
	// iter keys and print values
	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
