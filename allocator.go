/*
 Copyright 2019 Bruce Ma

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package allocator

import "errors"

type Interface interface {
	Assign(RangeID) error
	Allocate() (RangeID, interface{}, error)
	Release(RangeID) error
	Has(RangeID) bool
}

var (
	ErrFull       = errors.New("range is full")
	ErrAllocated  = errors.New("already allocated")
	ErrOutOfRange = errors.New("out of range")
)

type allocator struct {
	ranger Range
	store  map[RangeID]struct{}
}

func NewAllocator(ranger Range) Interface {
	return &allocator{
		ranger: ranger,
		store:  make(map[RangeID]struct{}),
	}
}

func (a *allocator) Assign(id RangeID) error {
	if ! a.ranger.Contains(id) {
		return ErrOutOfRange
	}
	if a.Has(id) {
		return ErrAllocated
	}
	a.store[id] = struct{}{}
	return nil
}

func (a *allocator) Allocate() (RangeID, interface{}, error) {
	for ri := a.ranger.First(); ri.InRange(); ri.Next() {
		id, value := ri.Get()
		if a.Assign(id) == nil {
			return id, value, nil
		}
	}
	return "", nil, ErrFull
}

func (a *allocator) Release(id RangeID) error {
	delete(a.store, id)
	return nil
}

func (a *allocator) Has(id RangeID) bool {
	if _, exist := a.store[id]; exist {
		return true
	}
	return false
}
