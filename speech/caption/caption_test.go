// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

func TestRecognize(t *testing.T) {
	testutil.SystemTest(t)

	var buf bytes.Buffer

	if err := recognize(&buf, "../testdata/quit.raw"); err != nil {
		t.Fatal(err)
	}
	if len(buf.String()) == 0 {
		t.Fatal("got no results; want at least one")
	}

	if got, want := buf.String(), "quit"; !strings.Contains(got, want) {
		t.Errorf("Transcript: got %q; want %q", got, want)
	}
}

func TestRecognizeGCS(t *testing.T) {
	t.Skip("https://github.com/GoogleCloudPlatform/golang-samples/issues/2349")
	testutil.SystemTest(t)

	var buf bytes.Buffer

	if err := recognizeGCS(&buf, "gs://python-docs-samples-tests/speech/audio.raw"); err != nil {
		t.Fatal(err)
	}
	if len(buf.String()) == 0 {
		t.Fatal("got no results; want at least one")
	}

	if got, want := buf.String(), "how old is the Brooklyn Bridge"; !strings.Contains(got, want) {
		t.Errorf("Transcript: got %q; want %q", got, want)
	}
}
