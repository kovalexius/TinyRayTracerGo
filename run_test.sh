#go test -bench=. -benchmem -gcflags -m
go test -bench=BenchmarkMain -benchmem -cpuprofile cpu.out
