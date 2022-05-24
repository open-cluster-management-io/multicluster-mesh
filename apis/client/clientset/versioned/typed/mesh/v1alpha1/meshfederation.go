/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	scheme "github.com/stolostron/multicluster-mesh-addon/apis/client/clientset/versioned/scheme"
	v1alpha1 "github.com/stolostron/multicluster-mesh-addon/apis/mesh/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MeshFederationsGetter has a method to return a MeshFederationInterface.
// A group's client should implement this interface.
type MeshFederationsGetter interface {
	MeshFederations(namespace string) MeshFederationInterface
}

// MeshFederationInterface has methods to work with MeshFederation resources.
type MeshFederationInterface interface {
	Create(ctx context.Context, meshFederation *v1alpha1.MeshFederation, opts v1.CreateOptions) (*v1alpha1.MeshFederation, error)
	Update(ctx context.Context, meshFederation *v1alpha1.MeshFederation, opts v1.UpdateOptions) (*v1alpha1.MeshFederation, error)
	UpdateStatus(ctx context.Context, meshFederation *v1alpha1.MeshFederation, opts v1.UpdateOptions) (*v1alpha1.MeshFederation, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MeshFederation, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MeshFederationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MeshFederation, err error)
	MeshFederationExpansion
}

// meshFederations implements MeshFederationInterface
type meshFederations struct {
	client rest.Interface
	ns     string
}

// newMeshFederations returns a MeshFederations
func newMeshFederations(c *MeshV1alpha1Client, namespace string) *meshFederations {
	return &meshFederations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the meshFederation, and returns the corresponding meshFederation object, and an error if there is any.
func (c *meshFederations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MeshFederation, err error) {
	result = &v1alpha1.MeshFederation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("meshfederations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MeshFederations that match those selectors.
func (c *meshFederations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MeshFederationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MeshFederationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("meshfederations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested meshFederations.
func (c *meshFederations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("meshfederations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a meshFederation and creates it.  Returns the server's representation of the meshFederation, and an error, if there is any.
func (c *meshFederations) Create(ctx context.Context, meshFederation *v1alpha1.MeshFederation, opts v1.CreateOptions) (result *v1alpha1.MeshFederation, err error) {
	result = &v1alpha1.MeshFederation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("meshfederations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(meshFederation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a meshFederation and updates it. Returns the server's representation of the meshFederation, and an error, if there is any.
func (c *meshFederations) Update(ctx context.Context, meshFederation *v1alpha1.MeshFederation, opts v1.UpdateOptions) (result *v1alpha1.MeshFederation, err error) {
	result = &v1alpha1.MeshFederation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("meshfederations").
		Name(meshFederation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(meshFederation).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *meshFederations) UpdateStatus(ctx context.Context, meshFederation *v1alpha1.MeshFederation, opts v1.UpdateOptions) (result *v1alpha1.MeshFederation, err error) {
	result = &v1alpha1.MeshFederation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("meshfederations").
		Name(meshFederation.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(meshFederation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the meshFederation and deletes it. Returns an error if one occurs.
func (c *meshFederations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("meshfederations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *meshFederations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("meshfederations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched meshFederation.
func (c *meshFederations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MeshFederation, err error) {
	result = &v1alpha1.MeshFederation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("meshfederations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
