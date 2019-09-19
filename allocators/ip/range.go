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

package ip

import (
	"net"

	"github.com/mars1024/allocator"
)

type ranger struct {
	subnet  *net.IPNet
	start   net.IP
	end     net.IP
	gateway net.IP
}

func (*ranger) Contains(allocator.RangeID) bool {
	panic("implement me")
}

func (*ranger) First() allocator.RangeIterator {
	panic("implement me")
}

type iterator struct {
	ranger *ranger
	cur    net.IP
}

func (*iterator) Get() (allocator.RangeID, interface{}) {
	panic("implement me")
}

func (*iterator) Next() {
	panic("implement me")
}

func (*iterator) InRange() bool {
	panic("implement me")
}
