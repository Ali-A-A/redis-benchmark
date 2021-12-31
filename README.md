# redis-benchmark

### Simple get, mget and pipelined get benchmark.

## Usage

1. `git clone https://github.com/Ali-A-A/redis-benchmark.git`
2. `cd ./redis-benchmark`
3. `docker-compose up -d`
4. `cd ./test`
5. `go test -bench=.  -benchmem  -cpu 2  -run=^a`

## Sample Output

```
goos: darwin
goarch: arm64
pkg: bench/test
BenchmarkFindMGet-2                   24          46734172 ns/op         2726839 B/op        606 allocs/op
BenchmarkFindPipeline-2               22          50437112 ns/op         2750784 B/op        816 allocs/op
BenchmarkFindGet-2                     6         185732521 ns/op         2752810 B/op       1200 allocs/op
PASS
ok      bench/test      5.577s
```


