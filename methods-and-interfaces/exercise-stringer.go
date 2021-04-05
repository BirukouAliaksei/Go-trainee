package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (a IPAddr) String() string {
	var s string
	for _, bit := range a {
		s += fmt.Sprintf("%v.", bit)
	}
	return fmt.Sprintf("%v", s)
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
