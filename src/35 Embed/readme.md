# Embed

- The `embed` package in Go provides a way to include files and directories directly in the Go source code,  making it easier to distribute self-contained Go programs or libraries.
- With the `embed` package, you can include files or entire directories in your Go code, and the content of these files or directories will be embedded in the compiled binary.

### Syntax to declare embed in go file

```markdown
1. First import "embed" library then declare it above main function
> //go:embed <file-name>
> var data string
> Ex. //go:embed test.json
```

### Command to build binary

```bash
> go build .
```

### Command to run binary

```bash
> ./<module-name>
> Ex. ./demo
```