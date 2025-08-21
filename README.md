# Experimental `Error` Wrapper

Adds `file:line` prefix to Go errors for quick debugging.

```go
err := errors.New("kanna")

err = Inject(err)
fmt.Println(err) // main.go:11: kanna

file, line, originalErr := Extract(err)
fmt.Println(file, line, originalErr) // main.go 11 kanna
```

- Re-wrapping replaces the old prefix.
- Nil errors stay nil.
- Only a single prefix is added.

Experimental & unidiomatic: I prefer to wrap errors with clear messages to easily locate them.
That's why it's not a library but a "program"
So it's more like a snippet.
