# Go Arango

This package offers a simple syntax to construct [AQL queries](https://docs.arangodb.com/3.11/aql/) with bind parameters, which can be used with the [official ArangoDB driver](https://pkg.go.dev/github.com/arangodb/go-driver).

While it's entirely possible to write static queries and bind parameters externally, sometimes you need more flexibility to create queries, particularly if you are providing a consumer API with querying capabilities. See also:

- [Go Query String](../qs/README.md)
- [Go REST](../rest/README.md)

This package allows you to build your query piece-by-piece and attach parameters in whatever way serves you best.

> :warning: This package has not yet been tested with the [v2 driver](https://pkg.go.dev/github.com/arangodb/go-driver/v2/arangodb).

## Examples

In this example, the `tastiness` value is bound to a parameter _while_ constructing the query.

```go
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/annybs/go/arango"
)

func main() {
	query := arango.NewQuery().Append("FOR f IN @@foodColl", "food")

	tastiness := 1
	if len(os.Args) > 1 {
		t, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(errors.New("invalid minimum tastiness"))
			os.Exit(1)
		}
		tastiness = t
	}

	query.
		Append("FILTER f.tastiness > @tastiness", tastiness).
		Append("RETURN f")

	fmt.Println("Query:", query.String())
	fmt.Println("Params:", query.Params)
}
```

For simpler cases, you may prefer to construct your query first and then bind parameters separately:

```go
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/annybs/go/arango"
)

func main() {
	query := arango.NewQuery().
		Append("FOR f IN @@foodColl", "food").
		Append("FILTER f.tastiness > @tastiness").
		Append("RETURN f")

	tastiness := 1
	if len(os.Args) > 1 {
		t, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(errors.New("invalid minimum tastiness"))
			os.Exit(1)
		}
		tastiness = t
	}

	query.Bind("tastiness", tastiness)

	fmt.Println("Query:", query.String())
	fmt.Println("Params:", query.Params)
}
```

## License

See [LICENSE.md](../LICENSE.md)
