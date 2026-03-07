package main

import (
	"flag"
	"fmt"
	"log"

	"example/go-netcalc/netcalc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cidr := flag.String("cidr", "", "CIDR input in format <ip>/<prefix> (e.g. 192.168.1.10/24)")
	flag.Parse()

	if *cidr == "" {
		log.Fatal("Usage: go run . -cidr 192.168.1.10/24")
	}

	network, err := netcalc.ParseCIDR(*cidr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Input:               %s/%d\n", netcalc.FormatIPv4(network.IPv4), network.Prefix)
	fmt.Println("Network Mask:       ", netcalc.FormatIPv4(network.Mask()))
	fmt.Println("Network Address:    ", netcalc.FormatIPv4(network.Address()))
	fmt.Println("Broadcast Address:  ", netcalc.FormatIPv4(network.BroadcastAddress()))
	fmt.Println("Host Address Range: ", netcalc.FormatIPv4(network.FirstHostAddress())+" - "+netcalc.FormatIPv4(network.LastHostAddress()))
	fmt.Println("Number of Host IPs: ", 1<<(32-network.Prefix)-2) // 2 ^ (32 - prefix) - 2
}
