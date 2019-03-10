package main

import (
	"fmt"

	"github.com/sergivb01/go-truckersmp"
)

func main() {
	cli := truckersmp.NewClient(nil)
	p, err := cli.GetPlayer(76561198057763917)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)

	// the example player has no bans - it will return an empty slice of bans
	bans, err := cli.GetPlayerBans(76561198057763917)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bans)

	servers, err := cli.GetServers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(servers)

	time, err := cli.GetGameTime()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d\n", time)

}
