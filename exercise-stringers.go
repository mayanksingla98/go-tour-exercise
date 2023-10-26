package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (m IPAddr) String() string {
	var ok = make([]string, 4)
	for i, v := range m {
		ok[i] = strconv.Itoa(int(v))
	}
	return strings.Join(ok, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
