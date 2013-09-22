// Copyright 2013 Steven Le. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package queue

import (
	"errors"
)

const (
	InitialBufferSize = 4
)

var (
	ErrEmptyQueue = errors.New("No items in queue!")
)

type Queue struct {
	items      []interface{}
	head       uint
	tail       uint
	length     uint
	bufferSize uint
}

type Item interface{}

func NewQueue() *Queue {
	q := Queue{
		items:      make([]interface{}, InitialBufferSize),
		head:       0,
		tail:       0,
		length:     0,
		bufferSize: InitialBufferSize,
	}
	return &q
}

func (q *Queue) Push(item Item) {
	// If the queue is at capacity, double the items array and copy the old array
	// to the new one.
	if q.length == q.bufferSize {
		q.doubleBufferSize()
	}

	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.bufferSize
	q.length += 1
}

func (q *Queue) Pop() (Item, error) {
	if q.length == 0 {
		return nil, ErrEmptyQueue
	}
	item := q.items[q.head]
	q.items[q.head] = nil
	q.head = (q.head + 1) % q.bufferSize
	q.length -= 1
	return item, nil
}

func (q *Queue) Length() uint {
	return q.length
}

func (q *Queue) doubleBufferSize() {
	items := make([]interface{}, q.bufferSize*2)
	for i := uint(0); i < q.length; i += 1 {
		items[i] = q.items[(q.head+i)%q.length]
	}
	q.items = items
	q.bufferSize *= 2
	q.head = 0
	q.tail = q.length
}
