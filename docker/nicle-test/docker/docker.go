package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/nicle-lin/dockerV1.2.0/docker/api/client"
)


func main() {
	flHosts := []string{"unix://var/run/socket","tcp://host:port","unix:///path/to/socket" ,"fd://socketfd"}
	protoAddrParts := strings.SplitN(flHosts[0], "://", 2)
	fmt.Println(protoAddrParts)
	mainDaemon()
	cli := client.NewDockerCli(os.Stdin, os.Stdout, os.Stderr, protoAddrParts[0], protoAddrParts[1], nil)
	if err := cli.Cmd([]string{"pull","ubuntu"}); err != nil {
		fmt.Println(err)
	}

}
