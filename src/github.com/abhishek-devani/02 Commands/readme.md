# Commands

1. **go run:** Compiles and runs a Go program in one step. For example:

```bash
go run filename.go
```

2. **go build:** Compiles the Go program and generates an executable file. For example:

```bash
go build filename.go
```

3. **go install:** Compiles and installs the Go program, placing the executable in the `bin` directory of the workspace. For example:

```bash
go install
```

4. **go get:** Downloads and installs packages and their dependencies. It's often used to fetch external packages from version control systems like Git. For example:

```bash
go get github.com/example/package
```

5. **go test:** Runs tests in the current package. For example:

```bash
go test
```

6. **go fmt:** Formats Go source code files. For example:

```bash
go fmt filename.go
```

7. **go mod init:** Initializes a new module, creating a `go.mod` file. For example:

```bash
go mod init module-name
```

8. **go mod tidy:** Prunes any dependencies that are no longer required and adds any dependencies needed for the code. For example:

```bash
go mod tidy
```

9.  **go version:** Displays the Go version installed on your system.