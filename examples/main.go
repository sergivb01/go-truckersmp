package main

import (
	"fmt"

	"github.com/sergivb01/go-truckersmp"
)

func main() {
	cli := truckersmp.NewClient(nil)
	p, err := cli.PlayerLookup("76561198057763917")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p)

}
