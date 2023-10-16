# Test Coverage

### Run tests with coverage
`go test -coverprofile=coverage.out ./...`

### View coverage
`go tool cover -html=coverage.out`

# Benchmarks

### Run benchmarks
`go test -bench=. ./...`

### Run benchmarks with count
`go test -bench=. -count=10 ./...`

### Run benchmarks with time
`go test -bench=. -benchtime=10s ./...`

### Run benchmarks with memory allocations
`go test -bench=. -benchmem ./...`

### Run benchmarks with memory allocations and CPU profiling
`go test -bench=. -benchmem -cpuprofile=cpu.out ./...`

# Fuzzing Tests

### Run fuzzing tests
`go test -fuzz=. ./...`

### Run fuzzing tests with count
`go test -fuzz=. -count=10 ./...`
