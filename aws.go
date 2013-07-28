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
	"launchpad.net/goamz/aws"
	goec2 "launchpad.net/goamz/ec2"
)

func Hosts(region aws.Region, tags map[string]string) ([]string, error) {
	instances, err := Instances(region, tags)
	if err != nil {
		return nil, err
	}

	hosts := make([]string, 0)
	for _, instance := range instances {
		hosts = append(hosts, instance.DNSName)
	}

	return hosts, nil
}

func Instances(region aws.Region, tags map[string]string) ([]goec2.Instance, error) {
	auth, err := aws.EnvAuth()
	if err != nil {
		return nil, err
	}

	ec2 := goec2.New(auth, region)
	response, err := ec2.Instances(nil, nil)
	if err != nil {
		return nil, err
	}

	instances := make([]goec2.Instance, 0)
	for _, reservation := range response.Reservations {
		for _, instance := range reservation.Instances {
			found := true

			// convert the tags into a map[string]string
			actual := map[string]string{}
			for _, tag := range instance.Tags {
				actual[tag.Key] = tag.Value
			}

			// now verify that all the requested values have been found
			for key, value := range tags {
				if actual[key] != value {
					found = false
				}
			}

			// if found, add this to the list
			if found {
				instances = append(instances, instance)
			}
		}
	}

	return instances, nil
}
