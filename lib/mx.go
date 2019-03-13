package lib

import (
	"fmt"
	"net"
)

func MX(host string) error {
	mx, err := net.LookupMX(host)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < len(mx); i++ {
		fmt.Println(mx[i])
	}
	return nil
}
