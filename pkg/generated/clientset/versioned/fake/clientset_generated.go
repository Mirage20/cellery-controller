/*
 * Copyright (c) 2019 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "cellery.io/cellery-controller/pkg/generated/clientset/versioned"
	authenticationv1alpha1 "cellery.io/cellery-controller/pkg/generated/clientset/versioned/typed/authentication/v1alpha1"
	fakeauthenticationv1alpha1 "cellery.io/cellery-controller/pkg/generated/clientset/versioned/typed/authentication/v1alpha1/fake"
	meshv1alpha2 "cellery.io/cellery-controller/pkg/generated/clientset/versioned/typed/mesh/v1alpha2"
	fakemeshv1alpha2 "cellery.io/cellery-controller/pkg/generated/clientset/versioned/typed/mesh/v1alpha2/fake"
	networkingv1alpha3 "cellery.io/cellery-controller/pkg/generated/clientset/versioned/typed/networking/v1alpha3"
	fakenetworkingv1alpha3 "cellery.io/cellery-controller/pkg/generated/clientset/versioned/typed/networking/v1alpha3/fake"
	servingv1alpha1 "cellery.io/cellery-controller/pkg/generated/clientset/versioned/typed/serving/v1alpha1"
	fakeservingv1alpha1 "cellery.io/cellery-controller/pkg/generated/clientset/versioned/typed/serving/v1alpha1/fake"
	servingv1beta1 "cellery.io/cellery-controller/pkg/generated/clientset/versioned/typed/serving/v1beta1"
	fakeservingv1beta1 "cellery.io/cellery-controller/pkg/generated/clientset/versioned/typed/serving/v1beta1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// AuthenticationV1alpha1 retrieves the AuthenticationV1alpha1Client
func (c *Clientset) AuthenticationV1alpha1() authenticationv1alpha1.AuthenticationV1alpha1Interface {
	return &fakeauthenticationv1alpha1.FakeAuthenticationV1alpha1{Fake: &c.Fake}
}

// MeshV1alpha2 retrieves the MeshV1alpha2Client
func (c *Clientset) MeshV1alpha2() meshv1alpha2.MeshV1alpha2Interface {
	return &fakemeshv1alpha2.FakeMeshV1alpha2{Fake: &c.Fake}
}

// NetworkingV1alpha3 retrieves the NetworkingV1alpha3Client
func (c *Clientset) NetworkingV1alpha3() networkingv1alpha3.NetworkingV1alpha3Interface {
	return &fakenetworkingv1alpha3.FakeNetworkingV1alpha3{Fake: &c.Fake}
}

// ServingV1alpha1 retrieves the ServingV1alpha1Client
func (c *Clientset) ServingV1alpha1() servingv1alpha1.ServingV1alpha1Interface {
	return &fakeservingv1alpha1.FakeServingV1alpha1{Fake: &c.Fake}
}

// ServingV1beta1 retrieves the ServingV1beta1Client
func (c *Clientset) ServingV1beta1() servingv1beta1.ServingV1beta1Interface {
	return &fakeservingv1beta1.FakeServingV1beta1{Fake: &c.Fake}
}
