#go test -bench=. -benchmem -gcflags -m
go test -bench=. -benchmem -cpuprofile cpu.out
