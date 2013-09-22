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
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	_, err := q.Pop()
	if err == nil {
		t.Error("Expected nil")
	}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	i, err := q.Pop()
	if i != 1 {
		t.Errorf("Failed!!!")
		return
	}
	q.Push(5)
	q.Push(6)

	for j := 2; j <= 6; j += 1 {
		i, err = q.Pop()
		if err != nil {
			t.Errorf("Error: %v", err)
			return
		}
		if i != j {
			t.Errorf("%v != %v", i, j)
			return
		}
	}
}
