# speedresize

Test speed of some image libraries

Run test in docker:

```bash

#check out repo
# run in docker, mount currend dir into docker:
docker run --rm -it -v $(pwd):/src debian:bookworm /bin/bash

# install deps:
apt install -y libjpeg62-turbo-dev ca-certificates golang libvips-dev

# run tests
go test -bench .
```

## example test output
```txt
goos: linux
goarch: amd64
pkg: resizehack
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkVips-8    	      16	  81217453 ns/op
--- BENCH: BenchmarkVips-8
    main_test.go:13: Took 68.599523ms
    main_test.go:13: Took 66.074989ms
    main_test.go:13: Took 89.021341ms
    main_test.go:13: Took 87.814882ms
    main_test.go:13: Took 76.589917ms
    main_test.go:13: Took 71.086804ms
    main_test.go:13: Took 70.026369ms
    main_test.go:13: Took 71.041156ms
    main_test.go:13: Took 65.054919ms
    main_test.go:13: Took 90.342536ms
	... [output truncated]
BenchmarkTurbo-8   	      15	  70350140 ns/op
--- BENCH: BenchmarkTurbo-8
    main_test.go:28: Took 76.386543ms
    main_test.go:28: Took 81.33853ms
    main_test.go:28: Took 72.704254ms
    main_test.go:28: Took 66.325781ms
    main_test.go:28: Took 62.094701ms
    main_test.go:28: Took 73.301152ms
    main_test.go:28: Took 66.257635ms
    main_test.go:28: Took 68.158708ms
    main_test.go:28: Took 81.098838ms
    main_test.go:28: Took 85.938362ms
	... [output truncated]
BenchmarkNfnt-8    	       2	 614070016 ns/op
--- BENCH: BenchmarkNfnt-8
    main_test.go:43: Took 676.976127ms
    main_test.go:43: Took 667.301415ms
    main_test.go:43: Took 560.729778ms
PASS
ok  	resizehack	5.147s
```
