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
BenchmarkCompress/#0-0-8         	10000000	       129 ns/op
BenchmarkCompress/#1-4-8         	10000000	       210 ns/op	  19.04 MB/s
BenchmarkCompress/#2-5-8         	 5000000	       288 ns/op	  17.32 MB/s
BenchmarkCompress/#3-240-8       	  300000	      4324 ns/op	  55.50 MB/s
BenchmarkCompress/#4-58-8        	 1000000	      1014 ns/op	  57.15 MB/s
BenchmarkCompress/#5-20-8        	 3000000	       555 ns/op	  36.03 MB/s
BenchmarkCompress/#6-13-8        	 3000000	       404 ns/op	  32.12 MB/s
BenchmarkCompress/#7-111-8       	 1000000	      2401 ns/op	  46.23 MB/s
BenchmarkCompress/#8-9-8         	 5000000	       341 ns/op	  26.36 MB/s
BenchmarkCompress/#9-13-8        	 3000000	       387 ns/op	  33.53 MB/s
BenchmarkCompress/#10-13-8       	 5000000	       427 ns/op	  30.43 MB/s
BenchmarkCompress/#11-10-8       	 5000000	       389 ns/op	  25.70 MB/s
BenchmarkCompress/#12-15-8       	 3000000	       439 ns/op	  34.12 MB/s
BenchmarkCompress/#13-35-8       	 2000000	       828 ns/op	  42.23 MB/s
BenchmarkCompress/#14-6-8        	10000000	       308 ns/op	  19.44 MB/s
BenchmarkCompress/#15-2-8        	10000000	       203 ns/op	   9.83 MB/s
BenchmarkCompress/#16-4-8        	10000000	       257 ns/op	  15.51 MB/s
BenchmarkCompress/#17-4-8        	10000000	       279 ns/op	  14.33 MB/s
BenchmarkCompress/#18-2-8        	10000000	       193 ns/op	  10.32 MB/s
BenchmarkCompress/#19-4-8        	10000000	       201 ns/op	  19.83 MB/s
BenchmarkCompress/#20-4-8        	10000000	       205 ns/op	  19.50 MB/s
BenchmarkDecompress/#0-0-8       	10000000	       150 ns/op
BenchmarkDecompress/#1-2-8       	 5000000	       259 ns/op	   7.71 MB/s
BenchmarkDecompress/#2-3-8       	 5000000	       262 ns/op	  11.41 MB/s
BenchmarkDecompress/#3-169-8     	  500000	      4191 ns/op	  40.32 MB/s
BenchmarkDecompress/#4-39-8      	 1000000	      1219 ns/op	  31.98 MB/s
BenchmarkDecompress/#5-24-8      	 3000000	       448 ns/op	  53.55 MB/s
BenchmarkDecompress/#6-17-8      	 5000000	       336 ns/op	  50.46 MB/s
BenchmarkDecompress/#7-79-8      	 1000000	      2187 ns/op	  36.11 MB/s
BenchmarkDecompress/#8-18-8      	 5000000	       287 ns/op	  62.54 MB/s
BenchmarkDecompress/#9-22-8      	 5000000	       344 ns/op	  63.95 MB/s
BenchmarkDecompress/#10-22-8     	 5000000	       336 ns/op	  65.31 MB/s
BenchmarkDecompress/#11-20-8     	 5000000	       317 ns/op	  62.93 MB/s
BenchmarkDecompress/#12-25-8     	 5000000	       363 ns/op	  68.82 MB/s
BenchmarkDecompress/#13-46-8     	 2000000	       885 ns/op	  51.94 MB/s
BenchmarkDecompress/#14-12-8     	 5000000	       267 ns/op	  44.81 MB/s
BenchmarkDecompress/#15-4-8      	10000000	       195 ns/op	  20.46 MB/s
BenchmarkDecompress/#16-8-8      	10000000	       211 ns/op	  37.74 MB/s
BenchmarkDecompress/#17-8-8      	10000000	       202 ns/op	  39.43 MB/s
BenchmarkDecompress/#18-3-8      	10000000	       154 ns/op	  19.42 MB/s
BenchmarkDecompress/#19-5-8      	10000000	       154 ns/op	  32.31 MB/s
BenchmarkDecompress/#20-5-8      	10000000	       156 ns/op	  31.92 MB/s
```

## License

Unless otherwise noted, the shoco source files are distributed under the Modified BSD License
found in the LICENSE file.
