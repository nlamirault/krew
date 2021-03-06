// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integrationtest

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKrewList(t *testing.T) {
	skipShort(t)

	test, cleanup := NewTest(t)
	defer cleanup()

	initialList := test.WithIndex().Krew("list").RunOrFailOutput()
	initialOut := []byte{'\n'}

	if diff := cmp.Diff(initialList, initialOut); diff != "" {
		t.Fatalf("expected empty output from 'list':\n%s", diff)
	}

	test.Krew("install", validPlugin).RunOrFail()
	expected := []byte(validPlugin + "\n")

	eventualList := test.Krew("list").RunOrFailOutput()
	if diff := cmp.Diff(eventualList, expected); diff != "" {
		t.Fatalf("'list' output doesn't match:\n%s", diff)
	}
	// TODO(ahmetb): install multiple plugins and see if the output is sorted
}
