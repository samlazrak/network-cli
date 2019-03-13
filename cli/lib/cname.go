package lib

import (
	"fmt"
	"net"
)

func CNAME(host string) error {
	cname, err := net.LookupCNAME(host)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(cname)
	return nil
}
