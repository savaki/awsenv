package main

import (
	awsenv ".."
	"flag"
	"fmt"
	"launchpad.net/goamz/aws"
	"log"
	"strings"
)

var (
	env       = flag.String("env", "development", "env tag to filter on")
	regionStr = flag.String("region", "us-west-2", "the aws region to scan")
)

func toFilter(args []string) map[string]string {
	params := map[string]string{"env": *env}

	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			log.Fatalf("invalid argument, %s", arg)
		}
		key := parts[0]
		value := parts[1]
		params[key] = value
	}

	return params
}

func main() {
	flag.Parse()
	filter := toFilter(flag.Args())

	region, found := aws.Regions[*regionStr]
	if !found {
		log.Fatalf("unknown region, %s", regionStr)
	}

	hosts, err := awsenv.Hosts(region, filter)
	if err != nil {
		log.Fatal(err)
	}

	for _, host := range hosts {
		fmt.Println(host)
	}
}
