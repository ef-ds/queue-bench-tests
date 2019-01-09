# Performance

Below compares the queue [benchmark tests](BENCHMARK_TESTS.md) results with the other tested queues.

## Running the Tests
In the "testdata" directory, we have included the result of local test runs for all queues. Below uses this run to compare the queues, but it's possible and we highly encourage you to run the tests yourself to help validate the results.

To run the tests locally, clone the queue repo, cd to the queue main directory and run below command.

```
go test -benchmem -timeout 60m -bench=. -run=^$
```

This command will run all tests for all queues locally once. This should be good enouh to give you a sense of the queues performance, but to do a proper comparison, elimating test variations, we recommend you to run the tests as detailed [here](BENCHMARK_TESTS.md) by running the tests with multiple counts, splitting the files with [test-splitter](https://github.com/ef-ds/tools/tree/master/testsplitter) and using the [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) tool to aggregate the results.


## Bottom Line
As a general purpose FIFO queue, queue is the data structure that displays the most balanced performance, performing either very competitively or besting all other queues in all the different test scenarios.


## Results
Given the enormous amount of test data, it can be difficult and time consuming to find out the net impact of all the tests, so we generally spend most of the time on the results of the very simple fill tests, which sequentially add and remove N number of items, and the Microservice test, which is a composite test of all other tests.

Below results is for queue [v1.0.0](https://github.com/ef-ds/queue/blob/master/CHANGELOG.md).


### Fill Test Results
queue vs [list](https://github.com/golang/go/tree/master/src/container/list) - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillQueuev1.0.0.txt testdata/BenchmarkFillList.txt
name        old time/op    new time/op    delta
/0-4          35.7ns ± 1%    39.2ns ± 4%    +9.83%  (p=0.000 n=9+9)
/1-4           159ns ± 0%     107ns ± 1%   -32.89%  (p=0.000 n=9+10)
/10-4          549ns ± 1%     748ns ± 2%   +36.26%  (p=0.000 n=9+10)
/100-4        4.55µs ± 2%    7.11µs ± 5%   +56.32%  (p=0.000 n=10+10)
/1000-4       35.0µs ± 1%    70.8µs ± 3%  +102.20%  (p=0.000 n=10+10)
/10000-4       350µs ± 1%     742µs ± 5%  +112.13%  (p=0.000 n=10+10)
/100000-4     3.69ms ± 1%   20.05ms ±15%  +443.63%  (p=0.000 n=9+10)
/1000000-4    42.5ms ± 1%   173.2ms ± 5%  +307.81%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     48.0B ± 0%      ~     (all equal)
/1-4            160B ± 0%      112B ± 0%   -30.00%  (p=0.000 n=10+10)
/10-4           560B ± 0%      688B ± 0%   +22.86%  (p=0.000 n=10+10)
/100-4        7.15kB ± 0%    6.45kB ± 0%    -9.84%  (p=0.000 n=10+10)
/1000-4       33.9kB ± 0%    64.0kB ± 0%   +88.73%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     640kB ± 0%   +98.52%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.40MB ± 0%   +99.04%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%    64.0MB ± 0%   +99.22%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      21.0 ± 0%   +50.00%  (p=0.000 n=10+10)
/100-4           107 ± 0%       201 ± 0%   +87.85%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     2.00k ± 0%   +97.53%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     20.0k ± 0%   +98.36%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      200k ± 0%   +98.44%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     2.00M ± 0%   +98.45%  (p=0.000 n=10+10)
```

queue vs [CustomSliceQueue](testdata_test.go) - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillQueuev1.0.0.txt testdata/BenchmarkFillSlice.txt
name        old time/op    new time/op    delta
/0-4          35.7ns ± 1%    39.4ns ± 3%   +10.22%  (p=0.000 n=9+10)
/1-4           159ns ± 0%      94ns ± 1%   -40.61%  (p=0.000 n=9+8)
/10-4          549ns ± 1%     580ns ± 2%    +5.80%  (p=0.000 n=9+10)
/100-4        4.55µs ± 2%    3.73µs ± 6%   -17.92%  (p=0.000 n=10+9)
/1000-4       35.0µs ± 1%    31.3µs ± 1%   -10.51%  (p=0.000 n=10+10)
/10000-4       350µs ± 1%     381µs ± 1%    +8.99%  (p=0.000 n=10+9)
/100000-4     3.69ms ± 1%    8.45ms ± 1%  +129.17%  (p=0.000 n=9+8)
/1000000-4    42.5ms ± 1%   102.9ms ± 7%  +142.19%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     32.0B ± 0%   -33.33%  (p=0.000 n=10+10)
/1-4            160B ± 0%       56B ± 0%   -65.00%  (p=0.000 n=10+10)
/10-4           560B ± 0%      440B ± 0%   -21.43%  (p=0.000 n=10+10)
/100-4        7.15kB ± 0%    3.67kB ± 0%   -48.66%  (p=0.000 n=10+10)
/1000-4       33.9kB ± 0%    32.4kB ± 0%    -4.50%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     546kB ± 0%   +69.45%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.25MB ± 0%   +94.51%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%    61.2MB ± 0%   +90.47%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      16.0 ± 0%   +14.29%  (p=0.000 n=10+10)
/100-4           107 ± 0%       109 ± 0%    +1.87%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    -0.10%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.61%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

queue vs [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go) - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillQueuev1.0.0.txt testdata/BenchmarkFillImpl7.txt
name        old time/op    new time/op    delta
/0-4          35.7ns ± 1%    36.0ns ± 1%   +0.90%  (p=0.014 n=9+9)
/1-4           159ns ± 0%     134ns ± 1%  -16.04%  (p=0.000 n=9+10)
/10-4          549ns ± 1%     733ns ± 2%  +33.52%  (p=0.000 n=9+10)
/100-4        4.55µs ± 2%    4.08µs ± 1%  -10.32%  (p=0.000 n=10+9)
/1000-4       35.0µs ± 1%    35.9µs ± 1%   +2.49%  (p=0.000 n=10+10)
/10000-4       350µs ± 1%     356µs ± 1%   +1.77%  (p=0.000 n=10+10)
/100000-4     3.69ms ± 1%    3.75ms ± 1%   +1.75%  (p=0.000 n=9+10)
/1000000-4    42.5ms ± 1%    44.1ms ± 1%   +3.76%  (p=0.000 n=9+7)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     48.0B ± 0%     ~     (all equal)
/1-4            160B ± 0%      112B ± 0%  -30.00%  (p=0.000 n=10+10)
/10-4           560B ± 0%      736B ± 0%  +31.43%  (p=0.000 n=10+10)
/100-4        7.15kB ± 0%    4.26kB ± 0%  -40.49%  (p=0.000 n=10+10)
/1000-4       33.9kB ± 0%    33.2kB ± 0%   -2.12%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     323kB ± 0%   +0.12%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.23MB ± 0%   +0.36%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%    32.3MB ± 0%   +0.39%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%     ~     (all equal)
/10-4           14.0 ± 0%      17.0 ± 0%  +21.43%  (p=0.000 n=10+10)
/100-4           107 ± 0%       109 ± 0%   +1.87%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.02k ± 0%   +0.99%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.2k ± 0%   +0.79%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      102k ± 0%   +0.78%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.02M ± 0%   +0.78%  (p=0.000 n=10+10)
```

queue vs [eapache](https://github.com/eapache/queue) - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillQueuev1.0.0.txt testdata/BenchmarkFillEapache.txt
name        old time/op    new time/op    delta
/0-4          35.7ns ± 1%   147.2ns ± 2%  +312.34%  (p=0.000 n=9+8)
/1-4           159ns ± 0%     182ns ± 6%   +14.34%  (p=0.000 n=9+10)
/10-4          549ns ± 1%     482ns ± 1%   -12.21%  (p=0.000 n=9+10)
/100-4        4.55µs ± 2%    5.17µs ± 4%   +13.56%  (p=0.000 n=10+10)
/1000-4       35.0µs ± 1%    43.4µs ± 1%   +23.85%  (p=0.000 n=10+9)
/10000-4       350µs ± 1%     494µs ± 1%   +41.32%  (p=0.000 n=10+10)
/100000-4     3.69ms ± 1%    6.77ms ± 2%   +83.49%  (p=0.000 n=9+10)
/1000000-4    42.5ms ± 1%    73.0ms ± 3%   +71.82%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%    304.0B ± 0%  +533.33%  (p=0.000 n=10+10)
/1-4            160B ± 0%      320B ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           560B ± 0%      464B ± 0%   -17.14%  (p=0.000 n=10+10)
/100-4        7.15kB ± 0%    7.28kB ± 0%    +1.79%  (p=0.000 n=10+10)
/1000-4       33.9kB ± 0%    64.7kB ± 0%   +90.62%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     946kB ± 0%  +193.40%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +145.41%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%    66.3MB ± 0%  +106.48%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      12.0 ± 0%   -14.29%  (p=0.000 n=10+10)
/100-4           107 ± 0%       108 ± 0%    +0.93%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    +0.10%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.60%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

queue vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/queue/queue.go) - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillQueuev1.0.0.txt testdata/BenchmarkFillCookiejar.txt
name        old time/op    new time/op     delta
/0-4          35.7ns ± 1%  10784.8ns ± 3%   +30099.98%  (p=0.000 n=9+8)
/1-4           159ns ± 0%     9871ns ± 5%    +6108.11%  (p=0.000 n=9+9)
/10-4          549ns ± 1%    10897ns ±10%    +1886.16%  (p=0.000 n=9+10)
/100-4        4.55µs ± 2%    12.86µs ± 1%     +182.63%  (p=0.000 n=10+10)
/1000-4       35.0µs ± 1%     40.2µs ± 1%      +14.72%  (p=0.000 n=10+10)
/10000-4       350µs ± 1%      333µs ± 1%       -4.85%  (p=0.000 n=10+10)
/100000-4     3.69ms ± 1%     3.55ms ± 1%       -3.67%  (p=0.000 n=9+10)
/1000000-4    42.5ms ± 1%     44.8ms ± 3%       +5.44%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op    delta
/0-4           48.0B ± 0%   65680.0B ± 0%  +136733.33%  (p=0.000 n=10+10)
/1-4            160B ± 0%     65696B ± 0%   +40960.00%  (p=0.000 n=10+10)
/10-4           560B ± 0%     65840B ± 0%   +11657.14%  (p=0.000 n=10+10)
/100-4        7.15kB ± 0%    67.28kB ± 0%     +840.72%  (p=0.000 n=10+10)
/1000-4       33.9kB ± 0%     81.7kB ± 0%     +140.69%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%      357kB ± 0%      +10.69%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%     3.25MB ± 0%       +0.97%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%     32.8MB ± 0%       +2.19%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op   delta
/0-4            1.00 ± 0%       3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%       4.00 ± 0%         ~     (all equal)
/10-4           14.0 ± 0%       13.0 ± 0%       -7.14%  (p=0.000 n=10+10)
/100-4           107 ± 0%        103 ± 0%       -3.74%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%      1.00k ± 0%       -0.99%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%      10.0k ± 0%       -0.75%  (p=0.000 n=10+10)
/100000-4       101k ± 0%       100k ± 0%       -0.73%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%      1.00M ± 0%       -0.73%  (p=0.000 n=10+10)
```

queue vs [deque](https://github.com/ef-ds/deque) - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillQueuev1.0.0.txt testdata/BenchmarkFillDeque.txt
name        old time/op    new time/op    delta
/0-4          35.7ns ± 1%    38.0ns ± 3%   +6.33%  (p=0.000 n=9+10)
/1-4           159ns ± 0%     179ns ± 7%  +12.52%  (p=0.000 n=9+10)
/10-4          549ns ± 1%     589ns ± 1%   +7.35%  (p=0.000 n=9+10)
/100-4        4.55µs ± 2%    4.78µs ± 0%   +5.04%  (p=0.000 n=10+9)
/1000-4       35.0µs ± 1%    37.2µs ± 0%   +6.21%  (p=0.000 n=10+8)
/10000-4       350µs ± 1%     370µs ± 1%   +5.79%  (p=0.000 n=10+9)
/100000-4     3.69ms ± 1%    3.89ms ± 1%   +5.39%  (p=0.000 n=9+10)
/1000000-4    42.5ms ± 1%    44.6ms ± 2%   +4.88%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     64.0B ± 0%  +33.33%  (p=0.000 n=10+10)
/1-4            160B ± 0%      192B ± 0%  +20.00%  (p=0.000 n=10+10)
/10-4           560B ± 0%      592B ± 0%   +5.71%  (p=0.000 n=10+10)
/100-4        7.15kB ± 0%    7.20kB ± 0%   +0.67%  (p=0.000 n=10+10)
/1000-4       33.9kB ± 0%    34.0kB ± 0%   +0.28%  (p=0.000 n=10+10)
/10000-4       322kB ± 0%     323kB ± 0%   +0.20%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.22MB ± 0%   +0.20%  (p=0.000 n=10+10)
/1000000-4    32.1MB ± 0%    32.2MB ± 0%   +0.19%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%     ~     (all equal)
/10-4           14.0 ± 0%      14.0 ± 0%     ~     (all equal)
/100-4           107 ± 0%       107 ± 0%     ~     (all equal)
/1000-4        1.01k ± 0%     1.01k ± 0%     ~     (all equal)
/10000-4       10.1k ± 0%     10.1k ± 0%     ~     (all equal)
/100000-4       101k ± 0%      101k ± 0%     ~     (all equal)
/1000000-4     1.01M ± 0%     1.01M ± 0%     ~     (all equal)
```


### Microservice Test Results
queue vs [list](https://github.com/golang/go/tree/master/src/container/list) - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceQueuev1.0.0.txt testdata/BenchmarkMicroserviceList.txt
name        old time/op    new time/op    delta
/0-4          36.2ns ± 1%    43.9ns ± 5%   +21.17%  (p=0.000 n=10+9)
/1-4           443ns ± 1%     568ns ±12%   +28.13%  (p=0.000 n=9+10)
/10-4         2.63µs ± 0%    5.27µs ±12%  +100.13%  (p=0.000 n=9+10)
/100-4        22.9µs ± 1%    49.7µs ± 4%  +117.20%  (p=0.000 n=10+10)
/1000-4        209µs ± 1%     492µs ± 3%  +134.79%  (p=0.000 n=10+8)
/10000-4      2.12ms ± 1%    5.35ms ±10%  +153.03%  (p=0.000 n=10+10)
/100000-4     23.2ms ± 0%    82.0ms ±11%  +252.90%  (p=0.000 n=9+10)
/1000000-4     239ms ± 1%     855ms ± 7%  +257.62%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     48.0B ± 0%      ~     (all equal)
/1-4            512B ± 0%      496B ± 0%    -3.12%  (p=0.000 n=10+10)
/10-4         2.54kB ± 0%    4.53kB ± 0%   +77.99%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    44.8kB ± 0%  +114.79%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     448kB ± 0%  +234.20%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    4.48MB ± 0%  +248.22%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    44.8MB ± 0%  +249.47%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     448MB ± 0%  +249.64%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            11.0 ± 0%      15.0 ± 0%   +36.36%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%     141.0 ± 0%   +88.00%  (p=0.000 n=10+10)
/100-4           709 ± 0%      1401 ± 0%   +97.60%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%    14.00k ± 0%   +99.59%  (p=0.000 n=10+10)
/10000-4       70.1k ± 0%    140.0k ± 0%   +99.76%  (p=0.000 n=10+10)
/100000-4       701k ± 0%     1400k ± 0%   +99.77%  (p=0.000 n=10+10)
/1000000-4     7.01M ± 0%    14.00M ± 0%   +99.78%  (p=0.000 n=10+10)
```

queue vs [CustomSliceQueue](testdata_test.go) - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceQueuev1.0.0.txt testdata/BenchmarkMicroserviceSlice.txt
name        old time/op    new time/op    delta
/0-4          36.2ns ± 1%    38.5ns ± 2%    +6.42%  (p=0.000 n=10+9)
/1-4           443ns ± 1%     432ns ± 1%    -2.41%  (p=0.000 n=9+9)
/10-4         2.63µs ± 0%    3.25µs ± 1%   +23.57%  (p=0.000 n=9+9)
/100-4        22.9µs ± 1%    24.1µs ± 1%    +5.43%  (p=0.000 n=10+10)
/1000-4        209µs ± 1%     250µs ± 3%   +19.53%  (p=0.000 n=10+10)
/10000-4      2.12ms ± 1%    2.81ms ± 2%   +32.72%  (p=0.000 n=10+9)
/100000-4     23.2ms ± 0%    43.7ms ± 3%   +88.08%  (p=0.000 n=9+8)
/1000000-4     239ms ± 1%     548ms ± 6%  +129.06%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     32.0B ± 0%   -33.33%  (p=0.000 n=10+10)
/1-4            512B ± 0%      232B ± 0%   -54.69%  (p=0.000 n=10+10)
/10-4         2.54kB ± 0%    2.17kB ± 0%   -14.78%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    21.3kB ± 0%    +2.18%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     214kB ± 0%   +59.98%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    2.95MB ± 0%  +129.11%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    33.1MB ± 0%  +158.50%  (p=0.000 n=10+8)
/1000000-4     128MB ± 0%     338MB ± 0%  +163.68%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            11.0 ± 0%      14.0 ± 0%   +27.27%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%     101.0 ± 0%   +34.67%  (p=0.000 n=10+10)
/100-4           709 ± 0%       822 ± 0%   +15.94%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     8.77k ± 0%   +24.96%  (p=0.000 n=10+10)
/10000-4       70.1k ± 0%     87.8k ± 0%   +25.33%  (p=0.000 n=10+10)
/100000-4       701k ± 0%      875k ± 0%   +24.93%  (p=0.000 n=10+10)
/1000000-4     7.01M ± 0%     8.86M ± 0%   +26.46%  (p=0.000 n=10+10)
```

queue vs [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go) - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceQueuev1.0.0.txt testdata/BenchmarkMicroserviceImpl7.txt
name        old time/op    new time/op    delta
/0-4          36.2ns ± 1%    36.2ns ± 1%      ~     (p=0.614 n=10+10)
/1-4           443ns ± 1%     612ns ± 0%   +38.15%  (p=0.000 n=9+8)
/10-4         2.63µs ± 0%    4.60µs ± 1%   +74.72%  (p=0.000 n=9+9)
/100-4        22.9µs ± 1%    30.7µs ± 2%   +34.41%  (p=0.000 n=10+10)
/1000-4        209µs ± 1%     290µs ± 0%   +38.60%  (p=0.000 n=10+8)
/10000-4      2.12ms ± 1%    2.93ms ± 1%   +38.24%  (p=0.000 n=10+10)
/100000-4     23.2ms ± 0%    31.7ms ± 1%   +36.45%  (p=0.000 n=9+10)
/1000000-4     239ms ± 1%     334ms ± 2%   +39.59%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     48.0B ± 0%      ~     (all equal)
/1-4            512B ± 0%      432B ± 0%   -15.62%  (p=0.000 n=10+10)
/10-4         2.54kB ± 0%    6.91kB ± 0%  +171.70%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    29.6kB ± 0%   +41.92%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     261kB ± 0%   +94.68%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    2.58MB ± 0%  +100.44%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    25.8MB ± 0%  +100.88%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     258MB ± 0%  +100.96%  (p=0.000 n=10+8)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            11.0 ± 0%      17.0 ± 0%   +54.55%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%     109.0 ± 0%   +45.33%  (p=0.000 n=10+10)
/100-4           709 ± 0%       927 ± 0%   +30.75%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     9.11k ± 0%   +29.88%  (p=0.000 n=10+10)
/10000-4       70.1k ± 0%     91.0k ± 0%   +29.78%  (p=0.000 n=10+10)
/100000-4       701k ± 0%      909k ± 0%   +29.77%  (p=0.000 n=10+10)
/1000000-4     7.01M ± 0%     9.09M ± 0%   +29.77%  (p=0.000 n=10+10)
```

queue vs [eapache](https://github.com/eapache/queue) - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceQueuev1.0.0.txt testdata/BenchmarkMicroserviceEapache.txt
name        old time/op    new time/op    delta
/0-4          36.2ns ± 1%   132.3ns ± 1%  +265.27%  (p=0.000 n=10+10)
/1-4           443ns ± 1%     351ns ± 3%   -20.74%  (p=0.000 n=9+10)
/10-4         2.63µs ± 0%    2.38µs ± 0%    -9.61%  (p=0.000 n=9+10)
/100-4        22.9µs ± 1%    25.3µs ± 1%   +10.86%  (p=0.000 n=10+9)
/1000-4        209µs ± 1%     239µs ± 1%   +14.33%  (p=0.000 n=10+10)
/10000-4      2.12ms ± 1%    2.54ms ± 1%   +20.00%  (p=0.000 n=10+10)
/100000-4     23.2ms ± 0%    30.7ms ± 1%   +31.97%  (p=0.000 n=9+10)
/1000000-4     239ms ± 1%     324ms ± 5%   +35.55%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%    304.0B ± 0%  +533.33%  (p=0.000 n=10+10)
/1-4            512B ± 0%      416B ± 0%   -18.75%  (p=0.000 n=10+10)
/10-4         2.54kB ± 0%    1.42kB ± 0%   -44.03%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    22.3kB ± 0%    +6.59%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     209kB ± 0%   +55.95%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    2.69MB ± 0%  +109.21%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    23.8MB ± 0%   +85.51%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     213MB ± 0%   +65.97%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            11.0 ± 0%       9.0 ± 0%   -18.18%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%      72.0 ± 0%    -4.00%  (p=0.000 n=10+10)
/100-4           709 ± 0%       714 ± 0%    +0.71%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%    +0.16%  (p=0.000 n=10+10)
/10000-4       70.1k ± 0%     70.0k ± 0%    -0.06%  (p=0.000 n=10+10)
/100000-4       701k ± 0%      700k ± 0%    -0.10%  (p=0.000 n=10+10)
/1000000-4     7.01M ± 0%     7.00M ± 0%    -0.11%  (p=0.000 n=10+10)
```

queue vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/queue/queue.go) - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceQueuev1.0.0.txt testdata/BenchmarkMicroserviceCookiejar.txt
name        old time/op    new time/op    delta
/0-4          36.2ns ± 1%  9879.0ns ± 1%   +27174.99%  (p=0.000 n=10+10)
/1-4           443ns ± 1%   10146ns ± 1%    +2190.34%  (p=0.000 n=9+10)
/10-4         2.63µs ± 0%   12.07µs ± 1%     +358.83%  (p=0.000 n=9+10)
/100-4        22.9µs ± 1%    29.4µs ± 1%      +28.48%  (p=0.000 n=10+10)
/1000-4        209µs ± 1%     212µs ± 1%       +1.35%  (p=0.000 n=10+9)
/10000-4      2.12ms ± 1%    1.99ms ± 1%       -5.83%  (p=0.000 n=10+9)
/100000-4     23.2ms ± 0%    22.6ms ± 3%       -2.50%  (p=0.000 n=9+10)
/1000000-4     239ms ± 1%     238ms ± 1%         ~     (p=0.182 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%  65680.0B ± 0%  +136733.33%  (p=0.000 n=10+10)
/1-4            512B ± 0%    65792B ± 0%   +12750.00%  (p=0.000 n=10+10)
/10-4         2.54kB ± 0%   66.80kB ± 0%    +2525.79%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    76.9kB ± 0%     +268.20%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     243kB ± 0%      +81.45%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    1.38MB ± 0%       +7.46%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    12.9MB ± 0%       +0.73%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     129MB ± 0%       +0.60%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            11.0 ± 0%      10.0 ± 0%       -9.09%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%      73.0 ± 0%       -2.67%  (p=0.000 n=10+10)
/100-4           709 ± 0%       703 ± 0%       -0.85%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.00k ± 0%       -0.14%  (p=0.000 n=10+10)
/10000-4       70.1k ± 0%     70.0k ± 0%       -0.11%  (p=0.000 n=10+10)
/100000-4       701k ± 0%      700k ± 0%       -0.11%  (p=0.000 n=10+10)
/1000000-4     7.01M ± 0%     7.00M ± 0%       -0.10%  (p=0.000 n=10+10)
```

queue vs [deque](https://github.com/ef-ds/deque) - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceQueuev1.0.0.txt testdata/BenchmarkMicroserviceDeque.txt
name        old time/op    new time/op    delta
/0-4          36.2ns ± 1%    38.2ns ± 3%   +5.44%  (p=0.000 n=10+10)
/1-4           443ns ± 1%     466ns ± 1%   +5.28%  (p=0.000 n=9+10)
/10-4         2.63µs ± 0%    2.79µs ± 2%   +5.89%  (p=0.000 n=9+9)
/100-4        22.9µs ± 1%    24.3µs ± 0%   +6.10%  (p=0.000 n=10+10)
/1000-4        209µs ± 1%     223µs ± 0%   +6.52%  (p=0.000 n=10+8)
/10000-4      2.12ms ± 1%    2.27ms ± 0%   +7.48%  (p=0.000 n=10+9)
/100000-4     23.2ms ± 0%    24.6ms ± 1%   +6.01%  (p=0.000 n=9+10)
/1000000-4     239ms ± 1%     258ms ± 1%   +7.98%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           48.0B ± 0%     64.0B ± 0%  +33.33%  (p=0.000 n=10+10)
/1-4            512B ± 0%      544B ± 0%   +6.25%  (p=0.000 n=10+10)
/10-4         2.54kB ± 0%    2.58kB ± 0%   +1.26%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    20.9kB ± 0%   +0.31%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     134kB ± 0%   +0.08%  (p=0.000 n=10+10)
/10000-4      1.29MB ± 0%    1.43MB ± 0%  +11.33%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    14.4MB ± 0%  +12.53%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     144MB ± 0%  +12.67%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            11.0 ± 0%      11.0 ± 0%     ~     (all equal)
/10-4           75.0 ± 0%      75.0 ± 0%     ~     (all equal)
/100-4           709 ± 0%       709 ± 0%     ~     (all equal)
/1000-4        7.01k ± 0%     7.01k ± 0%     ~     (all equal)
/10000-4       70.1k ± 0%     70.2k ± 0%   +0.10%  (p=0.000 n=10+10)
/100000-4       701k ± 0%      702k ± 0%   +0.11%  (p=0.000 n=10+10)
/1000000-4     7.01M ± 0%     7.02M ± 0%   +0.11%  (p=0.000 n=10+10)
```


### Other Test Results
#### queue vs [list](https://github.com/golang/go/tree/master/src/container/list)
queue vs list - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillQueuev1.0.0.txt testdata/BenchmarkRefillList.txt
name       old time/op    new time/op    delta
/1-4         3.38µs ± 1%    7.57µs ± 3%  +124.08%  (p=0.000 n=9+10)
/10-4        32.3µs ± 2%    71.1µs ± 6%  +119.88%  (p=0.000 n=10+10)
/100-4        312µs ± 1%     718µs ± 6%  +129.93%  (p=0.000 n=10+10)
/1000-4      3.17ms ± 2%    7.29ms ± 9%  +129.80%  (p=0.000 n=10+10)
/10000-4     31.7ms ± 2%    72.3ms ± 5%  +128.37%  (p=0.000 n=10+10)
/100000-4     351ms ± 1%    1991ms ±14%  +466.58%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +299.99%  (p=0.000 n=10+8)
/10000-4     16.0MB ± 0%    64.0MB ± 0%  +299.92%  (p=0.000 n=9+9)
/100000-4     161MB ± 0%     640MB ± 0%  +298.65%  (p=0.000 n=9+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     20.0M ± 0%   +99.99%  (p=0.000 n=10+9)
```

queue vs list - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullQueuev1.0.0.txt testdata/BenchmarkRefillFullList.txt
name       old time/op    new time/op    delta
/1-4         3.31µs ± 2%    9.08µs ± 5%  +173.76%  (p=0.000 n=10+10)
/10-4        33.4µs ± 2%    93.0µs ±11%  +178.94%  (p=0.000 n=10+9)
/100-4        324µs ± 2%     890µs ± 9%  +174.91%  (p=0.000 n=10+9)
/1000-4      3.20ms ± 2%   10.11ms ±12%  +215.98%  (p=0.000 n=10+10)
/10000-4     31.9ms ± 2%    93.2ms ± 3%  +192.61%  (p=0.000 n=10+9)
/100000-4     348ms ± 1%    1738ms ± 7%  +398.79%  (p=0.000 n=9+8)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=9+9)
/100000-4     160MB ± 0%     640MB ± 0%  +300.00%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     20.0M ± 0%  +100.00%  (p=0.000 n=10+10)
```

queue vs list - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseQueuev1.0.0.txt testdata/BenchmarkSlowIncreaseList.txt
name        old time/op    new time/op    delta
/1-4           189ns ± 0%     172ns ± 1%    -8.78%  (p=0.000 n=8+10)
/10-4         1.12µs ± 0%    1.38µs ± 1%   +23.34%  (p=0.000 n=9+10)
/100-4        7.49µs ± 0%   13.30µs ± 0%   +77.56%  (p=0.000 n=10+8)
/1000-4       65.6µs ± 1%   135.6µs ± 2%  +106.69%  (p=0.000 n=10+10)
/10000-4       648µs ± 1%    1413µs ± 1%  +117.88%  (p=0.000 n=9+9)
/100000-4     7.41ms ± 0%   23.89ms ± 6%  +222.55%  (p=0.000 n=8+10)
/1000000-4    78.6ms ± 1%   276.0ms ±13%  +250.89%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      176B ± 0%      ~     (all equal)
/10-4         1.74kB ± 0%    1.33kB ± 0%   -23.85%  (p=0.000 n=10+10)
/100-4        8.75kB ± 0%   12.85kB ± 0%   +46.80%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%   128.0kB ± 0%  +136.85%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%    1280kB ± 0%  +163.09%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%   12.80MB ± 0%  +165.81%  (p=0.000 n=10+9)
/1000000-4    48.1MB ± 0%   128.0MB ± 0%  +165.95%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%      ~     (all equal)
/10-4           25.0 ± 0%      41.0 ± 0%   +64.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       401 ± 0%   +93.72%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     4.00k ± 0%   +98.56%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     40.0k ± 0%   +99.16%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      400k ± 0%   +99.22%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     4.00M ± 0%   +99.22%  (p=0.000 n=10+10)
```

queue vs list - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseQueuev1.0.0.txt testdata/BenchmarkSlowDecreaseList.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 0%    67.5ns ± 1%   +98.94%  (p=0.000 n=8+10)
/10-4          343ns ± 1%     679ns ± 1%   +97.84%  (p=0.000 n=9+10)
/100-4        3.35µs ± 1%    6.68µs ± 1%   +99.42%  (p=0.000 n=10+10)
/1000-4       33.5µs ± 1%    67.1µs ± 1%  +100.15%  (p=0.000 n=10+9)
/10000-4       334µs ± 1%     670µs ± 1%  +100.48%  (p=0.000 n=9+10)
/100000-4     3.33ms ± 1%    6.65ms ± 1%   +99.62%  (p=0.000 n=9+10)
/1000000-4    33.4ms ± 1%    66.2ms ± 0%   +98.24%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

queue vs list - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableQueuev1.0.0.txt testdata/BenchmarkStableList.txt
name        old time/op    new time/op    delta
/1-4          31.8ns ± 2%    84.1ns ± 9%  +164.83%  (p=0.000 n=10+9)
/10-4          318ns ± 1%     872ns ±16%  +173.94%  (p=0.000 n=8+10)
/100-4        3.11µs ± 1%    8.37µs ± 9%  +169.44%  (p=0.000 n=10+10)
/1000-4       31.1µs ± 1%    86.0µs ± 8%  +176.45%  (p=0.000 n=10+10)
/10000-4       311µs ± 1%     846µs ± 5%  +172.14%  (p=0.000 n=10+10)
/100000-4     3.11ms ± 2%    8.77ms ±11%  +182.28%  (p=0.000 n=9+10)
/1000000-4    31.1ms ± 2%    86.8ms ±12%  +178.82%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

#### queue vs [CustomSliceQueue](https://github.com/ef-ds/queue-bench-tests/blob/master/testdata.go)
queue vs CustomSliceQueue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillQueuev1.0.0.txt testdata/BenchmarkRefillSlice.txt
name       old time/op    new time/op    delta
/1-4         3.38µs ± 1%    5.55µs ± 1%   +64.27%  (p=0.000 n=9+9)
/10-4        32.3µs ± 2%    43.3µs ± 0%   +34.07%  (p=0.000 n=10+8)
/100-4        312µs ± 1%     320µs ± 2%    +2.49%  (p=0.000 n=10+10)
/1000-4      3.17ms ± 2%    3.04ms ± 1%    -4.08%  (p=0.000 n=10+10)
/10000-4     31.7ms ± 2%    37.1ms ± 1%   +17.10%  (p=0.000 n=10+9)
/100000-4     351ms ± 1%     797ms ± 2%  +126.86%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    2.40kB ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    32.0kB ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     314kB ± 0%   +96.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    3.15MB ± 0%   +96.80%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    48.8MB ± 0%  +204.74%  (p=0.000 n=9+9)
/100000-4     161MB ± 0%     548MB ± 0%  +241.28%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.1k ± 0%    +1.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      100k ± 0%    +0.20%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%    +0.06%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%    +0.01%  (p=0.000 n=10+10)
```

queue vs CustomSliceQueue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullQueuev1.0.0.txt testdata/BenchmarkRefillFullSlice.txt
name       old time/op    new time/op    delta
/1-4         3.31µs ± 2%    3.82µs ± 1%   +15.11%  (p=0.000 n=10+8)
/10-4        33.4µs ± 2%    37.9µs ± 1%   +13.62%  (p=0.000 n=10+10)
/100-4        324µs ± 2%     371µs ± 1%   +14.62%  (p=0.000 n=10+10)
/1000-4      3.20ms ± 2%    3.91ms ± 1%   +22.03%  (p=0.000 n=10+9)
/10000-4     31.9ms ± 2%    39.9ms ± 1%   +25.14%  (p=0.000 n=10+10)
/100000-4     348ms ± 1%     830ms ± 3%  +138.18%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    4.82kB ± 0%  +200.94%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    48.2kB ± 0%  +201.09%  (p=0.000 n=10+8)
/100-4        160kB ± 0%     483kB ± 0%  +201.70%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    5.15MB ± 0%  +221.87%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    52.0MB ± 0%  +225.28%  (p=0.000 n=9+8)
/100000-4     160MB ± 0%     551MB ± 0%  +244.43%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    +0.03%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      100k ± 0%    +0.03%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%    +0.03%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%    +0.01%  (p=0.000 n=10+10)
```

queue vs CustomSliceQueue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseQueuev1.0.0.txt testdata/BenchmarkSlowIncreaseSlice.txt
name        old time/op    new time/op    delta
/1-4           189ns ± 0%     152ns ± 1%   -19.47%  (p=0.000 n=8+10)
/10-4         1.12µs ± 0%    1.06µs ± 3%    -5.22%  (p=0.000 n=9+10)
/100-4        7.49µs ± 0%    6.92µs ± 1%    -7.60%  (p=0.000 n=10+10)
/1000-4       65.6µs ± 1%    62.9µs ± 1%    -4.11%  (p=0.000 n=10+9)
/10000-4       648µs ± 1%     727µs ± 1%   +12.15%  (p=0.000 n=9+10)
/100000-4     7.41ms ± 0%   13.43ms ± 2%   +81.37%  (p=0.000 n=8+9)
/1000000-4    78.6ms ± 1%   166.2ms ± 6%  +111.27%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%       88B ± 0%   -50.00%  (p=0.000 n=10+10)
/10-4         1.74kB ± 0%    0.78kB ± 0%   -55.50%  (p=0.000 n=10+10)
/100-4        8.75kB ± 0%    6.66kB ± 0%   -23.86%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%    74.0kB ± 0%   +36.86%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%     991kB ± 0%  +103.60%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%   11.42MB ± 0%  +137.06%  (p=0.000 n=10+10)
/1000000-4    48.1MB ± 0%   114.6MB ± 0%  +138.05%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%      ~     (all equal)
/10-4           25.0 ± 0%      29.0 ± 0%   +16.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       214 ± 0%    +3.38%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.02k ± 0%    +0.25%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.25%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.37%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

queue vs CustomSliceQueue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseQueuev1.0.0.txt testdata/BenchmarkSlowDecreaseSlice.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 0%    55.7ns ± 0%   +64.19%  (p=0.000 n=8+8)
/10-4          343ns ± 1%     562ns ± 1%   +63.82%  (p=0.000 n=9+9)
/100-4        3.35µs ± 1%    5.54µs ± 1%   +65.22%  (p=0.000 n=10+10)
/1000-4       33.5µs ± 1%    55.2µs ± 1%   +64.75%  (p=0.000 n=10+8)
/10000-4       334µs ± 1%     554µs ± 1%   +65.77%  (p=0.000 n=9+10)
/100000-4     3.33ms ± 1%    5.55ms ± 1%   +66.41%  (p=0.000 n=9+10)
/1000000-4    33.4ms ± 1%    55.6ms ± 2%   +66.62%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     24.0B ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      240B ± 0%   +50.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    2.40kB ± 0%   +50.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    24.0kB ± 0%   +50.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     240kB ± 0%   +50.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    2.40MB ± 0%   +50.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    24.0MB ± 0%   +50.00%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

queue vs CustomSliceQueue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableQueuev1.0.0.txt testdata/BenchmarkStableSlice.txt
name        old time/op    new time/op    delta
/1-4          31.8ns ± 2%    36.3ns ± 1%   +14.40%  (p=0.000 n=10+9)
/10-4          318ns ± 1%     366ns ± 1%   +14.79%  (p=0.000 n=8+10)
/100-4        3.11µs ± 1%    3.57µs ± 1%   +15.01%  (p=0.000 n=10+10)
/1000-4       31.1µs ± 1%    35.6µs ± 0%   +14.58%  (p=0.000 n=10+9)
/10000-4       311µs ± 1%     366µs ± 3%   +17.80%  (p=0.000 n=10+10)
/100000-4     3.11ms ± 2%    3.68ms ± 1%   +18.44%  (p=0.000 n=9+8)
/1000000-4    31.1ms ± 2%    36.8ms ± 1%   +18.08%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     48.0B ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      481B ± 0%  +200.62%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    4.82kB ± 0%  +200.94%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    48.2kB ± 0%  +200.97%  (p=0.000 n=10+8)
/10000-4       160kB ± 0%     482kB ± 0%  +200.97%  (p=0.000 n=10+8)
/100000-4     1.60MB ± 0%    4.82MB ± 0%  +200.97%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    48.2MB ± 0%  +200.97%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%      ~     (all equal)
/100-4           100 ± 0%       100 ± 0%      ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    +0.03%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      100k ± 0%    +0.03%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.00M ± 0%    +0.03%  (p=0.000 n=10+10)
```

#### queue vs [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go)
queue vs impl7 - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillQueuev1.0.0.txt testdata/BenchmarkRefillImpl7.txt
name       old time/op    new time/op    delta
/1-4         3.38µs ± 1%    9.33µs ± 1%  +175.98%  (p=0.000 n=9+10)
/10-4        32.3µs ± 2%    69.4µs ± 1%  +114.57%  (p=0.000 n=10+10)
/100-4        312µs ± 1%     409µs ± 0%   +30.97%  (p=0.000 n=10+8)
/1000-4      3.17ms ± 2%    3.63ms ± 1%   +14.34%  (p=0.000 n=10+10)
/10000-4     31.7ms ± 2%    36.2ms ± 1%   +14.15%  (p=0.000 n=10+9)
/100000-4     351ms ± 1%     380ms ± 1%    +8.27%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    68.8kB ± 0%  +330.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     421kB ± 0%  +163.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    3.32MB ± 0%  +107.29%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    32.3MB ± 0%  +101.69%  (p=0.000 n=9+10)
/100000-4     161MB ± 0%     323MB ± 0%  +101.01%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.60k ± 0%   +60.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      102k ± 0%    +2.20%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.02M ± 0%    +1.62%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.2M ± 0%    +1.57%  (p=0.000 n=10+10)
```

queue vs impl7 - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullQueuev1.0.0.txt testdata/BenchmarkRefillFullImpl7.txt
name       old time/op    new time/op    delta
/1-4         3.31µs ± 2%    3.77µs ± 1%   +13.63%  (p=0.000 n=10+10)
/10-4        33.4µs ± 2%    37.8µs ± 1%   +13.38%  (p=0.000 n=10+10)
/100-4        324µs ± 2%     366µs ± 1%   +13.13%  (p=0.000 n=10+10)
/1000-4      3.20ms ± 2%    3.66ms ± 2%   +14.39%  (p=0.000 n=10+10)
/10000-4     31.9ms ± 2%    36.9ms ± 1%   +15.82%  (p=0.000 n=10+10)
/100000-4     348ms ± 1%     375ms ± 2%    +7.65%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    3.23kB ± 0%  +101.56%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    32.2kB ± 0%  +101.56%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     322kB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    3.23MB ± 0%  +101.56%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    32.2MB ± 0%  +101.56%  (p=0.000 n=9+9)
/100000-4     160MB ± 0%     322MB ± 0%  +101.56%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       101 ± 0%    +1.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.01k ± 0%    +1.50%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.2k ± 0%    +1.56%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      102k ± 0%    +1.56%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.02M ± 0%    +1.56%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.2M ± 0%    +1.56%  (p=0.000 n=10+10)
```

queue vs impl7 - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseQueuev1.0.0.txt testdata/BenchmarkSlowIncreaseImpl7.txt
name        old time/op    new time/op    delta
/1-4           189ns ± 0%     225ns ± 0%  +18.99%  (p=0.000 n=8+10)
/10-4         1.12µs ± 0%    1.55µs ± 1%  +38.12%  (p=0.000 n=9+10)
/100-4        7.49µs ± 0%    7.63µs ± 1%   +1.82%  (p=0.000 n=10+9)
/1000-4       65.6µs ± 1%    70.1µs ± 1%   +6.85%  (p=0.000 n=10+10)
/10000-4       648µs ± 1%     703µs ± 0%   +8.41%  (p=0.000 n=9+9)
/100000-4     7.41ms ± 0%    7.93ms ± 1%   +7.15%  (p=0.000 n=8+9)
/1000000-4    78.6ms ± 1%    88.3ms ± 2%  +12.28%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      160B ± 0%   -9.09%  (p=0.000 n=10+10)
/10-4         1.74kB ± 0%    2.98kB ± 0%  +70.64%  (p=0.000 n=10+10)
/100-4        8.75kB ± 0%    7.94kB ± 0%   -9.32%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%    65.9kB ± 0%  +21.81%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%     647kB ± 0%  +33.01%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    6.45MB ± 0%  +33.98%  (p=0.000 n=10+10)
/1000000-4    48.1MB ± 0%    64.5MB ± 0%  +34.01%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      6.00 ± 0%  +20.00%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      29.0 ± 0%  +16.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       211 ± 0%   +1.93%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.04k ± 0%   +1.19%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.3k ± 0%   +1.18%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      203k ± 0%   +1.17%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.03M ± 0%   +1.17%  (p=0.000 n=10+10)
```

queue vs impl7 - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseQueuev1.0.0.txt testdata/BenchmarkSlowDecreaseImpl7.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 0%    92.4ns ± 7%  +172.39%  (p=0.000 n=8+10)
/10-4          343ns ± 1%     910ns ± 1%  +165.36%  (p=0.000 n=9+10)
/100-4        3.35µs ± 1%    9.01µs ± 1%  +168.84%  (p=0.000 n=10+10)
/1000-4       33.5µs ± 1%    89.9µs ± 1%  +168.16%  (p=0.000 n=10+9)
/10000-4       334µs ± 1%     906µs ± 2%  +171.10%  (p=0.000 n=9+10)
/100000-4     3.33ms ± 1%    8.98ms ± 1%  +169.37%  (p=0.000 n=9+10)
/1000000-4    33.4ms ± 1%    90.1ms ± 1%  +169.84%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=9+8)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      3.00 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      30.0 ± 0%  +200.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     3.00k ± 0%  +200.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     30.0k ± 0%  +200.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      300k ± 0%  +200.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     3.00M ± 0%  +200.00%  (p=0.000 n=10+10)
```

queue vs impl7 - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableQueuev1.0.0.txt testdata/BenchmarkStableImpl7.txt
name        old time/op    new time/op    delta
/1-4          31.8ns ± 2%    36.4ns ± 1%   +14.55%  (p=0.000 n=10+10)
/10-4          318ns ± 1%     368ns ± 0%   +15.47%  (p=0.000 n=8+9)
/100-4        3.11µs ± 1%    3.60µs ± 1%   +15.78%  (p=0.000 n=10+10)
/1000-4       31.1µs ± 1%    35.9µs ± 1%   +15.55%  (p=0.000 n=10+10)
/10000-4       311µs ± 1%     359µs ± 0%   +15.35%  (p=0.000 n=10+10)
/100000-4     3.11ms ± 2%    3.59ms ± 1%   +15.70%  (p=0.000 n=9+9)
/1000000-4    31.1ms ± 2%    35.9ms ± 1%   +15.18%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     32.0B ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      322B ± 0%  +101.25%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    3.23kB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    32.2kB ± 0%  +101.56%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     322kB ± 0%  +101.56%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    3.23MB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    32.2MB ± 0%  +101.56%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%      ~     (all equal)
/100-4           100 ± 0%       101 ± 0%    +1.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     1.01k ± 0%    +1.50%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     10.2k ± 0%    +1.56%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      102k ± 0%    +1.56%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.02M ± 0%    +1.56%  (p=0.000 n=10+10)
```

#### queue vs [eapache](https://github.com/eapache/queue)
queue vs eapache - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillQueuev1.0.0.txt testdata/BenchmarkRefillEapache.txt
name       old time/op    new time/op    delta
/1-4         3.38µs ± 1%    3.65µs ± 2%    +8.10%  (p=0.000 n=9+10)
/10-4        32.3µs ± 2%    35.3µs ± 1%    +9.12%  (p=0.000 n=10+10)
/100-4        312µs ± 1%     499µs ± 1%   +59.72%  (p=0.000 n=10+9)
/1000-4      3.17ms ± 2%    4.37ms ± 1%   +37.83%  (p=0.000 n=10+9)
/10000-4     31.7ms ± 2%    50.0ms ± 1%   +57.80%  (p=0.000 n=10+9)
/100000-4     351ms ± 1%     686ms ± 2%   +95.32%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     698kB ± 0%  +336.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.44MB ± 0%  +302.39%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    94.6MB ± 0%  +490.92%  (p=0.000 n=9+8)
/100000-4     161MB ± 0%     789MB ± 0%  +391.51%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.20%  (p=0.000 n=10+10)
/10000-4      1.00M ± 0%     1.00M ± 0%    +0.20%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%    +0.02%  (p=0.000 n=10+10)
```

queue vs eapache - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullQueuev1.0.0.txt testdata/BenchmarkRefillFullEapache.txt
name       old time/op    new time/op    delta
/1-4         3.31µs ± 2%    3.64µs ± 2%    +9.79%  (p=0.000 n=10+9)
/10-4        33.4µs ± 2%    36.8µs ± 2%   +10.34%  (p=0.000 n=10+10)
/100-4        324µs ± 2%     357µs ± 3%   +10.20%  (p=0.000 n=10+10)
/1000-4      3.20ms ± 2%    3.59ms ± 4%   +12.25%  (p=0.000 n=10+10)
/10000-4     31.9ms ± 2%    36.7ms ± 4%   +15.09%  (p=0.000 n=10+10)
/100000-4     348ms ± 1%     671ms ± 1%   +92.42%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%      ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%      ~     (all equal)
/10000-4     16.0MB ± 0%    16.0MB ± 0%      ~     (all equal)
/100000-4     160MB ± 0%     632MB ± 0%  +294.91%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%      ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%      ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%      ~     (all equal)
/100000-4     10.0M ± 0%     10.0M ± 0%    +0.00%  (p=0.000 n=10+10)
```

queue vs eapache - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseQueuev1.0.0.txt testdata/BenchmarkSlowIncreaseEapache.txt
name        old time/op    new time/op    delta
/1-4           189ns ± 0%     204ns ± 0%    +7.94%  (p=0.000 n=8+10)
/10-4         1.12µs ± 0%    0.80µs ± 1%   -28.72%  (p=0.000 n=9+9)
/100-4        7.49µs ± 0%    8.30µs ± 1%   +10.73%  (p=0.000 n=10+10)
/1000-4       65.6µs ± 1%    75.3µs ± 1%   +14.82%  (p=0.000 n=10+10)
/10000-4       648µs ± 1%     879µs ± 1%   +35.50%  (p=0.000 n=9+9)
/100000-4     7.41ms ± 0%   10.79ms ± 1%   +45.73%  (p=0.000 n=8+10)
/1000000-4    78.6ms ± 1%   121.1ms ± 3%   +53.95%  (p=0.000 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      336B ± 0%   +90.91%  (p=0.000 n=10+10)
/10-4         1.74kB ± 0%    0.62kB ± 0%   -64.22%  (p=0.000 n=10+10)
/100-4        8.75kB ± 0%    8.88kB ± 0%    +1.46%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%    80.7kB ± 0%   +49.25%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%    1106kB ± 0%  +127.31%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    9.49MB ± 0%   +97.09%  (p=0.000 n=10+10)
/1000000-4    48.1MB ± 0%    82.3MB ± 0%   +71.06%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      4.00 ± 0%   -20.00%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      22.0 ± 0%   -12.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       208 ± 0%    +0.48%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.01k ± 0%    -0.05%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.31%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

queue vs eapache - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseQueuev1.0.0.txt testdata/BenchmarkSlowDecreaseEapache.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 0%    33.8ns ± 1%    ~     (p=0.160 n=8+9)
/10-4          343ns ± 1%     344ns ± 1%    ~     (p=0.375 n=9+10)
/100-4        3.35µs ± 1%    3.34µs ± 0%  -0.32%  (p=0.021 n=10+8)
/1000-4       33.5µs ± 1%    33.7µs ± 2%    ~     (p=0.280 n=10+10)
/10000-4       334µs ± 1%     334µs ± 1%    ~     (p=0.905 n=9+10)
/100000-4     3.33ms ± 1%    3.36ms ± 1%    ~     (p=0.050 n=9+9)
/1000000-4    33.4ms ± 1%    33.5ms ± 1%    ~     (p=0.105 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

queue vs eapache - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableQueuev1.0.0.txt testdata/BenchmarkStableEapache.txt
name        old time/op    new time/op    delta
/1-4          31.8ns ± 2%    34.3ns ± 2%   +7.93%  (p=0.000 n=10+10)
/10-4          318ns ± 1%     349ns ± 2%   +9.70%  (p=0.000 n=8+10)
/100-4        3.11µs ± 1%    3.44µs ± 2%  +10.71%  (p=0.000 n=10+10)
/1000-4       31.1µs ± 1%    34.3µs ± 3%  +10.47%  (p=0.000 n=10+10)
/10000-4       311µs ± 1%     340µs ± 2%   +9.34%  (p=0.000 n=10+10)
/100000-4     3.11ms ± 2%    3.39ms ± 1%   +9.01%  (p=0.000 n=9+9)
/1000000-4    31.1ms ± 2%    34.2ms ± 3%   +9.84%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

#### queue vs [cookiejar](https://github.com/karalabe/cookiejar/blob/v2/collections/queue/queue.go)
queue vs cookiejar - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillQueuev1.0.0.txt testdata/BenchmarkRefillCookiejar.txt
name       old time/op    new time/op    delta
/1-4         3.38µs ± 1%    3.39µs ± 2%    ~     (p=0.563 n=9+10)
/10-4        32.3µs ± 2%    31.5µs ± 2%  -2.42%  (p=0.000 n=10+9)
/100-4        312µs ± 1%     306µs ± 2%  -1.86%  (p=0.000 n=10+10)
/1000-4      3.17ms ± 2%    3.03ms ± 2%  -4.53%  (p=0.000 n=10+10)
/10000-4     31.7ms ± 2%    31.4ms ± 3%    ~     (p=0.315 n=10+10)
/100000-4     351ms ± 1%     354ms ± 2%    ~     (p=0.113 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%  +0.01%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     160kB ± 0%  +0.02%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%  +0.01%  (p=0.000 n=10+10)
/10000-4     16.0MB ± 0%    16.0MB ± 0%  +0.01%  (p=0.000 n=9+10)
/100000-4     161MB ± 0%     161MB ± 0%  +0.02%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%    ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%  -0.00%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.0M ± 0%  -0.00%  (p=0.000 n=10+10)
```

queue vs cookiejar - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullQueuev1.0.0.txt testdata/BenchmarkRefillFullCookiejar.txt
name       old time/op    new time/op    delta
/1-4         3.31µs ± 2%    3.17µs ± 2%  -4.27%  (p=0.000 n=10+10)
/10-4        33.4µs ± 2%    31.7µs ± 3%  -4.82%  (p=0.000 n=10+10)
/100-4        324µs ± 2%     308µs ± 2%  -5.05%  (p=0.000 n=10+10)
/1000-4      3.20ms ± 2%    3.11ms ± 1%  -2.92%  (p=0.000 n=10+8)
/10000-4     31.9ms ± 2%    31.3ms ± 3%  -1.68%  (p=0.035 n=10+10)
/100000-4     348ms ± 1%     346ms ± 3%  -0.68%  (p=0.040 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%    ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/10000-4     16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)
/100000-4     160MB ± 0%     160MB ± 0%    ~     (all equal)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%    ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%    ~     (all equal)
/10000-4      1.00M ± 0%     1.00M ± 0%    ~     (all equal)
/100000-4     10.0M ± 0%     10.0M ± 0%    ~     (all equal)
```

queue vs cookiejar - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseQueuev1.0.0.txt testdata/BenchmarkSlowIncreaseCookiejar.txt
name        old time/op    new time/op    delta
/1-4           189ns ± 0%    9984ns ± 1%   +5182.36%  (p=0.000 n=8+9)
/10-4         1.12µs ± 0%   10.63µs ± 1%    +847.31%  (p=0.000 n=9+9)
/100-4        7.49µs ± 0%   15.78µs ± 1%    +110.61%  (p=0.000 n=10+10)
/1000-4       65.6µs ± 1%    67.8µs ± 1%      +3.35%  (p=0.000 n=10+10)
/10000-4       648µs ± 1%     637µs ± 1%      -1.76%  (p=0.000 n=9+10)
/100000-4     7.41ms ± 0%    7.87ms ± 3%      +6.32%  (p=0.000 n=8+10)
/1000000-4    78.6ms ± 1%    82.1ms ± 3%      +4.34%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%    65712B ± 0%  +37236.36%  (p=0.000 n=10+10)
/10-4         1.74kB ± 0%   66.00kB ± 0%   +3684.40%  (p=0.000 n=10+10)
/100-4        8.75kB ± 0%   68.88kB ± 0%    +687.02%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%    97.7kB ± 0%     +80.67%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%     583kB ± 0%     +19.72%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.91MB ± 0%      +2.02%  (p=0.000 n=10+10)
/1000000-4    48.1MB ± 0%    48.9MB ± 0%      +1.60%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%        ~     (all equal)
/10-4           25.0 ± 0%      23.0 ± 0%      -8.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       203 ± 0%      -1.93%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.00k ± 0%      -0.60%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%      -0.38%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%      -0.37%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%      -0.36%  (p=0.000 n=10+10)
```

queue vs cookiejar - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseQueuev1.0.0.txt testdata/BenchmarkSlowDecreaseCookiejar.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 0%    33.0ns ± 1%  -2.70%  (p=0.000 n=8+10)
/10-4          343ns ± 1%     336ns ± 1%  -2.16%  (p=0.000 n=9+10)
/100-4        3.35µs ± 1%    3.27µs ± 1%  -2.55%  (p=0.000 n=10+9)
/1000-4       33.5µs ± 1%    33.0µs ± 5%    ~     (p=0.063 n=10+10)
/10000-4       334µs ± 1%     326µs ± 3%  -2.35%  (p=0.001 n=9+10)
/100000-4     3.33ms ± 1%    3.26ms ± 1%  -2.20%  (p=0.000 n=9+10)
/1000000-4    33.4ms ± 1%    32.7ms ± 1%  -2.05%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

queue vs cookiejar - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableQueuev1.0.0.txt testdata/BenchmarkStableCookiejar.txt
name        old time/op    new time/op    delta
/1-4          31.8ns ± 2%    31.7ns ± 6%    ~     (p=0.782 n=10+10)
/10-4          318ns ± 1%     332ns ± 4%  +4.08%  (p=0.000 n=8+10)
/100-4        3.11µs ± 1%    3.16µs ± 4%    ~     (p=0.148 n=10+10)
/1000-4       31.1µs ± 1%    31.3µs ± 5%    ~     (p=1.000 n=10+10)
/10000-4       311µs ± 1%     318µs ± 5%    ~     (p=0.243 n=10+9)
/100000-4     3.11ms ± 2%    3.24ms ± 6%  +4.33%  (p=0.003 n=9+10)
/1000000-4    31.1ms ± 2%    31.3ms ± 3%    ~     (p=0.408 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

#### queue vs [deque](https://github.com/ef-ds/deque)
queue vs Deque - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillQueuev1.0.0.txt testdata/BenchmarkRefillDeque.txt
name       old time/op    new time/op    delta
/1-4         3.38µs ± 1%    3.59µs ± 1%   +6.11%  (p=0.000 n=9+10)
/10-4        32.3µs ± 2%    35.0µs ± 1%   +8.36%  (p=0.000 n=10+10)
/100-4        312µs ± 1%     338µs ± 1%   +8.31%  (p=0.000 n=10+10)
/1000-4      3.17ms ± 2%    3.39ms ± 1%   +6.99%  (p=0.000 n=10+10)
/10000-4     31.7ms ± 2%    36.7ms ± 1%  +15.99%  (p=0.000 n=10+9)
/100000-4     351ms ± 1%     394ms ± 0%  +11.99%  (p=0.000 n=10+8)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     16.0MB ± 0%    30.1MB ± 0%  +88.19%  (p=0.000 n=9+9)
/100000-4     161MB ± 0%     320MB ± 0%  +99.21%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.00M ± 0%     1.01M ± 0%   +0.68%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.1M ± 0%   +0.77%  (p=0.000 n=10+10)
```

queue vs Deque - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullQueuev1.0.0.txt testdata/BenchmarkRefillFullDeque.txt
name       old time/op    new time/op    delta
/1-4         3.31µs ± 2%    3.53µs ± 2%   +6.54%  (p=0.000 n=10+10)
/10-4        33.4µs ± 2%    35.3µs ± 1%   +5.98%  (p=0.000 n=10+10)
/100-4        324µs ± 2%     345µs ± 1%   +6.38%  (p=0.000 n=10+10)
/1000-4      3.20ms ± 2%    3.43ms ± 1%   +7.05%  (p=0.000 n=10+10)
/10000-4     31.9ms ± 2%    38.1ms ± 5%  +19.73%  (p=0.000 n=10+10)
/100000-4     348ms ± 1%     392ms ± 1%  +12.49%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     16.0MB ± 0%    30.1MB ± 0%  +88.22%  (p=0.000 n=9+10)
/100000-4     160MB ± 0%     320MB ± 0%  +99.88%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.00M ± 0%     1.01M ± 0%   +0.68%  (p=0.000 n=10+10)
/100000-4     10.0M ± 0%     10.1M ± 0%   +0.77%  (p=0.000 n=10+10)
```

queue vs Deque - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseQueuev1.0.0.txt testdata/BenchmarkSlowIncreaseDeque.txt
name        old time/op    new time/op    delta
/1-4           189ns ± 0%     202ns ± 1%   +6.93%  (p=0.000 n=8+10)
/10-4         1.12µs ± 0%    1.18µs ± 1%   +4.83%  (p=0.000 n=9+9)
/100-4        7.49µs ± 0%    7.98µs ± 0%   +6.47%  (p=0.000 n=10+9)
/1000-4       65.6µs ± 1%    70.3µs ± 0%   +7.07%  (p=0.000 n=10+10)
/10000-4       648µs ± 1%     697µs ± 0%   +7.46%  (p=0.000 n=9+9)
/100000-4     7.41ms ± 0%    7.96ms ± 1%   +7.51%  (p=0.000 n=8+10)
/1000000-4    78.6ms ± 1%    83.6ms ± 1%   +6.35%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      208B ± 0%  +18.18%  (p=0.000 n=10+10)
/10-4         1.74kB ± 0%    1.78kB ± 0%   +1.83%  (p=0.000 n=10+10)
/100-4        8.75kB ± 0%    8.80kB ± 0%   +0.55%  (p=0.000 n=10+10)
/1000-4       54.1kB ± 0%    54.2kB ± 0%   +0.21%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%     487kB ± 0%   +0.14%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.82MB ± 0%   +0.13%  (p=0.000 n=10+10)
/1000000-4    48.1MB ± 0%    48.2MB ± 0%   +0.13%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%     ~     (all equal)
/10-4           25.0 ± 0%      25.0 ± 0%     ~     (all equal)
/100-4           207 ± 0%       207 ± 0%     ~     (all equal)
/1000-4        2.02k ± 0%     2.02k ± 0%     ~     (all equal)
/10000-4       20.1k ± 0%     20.1k ± 0%     ~     (all equal)
/100000-4       201k ± 0%      201k ± 0%     ~     (all equal)
/1000000-4     2.01M ± 0%     2.01M ± 0%     ~     (all equal)
```

queue vs Deque - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseQueuev1.0.0.txt testdata/BenchmarkSlowDecreaseDeque.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 0%    34.1ns ± 1%  +0.63%  (p=0.001 n=8+10)
/10-4          343ns ± 1%     348ns ± 1%  +1.31%  (p=0.000 n=9+10)
/100-4        3.35µs ± 1%    3.41µs ± 1%  +1.76%  (p=0.000 n=10+9)
/1000-4       33.5µs ± 1%    33.8µs ± 1%  +0.80%  (p=0.003 n=10+10)
/10000-4       334µs ± 1%     338µs ± 1%  +1.06%  (p=0.000 n=9+10)
/100000-4     3.33ms ± 1%    3.38ms ± 1%  +1.46%  (p=0.000 n=9+10)
/1000000-4    33.4ms ± 1%    33.8ms ± 1%  +1.22%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

queue vs Deque - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableQueuev1.0.0.txt testdata/BenchmarkStableDeque.txt
name        old time/op    new time/op    delta
/1-4          31.8ns ± 2%    33.7ns ± 1%  +6.21%  (p=0.000 n=10+9)
/10-4          318ns ± 1%     340ns ± 1%  +6.69%  (p=0.000 n=8+10)
/100-4        3.11µs ± 1%    3.33µs ± 1%  +7.18%  (p=0.000 n=10+10)
/1000-4       31.1µs ± 1%    33.2µs ± 1%  +6.72%  (p=0.000 n=10+10)
/10000-4       311µs ± 1%     335µs ± 3%  +7.64%  (p=0.000 n=10+10)
/100000-4     3.11ms ± 2%    3.41ms ± 6%  +9.77%  (p=0.000 n=9+10)
/1000000-4    31.1ms ± 2%    33.3ms ± 1%  +7.04%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```
