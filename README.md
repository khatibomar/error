# Experimental `Error` Wrapper

Adds `file:line` prefix to Go errors for quick debugging.

```go
err := errors.New("kanna")

err = error.Inject(err)
fmt.Println(err) // main.go:11: kanna

file, line, originalErr := error.Extract(err)
fmt.Println(file, line, originalErr) // main.go 11 kanna
```

- Re-wrapping replaces the old prefix.
- Nil errors stay nil.
- Only a single prefix is added.

Experimental & unidiomatic: I prefer to wrap errors with clear messages to easily locate them.
But if for what ever reason you though this is a good idea. Don't pull the
library, just copy paste the file and use it in your project.
