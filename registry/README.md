# Go Registry

This package provides a simple registry implementation which allows storing arbitrary values in a memory map.

You can also lock specific keys to prevent them being written, even if no value has yet been set. Locked keys cannot be unlocked.

## Example

```go
package main

import (
	"fmt"

	"github.com/annybs/go/registry"
)

func main() {
	r := registry.New()

	r.Set("some key", "some text")
	fmt.Println(r.Get("some key"))

	r.Lock("some key")
	if err := r.Set("some key", "different text"); err != nil {
		fmt.Println(err)
	}
}
```

## License

See [LICENSE.md](../LICENSE.md)
