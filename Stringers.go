package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var output string
	output = fmt.Sprintf("%v", ip[0])

	for i := 1; i < len(ip); i++ {
		output = output + "." + fmt.Sprintf("%v", ip[i])
	}

	return output
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
