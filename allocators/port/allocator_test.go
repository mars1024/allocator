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

package port

import (
	"github.com/mars1024/allocator"
	"reflect"
	"testing"
)

func TestNewPortAllocator(t *testing.T) {
	tests := []struct {
		name  string
		lower int
		upper int
		err   error
	}{
		{
			"lower port is bigger than upper port",
			100,
			99,
			ErrUpperPort,
		},
		{
			"port is le than 0",
			-1,
			0,
			ErrOutOfRange,
		},
		{
			"port is greater than 65535",
			100,
			100000,
			ErrOutOfRange,
		},
		{
			"valid input",
			100,
			200,
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewPortAllocator(test.lower, test.upper)
			if !reflect.DeepEqual(err, test.err) {
				t.Errorf("test %s fails", test.name)
				return
			}
		})
	}
}

func TestAssign(t *testing.T) {
	pa, _ := NewPortAllocator(100, 200)
	_ = pa.Assign("150")

	tests := []struct {
		name    string
		rangeID allocator.RangeID
		err     error
	}{
		{
			"out of range",
			"1000",
			allocator.ErrOutOfRange,
		},
		{
			"allocated",
			"150",
			allocator.ErrAllocated,
		},
		{
			"valid assign",
			"160",
			nil,
		},
		{
			"invalid range id",
			"abcd",
			allocator.ErrOutOfRange,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := pa.Assign(test.rangeID)
			if !reflect.DeepEqual(err, test.err) {
				t.Errorf("test %s fails", test.name)
				return
			}
		})
	}
}

func TestHas(t *testing.T) {
	pa, _ := NewPortAllocator(100, 200)
	_ = pa.Assign("150")

	tests := []struct {
		name    string
		rangeID allocator.RangeID
		has     bool
	}{
		{
			"has",
			"150",
			true,
		},
		{
			"not has",
			"120",
			false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if pa.Has(test.rangeID) != test.has {
				t.Errorf("test fails %s", test.name)
				return
			}
		})
	}
}

func TestRelease(t *testing.T) {
	pa, _ := NewPortAllocator(100, 200)
	_ = pa.Assign("150")

	tests := []struct {
		name    string
		rangeID allocator.RangeID
	}{
		{
			"release",
			"150",
		},
		{
			"release again",
			"150",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_ = pa.Release(test.rangeID)
			if pa.Has(test.rangeID) {
				t.Errorf("test %s fails", test.name)
				return
			}
		})
	}
}

func TestAllocate(t *testing.T) {
	pa, _ := NewPortAllocator(10, 15)
	_ = pa.Assign("13")

	tests := []struct {
		name          string
		expectRangeID allocator.RangeID
	}{
		{
			"allocate 1",
			"10",
		},
		{
			"allocate 2",
			"11",
		},
		{
			"allocate 3",
			"12",
		},
		{
			"allocate 4",
			"14",
		},
		{
			"allocate 5",
			"15",
		},
		{
			"allocate 6",
			"",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if e, _, _ := pa.Allocate(); e != test.expectRangeID {
				t.Errorf("test %s: expect %v but got %v", test.name, test.expectRangeID, e)
				return
			}
		})
	}
}
