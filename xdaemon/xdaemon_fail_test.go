/*
 * Copyright 2012-2019 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * A toolkit for Golang development
 * https://www.likexian.com/
 */

package xdaemon

import (
	"github.com/likexian/gokit/assert"
	"os"
	"testing"
)

func TestFailDaemon(t *testing.T) {
	v := os.Args[0]
	os.Args[0] = "xx"

	c := Config{
		Pid:   "",
		Log:   "",
		User:  "",
		Chdir: "",
	}
	err := c.Daemon()
	assert.NotNil(t, err)

	c = Config{
		Pid:   "/tmp/test.pid",
		Log:   "/tmp/test.log",
		User:  "nobody",
		Chdir: "/",
	}
	err = c.Daemon()
	assert.NotNil(t, err)

	os.Args[0] = v
}
