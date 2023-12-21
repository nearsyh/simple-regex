This is a simple implementation of [Regular Expression Matching Can Be Simple And Fast](https://swtch.com/~rsc/regexp/regexp1.html) in Go.

Example
```
go run cmd/main.go --str="abc" --pat="ab+(c|d)e"
```

You can also provide `--debug` to generate a dot file showing the regex statemachine.