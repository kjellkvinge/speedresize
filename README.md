# speedresize

Test speed of some image libraries

Run test in docker:

```bash

#check out repo
# run in docker, mount currend dir into docker:

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
BenchmarkVips-8    	       2	 639326164 ns/op
--- BENCH: BenchmarkVips-8
    main_test.go:13: Took 649.669641ms
    main_test.go:13: Took 667.909559ms
    main_test.go:13: Took 610.703419ms
BenchmarkTurbo-8   	      20	  57012006 ns/op
--- BENCH: BenchmarkTurbo-8
    main_test.go:28: Took 69.73603ms
    main_test.go:28: Took 61.513814ms
    main_test.go:28: Took 66.676875ms
    main_test.go:28: Took 53.226282ms
    main_test.go:28: Took 55.783916ms
    main_test.go:28: Took 61.697351ms
    main_test.go:28: Took 54.503784ms
    main_test.go:28: Took 51.944645ms
    main_test.go:28: Took 55.309774ms
    main_test.go:28: Took 60.28305ms
	... [output truncated]
BenchmarkNfnt-8    	       2	 560933163 ns/op
--- BENCH: BenchmarkNfnt-8
    main_test.go:43: Took 547.030001ms
    main_test.go:43: Took 518.591504ms
    main_test.go:43: Took 603.217821ms
PASS
ok  	resizehack	7.329s

```