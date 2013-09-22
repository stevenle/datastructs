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

package set

import (
  "testing"
)

func TestSet(t *testing.T) {
  s := New()
  s.Add(1)
  s.Add(1)
  if s.Len() != 1 {
    t.Errorf("Expected length to be 1, but was %v", s.Len())
  }
  if !s.Contains(1) {
    t.Error("1 was not found in set")
  }

  s.Add(2)
  if s.Len() != 2 {
    t.Errorf("Expected length to be 2, but was %v", s.Len())
  }
  if !s.Contains(2) {
    t.Error("2 was not found in set")
  }

  s.Remove(1)
  if s.Len() != 1 {
    t.Errorf("Expected length to be 1, but was %v", s.Len())
  }
  if s.Contains(1) {
    t.Error("1 was not removed from set")
  }
}
