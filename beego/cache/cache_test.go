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

package cache

import (
	"github.com/astaxie/beego/logs"
	"os"
	"testing"
	"time"
)

func TestCache(t *testing.T) {

	InitCache("memory", `{"interval":20}`)

	err := Put("learn", 1, 10)
	if err != nil {
		t.Error("set Error", err)
	}

	a := Get("learn")
	logs.Info(a)

	if !IsExist("learn") {
		t.Error("check err")
	}

	if v := Get("learn"); v.(int) != 1 {
		t.Error("get err")
	}

	time.Sleep(30 * time.Second)

	if IsExist("learn") {
		t.Error("check err")
	}

	if err := Put("learn", 1, 10); err != nil {
		t.Error("set Error", err)
	}

	if err := Incr("learn"); err != nil {
		t.Error("Incr Error", err)
	}

	if v := Get("learn"); v.(int) != 2 {
		t.Error("get err")
	}

	if err := Decr("learn"); err != nil {
		t.Error("Decr Error", err)
	}

	if v := Get("learn"); v.(int) != 1 {
		t.Error("get err")
	}
	Delete("learn")
	if IsExist("learn") {
		t.Error("delete err")
	}

	//test GetMulti
	if err := Put("learn", "author", 10); err != nil {
		t.Error("set Error", err)
	}
	if !IsExist("learn") {
		t.Error("check err")
	}
	if v := Get("learn"); v.(string) != "author" {
		t.Error("get err")
	}

	if err := Put("learn1", "author1", 10); err != nil {
		t.Error("set Error", err)
	}
	if !IsExist("learn1") {
		t.Error("check err")
	}

	vv := GetMulti([]string{"learn", "learn1"})
	if len(vv) != 2 {
		t.Error("GetMulti ERROR")
	}
	if vv[0].(string) != "author" {
		t.Error("GetMulti ERROR")
	}
	if vv[1].(string) != "author1" {
		t.Error("GetMulti ERROR")
	}
}

func TestFileCache(t *testing.T) {
	InitCache("file", `{"CachePath":"cache","FileSuffix":".bin","DirectoryLevel":2,"EmbedExpiry":0}`)

	if err := Put("learn", 1, 10); err != nil {
		t.Error("set Error", err)
	}
	if !IsExist("learn") {
		t.Error("check err")
	}

	if v := Get("learn"); v.(int) != 1 {
		t.Error("get err")
	}

	if err := Incr("learn"); err != nil {
		t.Error("Incr Error", err)
	}

	if v := Get("learn"); v.(int) != 2 {
		t.Error("get err")
	}

	if err := Decr("learn"); err != nil {
		t.Error("Decr Error", err)
	}

	if v := Get("learn"); v.(int) != 1 {
		t.Error("get err")
	}
	Delete("learn")
	if IsExist("learn") {
		t.Error("delete err")
	}

	//test string
	if err := Put("learn", "author", 10); err != nil {
		t.Error("set Error", err)
	}
	if !IsExist("learn") {
		t.Error("check err")
	}
	if v := Get("learn"); v.(string) != "author" {
		t.Error("get err")
	}

	//test GetMulti
	if err := Put("learn1", "author1", 10); err != nil {
		t.Error("set Error", err)
	}
	if !IsExist("learn1") {
		t.Error("check err")
	}

	vv := GetMulti([]string{"learn", "learn1"})
	if len(vv) != 2 {
		t.Error("GetMulti ERROR")
	}
	if vv[0].(string) != "author" {
		t.Error("GetMulti ERROR")
	}
	if vv[1].(string) != "author1" {
		t.Error("GetMulti ERROR")
	}

	os.RemoveAll("cache")
}
