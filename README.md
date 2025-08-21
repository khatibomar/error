# Experimental `Error` Wrapper

Adds `file:line` prefix to Go errors for quick debugging.

```go
10 err := errors.New("kanna")
11 err = error.Inject(err) // adds caller info
12 fmt.Println(err)  // main.go:11: kanna
```

Extract the file:line from an error:

```go
15 file, line := error.Extract(err)
16 fmt.Println(file, line) // main.go 11
```

- Re-wrapping replaces the old prefix.
- Nil errors stay nil.
- Only a single prefix is added.

Experimental & unidiomatic I prefer to wrap errors with clear messages to easily locate them.
But if you like it for some reason, copy the 2 small methods, no need to pull it as a library for such simple thing :)
