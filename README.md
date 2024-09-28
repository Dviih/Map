# Map

## An enhancement for `sync.Map` implementing types and map output

---

## Map
### Implemented on top of `sync.Map` with more methods.

---

# Usage

## Map
- `Load` - Loads a key and returns value.
- `Store` - Stores a value as a key.
- `Delete` - Deletes a key.
- `Clear` - Clears the map.
- `Swap` - Swaps a key to another value and returns previous value.
- `LoadOrStore` - Loads a value if exists or store it.
- `LoadAndDelete` - Loads a value and deletes it.
- `CompareAndSwap` - Compares a value and if matches swap it.
- `CompareAndDelete` - Compares a value and if it matches delete it.
- `Range` - Ranges through all keys of the map.
- `Map` - Returns a `map[K]V`.
- `Len` - Returns the length of the map.

# Example
```go
package main

import (
	"fmt"
	"github.com/Dviih/Map"
)

func main() {
	m := Map.Map[string, int]{}
	
	m.Store("one", 1)
	
	fmt.Println(m.Load("one"))
}
```

---
#### Made for Gophers by @Dviih