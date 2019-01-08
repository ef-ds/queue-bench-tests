# Benchmark Tests

## Tests
The benchmark tests are composed of all the tests implemented in the [benchmark package](https://github.com/ef-ds/benchmark).


## Tested Queues

Besides queue, the tests also probe a few high quality open source queue implementations, alongside the standard list package as well as using simple slice as a queue.

- List based queue: uses the standard [list](https://github.com/golang/go/tree/master/src/container/list) package as a FIFO queue.
- [CustomSliceQueue](testdata.go): uses a simple, dynamically growing slice as its underlying data structure.
- [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go): experimental queue implementation that stores the values in linked slices. This implementation tests the queue performance when performing lazy creation of the internal slice as well as starting with a 1-sized slice, allowing it to grow up to 16 by using the built in append function. Subsequent slices are created with 128 fixed size.
- [eapache](https://github.com/eapache/queue): this is a ring-buffer slice based queue. The queue uses a minimum queue length of 16 and resizes down if buffer 1/4 full.
- [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/queue/queue.go): the queue implemented in this package uses a  uses a dynamically growing slice of blocks to store the elements.
- [deque](https://github.com/ef-ds/deque): deque stores the elements in a dynamic growing circular doubly linked list of arrays. The queue implemented on this package is actually based on this deque.

We're actively looking for other, high quality queues to add to our tests. Due to the large volume of open source queues available, it is not possible to add all of them to the tests. However, all the new tested ones we're adding to this [issue](https://github.com/ef-ds/queue/issues/1).


### Efficient Data Structures queue vs deque

Efficient Data Structures implements this queue package as well as the [deque](https://github.com/ef-ds/deque) package which can also be used as a FIFO queue.

The queue implementated in this queue package is a simplified version of this deque package. When it comes to using the packages as a FIFO queue, the main differences are:

1) Queue is a simpler version of deque that performs better than deque on most, if not all, FIFO queue tests
2) Differently from deque which keeps only a limited number of unused slices, queue keeps all the previous allocated slices in memory. This has the advantage of less memory allocations, which realizes mostly on refill test scenarios, but also means a potentially larger final data structure size after use.


## Results

The raw results of a local run are stored under the [testdata](testdata) directory.

Refer [here](PERFORMANCE.md) for curated results.


## How To Run

From the package main directory, the tests can be run with below command.

```
go test -benchmem -timeout 60m -bench=. -run=^$
```

To run the test for a single queue, below command can be used.

```
go test -benchmem -timeout 60m -bench="QUEUE_NAME*" -run=^$
```

Replace the QUEUE_NAME with the desired queue such as "List", "Slice", "Impl7", "Eapache", "Cookiejar", "Deque", "Queue".


To run only a specific test suite, below command can be used.

```
go test -benchmem -timeout 60m -bench="TEST_SUITE_NAME*" -run=^$
```

Replace the TEST_SUITE_NAME with the desired test suite such as "Microservice", "Fill", "Refill", "RefillFull", "SlowIncrease", "SlowDecrease", "Stable".


## Test Variations

It is common to see significant variations in the test numbers with different test runs due to different reasons such as processes running in the hosting computer while the tests are running.

It is recommended to run the test multiple times and compare the aggregated results in order to help eliminate/smooth the test variations.

To run the tests multiple times, use the "go test" count parameter as below.

```
go test -benchmem -count 10 -timeout 600m -bench=. -run=^$
```

As the number of tests and now, test runs as well, is very large, it becomes very difficult to analyze and understand the results. In order to be able to analyze and compare the results between the different queues, the [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) tool can be used to aggregate the test results. But as benchstat was designed to compare the same set of tests, it is necessary to first split all the different tests into separate test files renaming each
test with the same name, so benchstat will be able to match the different tests.

First step is to run the test and output the results in a file. Below command can be used to run all tests 10 times.

```
go test -benchmem -count 10 -timeout 600m -bench=. -run=^$ > testdata/results.txt
```

Next step is to split the "results.txt" file into separate test files. The [test-splitter](https://github.com/ef-ds/tools/tree/master/testsplitter) tool can be used for this purpose. To run the tool, clone the repo and run test-splitter from the "testsplitter" directory as follow.

```
go run *.go --file PATH_TO_RESULTS.TXT
```

Test-splitter should output one file per test name in the tests results file. The file names are named after each test name.

The last step is to run the [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) tool to aggregate and compare the results.

Below are the set of benchstat commands that can be used to compare deque against the other tested queues.

Queue vs list
```
benchstat testdata/BenchmarkMicroserviceQueuev1.0.0.txt testdata/BenchmarkMicroserviceList.txt
benchstat testdata/BenchmarkFillQueuev1.0.0.txt testdata/BenchmarkFillList.txt
benchstat testdata/BenchmarkRefillQueuev1.0.0.txt testdata/BenchmarkRefillList.txt
benchstat testdata/BenchmarkRefillFullQueuev1.0.0.txt testdata/BenchmarkRefillFullList.txt
benchstat testdata/BenchmarkSlowIncreaseQueuev1.0.0.txt testdata/BenchmarkSlowIncreaseList.txt
benchstat testdata/BenchmarkSlowDecreaseQueuev1.0.0.txt testdata/BenchmarkSlowDecreaseList.txt
benchstat testdata/BenchmarkStableQueuev1.0.0.txt testdata/BenchmarkStableList.txt
```

Queue vs CustomSliceQueue
```
benchstat testdata/BenchmarkMicroserviceQueuev1.0.0.txt testdata/BenchmarkMicroserviceSlice.txt
benchstat testdata/BenchmarkFillQueuev1.0.0.txt testdata/BenchmarkFillSlice.txt
benchstat testdata/BenchmarkRefillQueuev1.0.0.txt testdata/BenchmarkRefillSlice.txt
benchstat testdata/BenchmarkRefillFullQueuevv1.0.0.txt testdata/BenchmarkRefillFullSlice.txt
benchstat testdata/BenchmarkSlowIncreaseQueuevv1.0.0.txt testdata/BenchmarkSlowIncreaseSlice.txt
benchstat testdata/BenchmarkSlowDecreaseQueuevv1.0.0.txt testdata/BenchmarkSlowDecreaseSlice.txt
benchstat testdata/BenchmarkStableQueuevv1.0.0.txt testdata/BenchmarkStableSlice.txt
```

Queue vs impl7
```
benchstat testdata/BenchmarkMicroserviceQueuev1.0.0.txt testdata/BenchmarkMicroserviceImpl7.txt
benchstat testdata/BenchmarkFillQueuev1.0.0.txt testdata/BenchmarkFillImpl7.txt
benchstat testdata/BenchmarkRefillQueuev1.0.0.txt testdata/BenchmarkRefillImpl7.txt
benchstat testdata/BenchmarkRefillFullQueuevv1.0.0.txt testdata/BenchmarkRefillFullImpl7.txt
benchstat testdata/BenchmarkSlowIncreaseQueuevv1.0.0.txt testdata/BenchmarkSlowIncreaseImpl7.txt
benchstat testdata/BenchmarkSlowDecreaseQueuevv1.0.0.txt testdata/BenchmarkSlowDecreaseImpl7.txt
benchstat testdata/BenchmarkStableQueuevv1.0.0.txt testdata/BenchmarkStableImpl7.txt
```

Queue vs eapache
```
benchstat testdata/BenchmarkMicroserviceQueuev1.0.0.txt testdata/BenchmarkMicroserviceEapache.txt
benchstat testdata/BenchmarkFillQueuev1.0.0.txt testdata/BenchmarkFillEapache.txt
benchstat testdata/BenchmarkRefillQueuev1.0.0.txt testdata/BenchmarkRefillEapache.txt
benchstat testdata/BenchmarkRefillFullQueuevv1.0.0.txt testdata/BenchmarkRefillFullEapache.txt
benchstat testdata/BenchmarkSlowIncreaseQueuevv1.0.0.txt testdata/BenchmarkSlowIncreaseEapache.txt
benchstat testdata/BenchmarkSlowDecreaseQueuevv1.0.0.txt testdata/BenchmarkSlowDecreaseEapache.txt
benchstat testdata/BenchmarkStableQueuevv1.0.0.txt testdata/BenchmarkStableEapache.txt
```

Queue vs cookiejar
```
benchstat testdata/BenchmarkMicroserviceQueuev1.0.0.txt testdata/BenchmarkMicroserviceCookiejar.txt
benchstat testdata/BenchmarkFillQueuev1.0.0.txt testdata/BenchmarkFillCookiejar.txt
benchstat testdata/BenchmarkRefillQueuev1.0.0.txt testdata/BenchmarkRefillCookiejar.txt
benchstat testdata/BenchmarkRefillFullQueuevv1.0.0.txt testdata/BenchmarkRefillFullCookiejar.txt
benchstat testdata/BenchmarkSlowIncreaseQueuevv1.0.0.txt testdata/BenchmarkSlowIncreaseCookiejar.txt
benchstat testdata/BenchmarkSlowDecreaseQueuevv1.0.0.txt testdata/BenchmarkSlowDecreaseCookiejar.txt
benchstat testdata/BenchmarkStableQueuevv1.0.0.txt testdata/BenchmarkStableCookiejar.txt
```

Queue vs deque
```
benchstat testdata/BenchmarkMicroserviceQueuev1.0.0.txt testdata/BenchmarkMicroserviceDeque.txt
benchstat testdata/BenchmarkFillQueuev1.0.0.txt testdata/BenchmarkFillDeque.txt
benchstat testdata/BenchmarkRefillQueuev1.0.0.txt testdata/BenchmarkRefillDeque.txt
benchstat testdata/BenchmarkRefillFullQueuevv1.0.0.txt testdata/BenchmarkRefillFullDeque.txt
benchstat testdata/BenchmarkSlowIncreaseQueuevv1.0.0.txt testdata/BenchmarkSlowIncreaseDeque.txt
benchstat testdata/BenchmarkSlowDecreaseQueuevv1.0.0.txt testdata/BenchmarkSlowDecreaseDeque.txt
benchstat testdata/BenchmarkStableQueuevv1.0.0.txt testdata/BenchmarkStableDeque.txt
```
