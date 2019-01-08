// Copyright (c) 2018 ef-ds
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package queuebenchtests

import "github.com/ef-ds/benchmark"

var (
	tests benchmark.Tests
)

// Pure slice based test queue implementation-------------------------------------------------------

// CustomSliceQueue represents an unbounded, dynamically growing deque customized
// to operate on testVale struct.
type CustomSliceQueue struct {
	// The queue valueq.
	v []*benchmark.TestValue
}

// NewCustomSliceQueue returns an initialized instance of CustomSliceQueue.
func NewCustomSliceQueue() *CustomSliceQueue {
	return new(CustomSliceQueue).Init()
}

// Init initializes or clears queue q.
func (q *CustomSliceQueue) Init() *CustomSliceQueue {
	q.v = make([]*benchmark.TestValue, 0)
	return q
}

// Len returns the length of queue q.
func (q *CustomSliceQueue) Len() int { return len(q.v) }

// Front returns the first element of queue q.
func (q *CustomSliceQueue) Front() (*benchmark.TestValue, bool) {
	if len(q.v) == 0 {
		return nil, false
	}
	return q.v[0], true
}

// Push adds element v to the back of queue q.
func (q *CustomSliceQueue) Push(v *benchmark.TestValue) {
	q.v = append(q.v, v)
}

// Pop removes and returns the first element of queue q.
func (q *CustomSliceQueue) Pop() (*benchmark.TestValue, bool) {
	if len(q.v) == 0 {
		return nil, false
	}

	v := q.v[0]
	q.v[0] = nil // Avoid memory leaks
	q.v = q.v[1:]
	return v, true
}
