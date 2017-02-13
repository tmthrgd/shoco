# shoco

[![GoDoc](https://godoc.org/github.com/tmthrgd/shoco?status.svg)](https://godoc.org/github.com/tmthrgd/shoco)
[![Build Status](https://travis-ci.org/tmthrgd/shoco.svg?branch=master)](https://travis-ci.org/tmthrgd/shoco)

**shoco** is a Golang package, based on [the **shoco** C library](https://github.com/Ed-von-Schleck/shoco), to compress and decompress short strings. It is very fast and easy to use. The default compression model is optimized for english words, but it is possible to generate your own compression model based on your specific input data.

## Download

```
go get github.com/tmthrgd/shoco
```

## Benchmark

```
BenchmarkCompress/#0-0-8         	10000000	       177 ns/op
BenchmarkCompress/#1-4-8         	 5000000	       264 ns/op	  15.14 MB/s
BenchmarkCompress/#2-5-8         	 5000000	       349 ns/op	  14.30 MB/s
BenchmarkCompress/#3-240-8       	  300000	      4768 ns/op	  50.33 MB/s
BenchmarkCompress/#4-58-8        	 1000000	      1180 ns/op	  49.15 MB/s
BenchmarkCompress/#5-20-8        	 2000000	       684 ns/op	  29.23 MB/s
BenchmarkCompress/#6-13-8        	 3000000	       450 ns/op	  28.83 MB/s
BenchmarkCompress/#7-111-8       	  500000	      2748 ns/op	  40.38 MB/s
BenchmarkCompress/#8-9-8         	 5000000	       400 ns/op	  22.45 MB/s
BenchmarkCompress/#9-13-8        	 3000000	       452 ns/op	  28.75 MB/s
BenchmarkCompress/#10-13-8       	 3000000	       433 ns/op	  30.02 MB/s
BenchmarkCompress/#11-10-8       	 3000000	       398 ns/op	  25.10 MB/s
BenchmarkCompress/#12-15-8       	 3000000	       462 ns/op	  32.44 MB/s
BenchmarkCompress/#13-35-8       	 2000000	       974 ns/op	  35.91 MB/s
BenchmarkCompress/#14-6-8        	 5000000	       330 ns/op	  18.18 MB/s
BenchmarkCompress/#15-2-8        	10000000	       218 ns/op	   9.14 MB/s
BenchmarkCompress/#16-4-8        	 5000000	       269 ns/op	  14.85 MB/s
BenchmarkCompress/#17-4-8        	 5000000	       269 ns/op	  14.82 MB/s
BenchmarkCompress/#18-9-8        	 5000000	       297 ns/op	  30.23 MB/s
BenchmarkCompress/#19-2-8        	10000000	       193 ns/op	  10.35 MB/s
BenchmarkCompress/#20-4-8        	10000000	       200 ns/op	  19.94 MB/s
BenchmarkCompress/#21-4-8        	10000000	       191 ns/op	  20.94 MB/s
BenchmarkDecompress/#0-0-8       	10000000	       120 ns/op
BenchmarkDecompress/#1-2-8       	10000000	       196 ns/op	  10.16 MB/s
BenchmarkDecompress/#2-3-8       	10000000	       214 ns/op	  14.02 MB/s
BenchmarkDecompress/#3-169-8     	  500000	      4170 ns/op	  40.52 MB/s
BenchmarkDecompress/#4-39-8      	 1000000	      1316 ns/op	  29.63 MB/s
BenchmarkDecompress/#5-24-8      	 3000000	       470 ns/op	  51.04 MB/s
BenchmarkDecompress/#6-17-8      	 5000000	       369 ns/op	  45.99 MB/s
BenchmarkDecompress/#7-79-8      	 1000000	      2255 ns/op	  35.02 MB/s
BenchmarkDecompress/#8-18-8      	 5000000	       284 ns/op	  63.29 MB/s
BenchmarkDecompress/#9-22-8      	 5000000	       333 ns/op	  65.96 MB/s
BenchmarkDecompress/#10-22-8     	 5000000	       327 ns/op	  67.26 MB/s
BenchmarkDecompress/#11-20-8     	 5000000	       304 ns/op	  65.77 MB/s
BenchmarkDecompress/#12-25-8     	 5000000	       360 ns/op	  69.35 MB/s
BenchmarkDecompress/#13-46-8     	 2000000	       858 ns/op	  53.60 MB/s
BenchmarkDecompress/#14-12-8     	10000000	       174 ns/op	  68.65 MB/s
BenchmarkDecompress/#15-4-8      	10000000	       176 ns/op	  22.71 MB/s
BenchmarkDecompress/#16-8-8      	10000000	       216 ns/op	  36.92 MB/s
BenchmarkDecompress/#17-8-8      	10000000	       222 ns/op	  36.00 MB/s
BenchmarkDecompress/#18-6-8      	 5000000	       344 ns/op	  17.43 MB/s
BenchmarkDecompress/#19-3-8      	10000000	       183 ns/op	  16.36 MB/s
BenchmarkDecompress/#20-5-8      	10000000	       190 ns/op	  26.31 MB/s
BenchmarkDecompress/#21-5-8      	10000000	       188 ns/op	  26.49 MB/s
```

```
BenchmarkWords/Compress-8        	     100	  21806321 ns/op	  43.05 MB/s
BenchmarkWords/Decompress-8      	     100	  16730975 ns/op	  39.60 MB/s
--- BENCH: BenchmarkWords
	shoco_test.go:228: len(in)  = 938848B
	shoco_test.go:229: len(out) = 662545B
	shoco_test.go:230: ratio    = 0.705700%
```

## License

Unless otherwise noted, the shoco source files are distributed under the Modified BSD License
found in the LICENSE file.
