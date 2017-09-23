// Copyright 2014 beego Author. All Rights Reserved.
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

package httplib

import (
	"testing"
)

func TestGet(t *testing.T) {
	b, err := Get("http://httpbin.org/get")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(b)

	if string(b) == "" {
		t.Fatal("request data not match")
	}
}

func TestPost(t *testing.T) {

	type param struct {
		V1 string `json:"v1"`
		V2 int8   `json:"v2"`
	}

	b, err := Post("http://httpbin.org/post", param{
		V1: "33",
		V2: 20,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(b)

	if len(b) <= 0 {
		t.Fatal(" not found in post")
	}
}

func TestPut(t *testing.T) {
	type param struct {
		V1 string `json:"v1"`
		V2 int8   `json:"v2"`
	}

	b, err := Put("http://httpbin.org/put", param{
		V1: "33",
		V2: 20,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(b)

	if len(b) <= 0 {
		t.Fatal(" not found in post")
	}
}
