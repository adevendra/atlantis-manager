/* Copyright 2014 Ooyala, Inc. All rights reserved.
 *
 * This file is licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is
 * distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

package datamodel

import (
	zookeeper "github.com/jigish/gozk-recipes"
	. "launchpad.net/gocheck"
	"testing"
)

const (
	app  = "my-app"
	sha  = "mysha"
	env  = "myenv"
	host = "my-host"
	pool = "my-pool"
	rule = "my-rule"
	trie = "my-trie"
	dep  = "my-dep"
	opt  = "my-opt"
	repo = "my-repo"
	root = "my-root"
)

func TestDatamodel(t *testing.T) { TestingT(t) }

type DatamodelSuite struct{}

var _ = Suite(&DatamodelSuite{})

var (
	zkTestServer *zookeeper.ZkTestServer
)

func (s *DatamodelSuite) SetUpSuite(c *C) {
	zkTestServer = zookeeper.NewZkTestServer()
	c.Assert(zkTestServer.Init(), IsNil)
	Zk = zkTestServer.Zk
}

func (s *DatamodelSuite) TearDownSuite(c *C) {
	err := zkTestServer.Destroy()
	c.Assert(err, IsNil)
}
