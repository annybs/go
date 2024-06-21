# Go Validate

A suite of straightforward validation functions. You put something in, you get back `nil` or an error.

## Error handling

You can use `errors.Is()` to ascertain the type of errors thrown by validation functions. This may be helpful to control side effects, particularly if using multiple validators and returning early (similar to a strongly-typed try-catch). For example:

## Example

```go
package main

import (
	"errors"
	"fmt"

	"github.com/annybs/go/validate"
)

func main() {
	v := validate.Equal("a")
	if err := v("b"); err != nil {
		if errors.Is(err, validate.ErrNotEqual) {
			fmt.Println("failed successfully")
		} else {
			fmt.Println("failed unsuccessfully")
		}
	}
}
```

## License

See [LICENSE.md](../LICENSE.md)
