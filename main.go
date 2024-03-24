package main

import (
	"fmt"
	"github.com/mgranderath/traceroute/methods"
	"github.com/mgranderath/traceroute/methods/tcp"
	"github.com/mgranderath/traceroute/methods/udp"
	"log"
	"net"
	"sort"
	"time"
)

func main() {
	ip := net.ParseIP("1.1.1.1")
	tcpTraceroute := tcp.New(ip, methods.TracerouteConfig{
		MaxHops:          12,
		NumMeasurements:  3,
		ParallelRequests: 1,
		Port:             53,
		Timeout:          time.Second / 2,
	})
	res, err := tcpTraceroute.Start()

	printResults(res)

	if err != nil {
		log.Fatal(err)
	}
	//log.Println(res)
	udpTraceroute := udp.New(ip, true, methods.TracerouteConfig{
		MaxHops:          20,
		NumMeasurements:  1,
		ParallelRequests: 24,
		Port:             784,
		Timeout:          2 * time.Second,
	})
	res, err = udpTraceroute.Start()
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(res)
}

func printResults(res *map[uint16][]methods.TracerouteHop) {
	var keys []int
	for key, _ := range *res {
		keys = append(keys, int(key))
	}
	sort.Ints(keys)
	for _, key := range keys {
		measures := (*res)[uint16(key)]
		fmt.Printf("%d\n", key)
		for _, measure := range measures {
			fmt.Printf("  %+v\n", measure)
		}
	}
	//fmt.Printf("%d: %+v\n", key, value)
}
