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

package fake

import (
	"context"

	v1alpha1 "github.com/stolostron/multicluster-mesh-addon/apis/mesh/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMeshes implements MeshInterface
type FakeMeshes struct {
	Fake *FakeMeshV1alpha1
	ns   string
}

var meshesResource = schema.GroupVersionResource{Group: "mesh.open-cluster-management.io", Version: "v1alpha1", Resource: "meshes"}

var meshesKind = schema.GroupVersionKind{Group: "mesh.open-cluster-management.io", Version: "v1alpha1", Kind: "Mesh"}

// Get takes name of the mesh, and returns the corresponding mesh object, and an error if there is any.
func (c *FakeMeshes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Mesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(meshesResource, c.ns, name), &v1alpha1.Mesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Mesh), err
}

// List takes label and field selectors, and returns the list of Meshes that match those selectors.
func (c *FakeMeshes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MeshList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(meshesResource, meshesKind, c.ns, opts), &v1alpha1.MeshList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MeshList{ListMeta: obj.(*v1alpha1.MeshList).ListMeta}
	for _, item := range obj.(*v1alpha1.MeshList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested meshes.
func (c *FakeMeshes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(meshesResource, c.ns, opts))

}

// Create takes the representation of a mesh and creates it.  Returns the server's representation of the mesh, and an error, if there is any.
func (c *FakeMeshes) Create(ctx context.Context, mesh *v1alpha1.Mesh, opts v1.CreateOptions) (result *v1alpha1.Mesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(meshesResource, c.ns, mesh), &v1alpha1.Mesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Mesh), err
}

// Update takes the representation of a mesh and updates it. Returns the server's representation of the mesh, and an error, if there is any.
func (c *FakeMeshes) Update(ctx context.Context, mesh *v1alpha1.Mesh, opts v1.UpdateOptions) (result *v1alpha1.Mesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(meshesResource, c.ns, mesh), &v1alpha1.Mesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Mesh), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMeshes) UpdateStatus(ctx context.Context, mesh *v1alpha1.Mesh, opts v1.UpdateOptions) (*v1alpha1.Mesh, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(meshesResource, "status", c.ns, mesh), &v1alpha1.Mesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Mesh), err
}

// Delete takes name of the mesh and deletes it. Returns an error if one occurs.
func (c *FakeMeshes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(meshesResource, c.ns, name, opts), &v1alpha1.Mesh{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMeshes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(meshesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MeshList{})
	return err
}

// Patch applies the patch and returns the patched mesh.
func (c *FakeMeshes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Mesh, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(meshesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Mesh{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Mesh), err
}
