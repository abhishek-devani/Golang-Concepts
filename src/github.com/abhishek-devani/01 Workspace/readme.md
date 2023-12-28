# Workspace

workspace is a specific directory structure that contains Go source code files and other related files. The workspace is the base directory where you organize your Go projects

```
workspace/
|-- bin/
|-- pkg/
|-- src/
    |-- github.com/
        |-- yourusername/
            |-- yourproject/
                |-- .git/
                |-- main.go
                |-- other.go
```

- **bin:** The compiled executable binaries will be placed here after running the `go install` command.
- **pkg:** Intermediate compiled packages (object files) will be stored here. This directory is used to store packages that are compiled from your code and are then linked to create the final executable.
- **src:** This is where your actual Go source code resides. It's organized using the package naming convention. The packages are organized under the src directory based on their import path. In the example above, the import path is github.com/yourusername/yourproject.

