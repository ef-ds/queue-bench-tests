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

import (
	"container/list"
	"testing"

	"github.com/christianrpetrin/queue-tests/queueimpl7"
	"github.com/ef-ds/benchmark"
	"github.com/ef-ds/deque"
	"github.com/ef-ds/queue"
	eapache "gopkg.in/eapache/queue.v1"
	cookiejar "gopkg.in/karalabe/cookiejar.v2/collections/queue"
)

// tests variable is of type benchmark.Tests (https://github.com/ef-ds/benchmark/blob/master/tests.go)
// and is declared in the testdata.go source file.

func BenchmarkRefillList(b *testing.B) {
	var l *list.List
	tests.Refill(
		b,
		func() {
			l = list.New()
		},
		func(v interface{}) {
			l.PushBack(v)
		},
		func() (interface{}, bool) {
			return l.Remove(l.Front()), true
		},
		func() bool {
			return l.Front() == nil
		},
	)
}

func BenchmarkRefillSlice(b *testing.B) {
	var q *CustomSliceQueue
	tests.Refill(
		b,
		func() {
			q = NewCustomSliceQueue()
		},
		func(v interface{}) {
			q.Push(v.(*benchmark.TestValue))
		},
		func() (interface{}, bool) {
			return q.Pop()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkRefillImpl7(b *testing.B) {
	var q *queueimpl7.Queueimpl7
	tests.Refill(
		b,
		func() {
			q = queueimpl7.New()
		},
		func(v interface{}) {
			q.Push(v)
		},
		func() (interface{}, bool) {
			return q.Pop()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkRefillEapache(b *testing.B) {
	var q *eapache.Queue
	tests.Refill(
		b,
		func() {
			q = eapache.New()
		},
		func(v interface{}) {
			q.Add(v)
		},
		func() (interface{}, bool) {
			return q.Remove(), true
		},
		func() bool {
			return q.Length() == 0
		},
	)
}
func BenchmarkRefillCookiejar(b *testing.B) {
	var q *cookiejar.Queue
	tests.Refill(
		b,
		func() {
			q = cookiejar.New()
		},
		func(v interface{}) {
			q.Push(v)
		},
		func() (interface{}, bool) {
			return q.Pop(), true
		},
		func() bool {
			return q.Size() == 0
		},
	)
}

func BenchmarkRefillDeque(b *testing.B) {
	var q *deque.Deque
	tests.Refill(
		b,
		func() {
			q = deque.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopFront()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkRefillQueue(b *testing.B) {
	var q *queue.Queue
	tests.Refill(
		b,
		func() {
			q = queue.New()
		},
		func(v interface{}) {
			q.Push(v)
		},
		func() (interface{}, bool) {
			return q.Pop()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}
