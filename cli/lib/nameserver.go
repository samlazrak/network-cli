package lib

import (
	"fmt"
	"net"
)

func Nameserver(host string) error {
	ns, err := net.LookupNS(host)
	if err != nil {
		fmt.Println(err)
		return (err)
	}
	for i := 0; i < len(ns); i++ {
		fmt.Println(ns[i].Host)
	}
	return nil
}
