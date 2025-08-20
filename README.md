# Experimental `Error` Wrapper

Adds `file:line` prefix to Go errors for quick debugging.

```go
10 err := errors.New("kanna")
11 err = error.Error(err) // adds caller info
12 fmt.Println(err) // main.go:11: kanna
```

- Re-wrapping replaces the old prefix.
- Nil errors stay nil.
- Only a single prefix is added.

Experimental & unidiomatic I prefer to wrap errors with clear messages to easily locate them.
But if you like it for some reason, copy the single method, no need to pull it as a library for such simple thing :)
