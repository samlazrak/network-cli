package lib

import (
	"fmt"
	"net"
)

func IP(host string) error {
	ip, err := net.LookupIP(host)
	if err != nil {
		fmt.Println(err)
		return (err)
	}
	for i := 0; i < len(ip); i++ {
		fmt.Println(ip[i])
	}
	return nil
}
