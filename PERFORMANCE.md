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
Given the enormous amount of test data, it can be difficult and time consuming to find out the net impact of all the tests,
so we generally spend most of the time on the results of the Microservice test, which is a composite test of all other tests, and the very simple Fill tests, which just sequentially add and remove N number of items.

Below results is for queue [v1.0.0](https://github.com/ef-ds/queue/blob/master/CHANGELOG.md).


### Microservice Test Results
TO BE POSTED.
