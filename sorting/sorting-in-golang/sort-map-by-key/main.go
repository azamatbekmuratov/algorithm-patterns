// A map is an unordered collection of key-value pairs.
// If you need a stable iteration order, you must maintain a separate data structure.
package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
