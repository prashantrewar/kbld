// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

// Using this code as is from: https://github.com/google/go-containerregistry/tree/master/pkg/v1/internal

// Copyright 2020 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package and

import (
	"bytes"
	"io"
	"testing"
)

func TestRead(t *testing.T) {
	want := "asdf"
	r := bytes.NewBufferString(want)
	called := false

	rac := &ReadCloser{
		Reader: r,
		CloseFunc: func() error {
			called = true
			return nil
		},
	}

	data, err := io.ReadAll(rac)
	if err != nil {
		t.Errorf("ReadAll(rac) = %v", err)
	}
	if got := string(data); got != want {
		t.Errorf("ReadAll(rac); got %q, want %q", got, want)
	}

	if called {
		t.Error("called before Close(); got true, wanted false")
	}
	if err := rac.Close(); err != nil {
		t.Errorf("Close() = %v", err)
	}
	if !called {
		t.Error("called after Close(); got false, wanted true")
	}
}

func TestWrite(t *testing.T) {
	w := bytes.NewBuffer([]byte{})
	called := false

	wac := &WriteCloser{
		Writer: w,
		CloseFunc: func() error {
			called = true
			return nil
		},
	}

	want := "asdf"
	if _, err := wac.Write([]byte(want)); err != nil {
		t.Errorf("Write(%q); = %v", want, err)
	}

	if called {
		t.Error("called before Close(); got true, wanted false")
	}
	if err := wac.Close(); err != nil {
		t.Errorf("Close() = %v", err)
	}
	if !called {
		t.Error("called after Close(); got false, wanted true")
	}

	if got := w.String(); got != want {
		t.Errorf("w.String(); got %q, want %q", got, want)
	}
}
