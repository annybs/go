# Go Libraries

General-purpose [Go](https://go.dev) packages for use in Recipeer projects.

## System Requirements

- [Go v1.21](https://go.dev/dl/)

## Testing

Some packages include tests. Run the following snippet to perform all tests automatically:

```sh
find . -type f -iname '*_test.go' | xargs dirname | uniq | sort | xargs go test
```

## License

See [LICENSE.md](./LICENSE.md)
