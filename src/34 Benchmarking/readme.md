# Benchmarking

- Benchmarking in Go is the process of measuring the performance of your code, specifically the time it takes to execute a particular operation or function.
- The Go testing package provides a built-in mechanism for writing benchmarks alongside your regular tests.
- To write a benchmark in Go, you use the `testing` package and prefix your benchmark function names with "Benchmark".
- The testing framework then runs these benchmarks and reports the time taken for each operation.

### Command

#### Run One Time

```bash
go test -bench .
```

#### Run Multiple Time

```bash
go test -bench . -count 10
```