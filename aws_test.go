//   Copyright 2013 Matt Ho
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

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
