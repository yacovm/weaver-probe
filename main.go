/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"github.com/yacovm/weaver-probe/relay"
	"net"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s --address <host:port>\n", os.Args[0])
	os.Exit(2)
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 || args[0] != "--address"{
		usage()
	}

	hostport := args[1]
	_, _, err := net.SplitHostPort(hostport)
	if err != nil {
		usage()
	}

	client := relay.Client{}
	client.RequestState(args[1])
}