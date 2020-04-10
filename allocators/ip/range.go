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
	"math/big"
	"net"

	"github.com/mars1024/allocator"
)

type IP struct {
	address *net.IPNet
	gateway net.IP
}

type ranger struct {
	subnet  *net.IPNet
	start   net.IP
	end     net.IP
	gateway net.IP
}

func NewIPRanger(subnet *net.IPNet, start, end, gateway net.IP) (allocator.Range, error) {
	// TODO: validation and defaulting

	return &ranger{
		subnet:  subnet,
		start:   start,
		end:     end,
		gateway: gateway,
	}, nil
}

func (r *ranger) Contains(rangeID allocator.RangeID) bool {
	ip := parseRangeID(rangeID)

	if !r.subnet.Contains(ip) {
		return false
	}
	if cmp(ip, r.start) < 0 {
		return false
	}
	if cmp(ip, r.end) > 0 {
		return false
	}
	if r.gateway.Equal(ip) {
		return false
	}

	return true
}

func (r *ranger) First() allocator.RangeIterator {
	return &iterator{
		ranger: r,
		cur:    r.start,
	}
}

type iterator struct {
	ranger *ranger
	cur    net.IP
}

func (i *iterator) Get() (allocator.RangeID, interface{}) {
	return generateRangeID(i.cur), &IP{
		address: &net.IPNet{
			IP:   i.cur,
			Mask: i.ranger.subnet.Mask,
		},
		gateway: i.ranger.gateway,
	}
}

func (i *iterator) Next() {
	i.cur = next(i.cur)
}

func (i *iterator) InRange() bool {
	if i.ranger.gateway.Equal(i.cur) {
		i.Next()
		return i.InRange()
	}
	return cmp(i.cur, i.ranger.end) <= 0
}

func parseRangeID(rangeID allocator.RangeID) net.IP {
	return net.ParseIP(string(rangeID))
}

func generateRangeID(ip net.IP) allocator.RangeID {
	return allocator.RangeID(ip.String())
}

func next(ip net.IP) net.IP {
	i := ipToInt(ip)
	return intToIP(i.Add(i, big.NewInt(1)))
}

func cmp(a, b net.IP) int {
	aa := ipToInt(a)
	bb := ipToInt(b)
	return aa.Cmp(bb)
}

func ipToInt(ip net.IP) *big.Int {
	if v := ip.To4(); v != nil {
		return big.NewInt(0).SetBytes(v)
	}
	return big.NewInt(0).SetBytes(ip.To16())
}

func intToIP(i *big.Int) net.IP {
	return net.IP(i.Bytes())
}
