// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crd_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ligato/networkservicemesh/plugins/crd"
	"github.com/ligato/networkservicemesh/utils/helper/deptools"
	"github.com/ligato/networkservicemesh/utils/registry/testsuites"
)

func TestSharedPlugin(t *testing.T) {
	RegisterTestingT(t)
	plugin1 := crd.SharedPlugin()
	Expect(plugin1).NotTo(BeNil())
	plugin2 := crd.SharedPlugin()
	Expect(plugin2 == plugin1).To(BeTrue())
	Expect(deptools.Check(plugin1)).To(Succeed())
}

func TestNewPluginNotShared(t *testing.T) {
	RegisterTestingT(t)
	plugin1 := crd.NewPlugin()
	Expect(plugin1).NotTo(BeNil())
	plugin2 := crd.SharedPlugin()
	Expect(plugin2 == plugin1).ToNot(BeTrue())
	Expect(deptools.Check(plugin1)).To(Succeed())
	Expect(deptools.Check(plugin2)).To(Succeed())
}

func TestSharedPluginRemoveOnClose(t *testing.T) {
	RegisterTestingT(t)
	plugin1 := crd.SharedPlugin()
	plugin1.Init()
	plugin1.Close()
	plugin2 := crd.SharedPlugin()
	Expect(plugin2 == plugin1).ToNot(BeTrue())
	Expect(deptools.Check(plugin1)).To(Succeed())
}

func TestNonDefaultName(t *testing.T) {
	RegisterTestingT(t)
	name := "foo"
	plugin := crd.NewPlugin(crd.UseDeps(&crd.Deps{Name: name}))
	Expect(plugin).NotTo(BeNil())
	Expect(deptools.Check(plugin)).To(Succeed())
}

func TestWithRegistry(t *testing.T) {
	name := "foo"
	p1 := crd.NewPlugin()
	p2 := crd.NewPlugin()
	p3 := crd.NewPlugin(crd.UseDeps(&crd.Deps{Name: name}))
	testsuites.SuiteRegistry(t, p1, p2, p3)
}
