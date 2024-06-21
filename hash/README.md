# Go Hash

Simple wrapper functions which take arbitrary string input and return string hashes in hex encoding.

The author is aware that this package is somewhat perfunctory, but likes not having to think about the `crypto` packages.

## Example

```go
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/annybs/go/hash"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(errors.New("input required"))
		os.Exit(1)
	}

	fmt.Println(hash.SHA256(os.Args[1]))
}
```

Execute `go run main.go "any test string you like"` to create SHA256 output.

## License

See [LICENSE.md](../LICENSE.md)
