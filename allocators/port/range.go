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
	"errors"
	"strconv"

	"github.com/mars1024/allocator"
)

var (
	ErrBadRangeID = errors.New("bad rangeID")
)

type ranger struct {
	lower int
	upper int
}

func parseRangeID(rangeID allocator.RangeID) (int, error) {
	port, err := strconv.Atoi(string(rangeID))
	if err != nil {
		return 0, ErrBadRangeID
	}

	return port, nil
}

func generateRangeID(port int) allocator.RangeID {
	return allocator.RangeID(strconv.Itoa(port))
}

func (r *ranger) Contains(rangeID allocator.RangeID) bool {
	port, err := parseRangeID(rangeID)

	return (err == nil) && (port >= r.lower) && (port <= r.upper)
}

func (r *ranger) First() allocator.RangeIterator {
	return &iterator{
		ranger: r,
		cur:    r.lower,
	}
}

type iterator struct {
	ranger *ranger
	cur    int
}

func (i *iterator) Get() (allocator.RangeID, interface{}) {
	return generateRangeID(i.cur), i.cur
}

func (i *iterator) Next() {
	i.cur++
}

func (i *iterator) InRange() bool {
	return (i.cur >= i.ranger.lower) && (i.cur <= i.ranger.upper)
}
