package awsenv

import (
	"fmt"
	"launchpad.net/goamz/aws"
	"testing"
)

func TestHosts(t *testing.T) {
	environments := []string{"development", "staging", "production"}
	for _, env := range environments {
		hosts, err := Hosts(aws.USWest2, map[string]string{"env": env})
		if err != nil {
			t.Fatal(err)
		}

		if hosts == nil {
			t.Fatal("no hosts returned")
		}
	}
}

func TestInstances(t *testing.T) {
	environments := []string{"development", "staging", "production"}
	for _, env := range environments {
		fmt.Printf("\n%s\n", env)
		instances, err := Instances(aws.USWest2, map[string]string{"env": env, "server": "tmtt"})
		if err != nil {
			t.Fatal(err)
		}

		for _, instance := range instances {
			fmt.Println(instance)
		}
	}
}
