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

package main

import (
	"testing"
)

func TestToFilter(t *testing.T) {
	// Given
	args := []string{"hello=world", "foo=bar"}

	// When
	filter := toFilter(args)

	// Then
	if len(filter) != 3 {
		t.Fatalf("expected 3 params; actual was %d\n", len(filter))
	}
	if filter["hello"] != "world" {
		t.Fatalf("expected hello => world ; actual => %s\n", filter["hello"])
	}
	if filter["foo"] != "bar" {
		t.Fatalf("expected foo => bar ; actual => %s\n", filter["foo"])
	}
}
