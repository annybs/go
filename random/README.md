# Go Random

Functions to create random values. That's about the size of it.

## Example

```go
package main

import (
	"fmt"

	"github.com/annybs/go/random"
)

func main() {
	n := make([]bool, 10)
	for i := range n {
		fmt.Println(random.Str(4 + i))
	}
}
```

## License

See [LICENSE.md](../LICENSE.md)
