# allocations
Understanding the differences between the stack and the heap.

## How to Use It

### View compiler output
```
go build -gcflags '-m -l'
```

### Run benchmarks
```
go test -bench . -benchmem
```

### Generate & view `CreateCopy` trace
```
go test -run TestCopyIt -trace=copy_trace.out
go tool trace copy_trace.out
```

### Generate & view `CreatePointer` trace
```
go test -run TestPointerIt -trace=copy_trace.out
go tool trace pointer_trace.out
```
