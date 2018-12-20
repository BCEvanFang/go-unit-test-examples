# Go Unit Test

## Test Commands
```sh
go test

# Output details
go test -v

# Run single test file
go test ./mymath/mymath_test.go

# Run all tests
go test ./...

# Get code test coverage
go test ./... -v -cover

# Code coverage with HTML report
go test ./... -coverprofile=cover.out
go tool cover -html=cover.out
```