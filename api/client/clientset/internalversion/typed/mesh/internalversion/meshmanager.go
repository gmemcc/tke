/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	mesh "tkestack.io/tke/api/mesh"
)

// MeshManagersGetter has a method to return a MeshManagerInterface.
// A group's client should implement this interface.
type MeshManagersGetter interface {
	MeshManagers() MeshManagerInterface
}

// MeshManagerInterface has methods to work with MeshManager resources.
type MeshManagerInterface interface {
	Create(ctx context.Context, meshManager *mesh.MeshManager, opts v1.CreateOptions) (*mesh.MeshManager, error)
	Update(ctx context.Context, meshManager *mesh.MeshManager, opts v1.UpdateOptions) (*mesh.MeshManager, error)
	UpdateStatus(ctx context.Context, meshManager *mesh.MeshManager, opts v1.UpdateOptions) (*mesh.MeshManager, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*mesh.MeshManager, error)
	List(ctx context.Context, opts v1.ListOptions) (*mesh.MeshManagerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *mesh.MeshManager, err error)
	MeshManagerExpansion
}

// meshManagers implements MeshManagerInterface
type meshManagers struct {
	client rest.Interface
}

// newMeshManagers returns a MeshManagers
func newMeshManagers(c *MeshClient) *meshManagers {
	return &meshManagers{
		client: c.RESTClient(),
	}
}

// Get takes name of the meshManager, and returns the corresponding meshManager object, and an error if there is any.
func (c *meshManagers) Get(ctx context.Context, name string, options v1.GetOptions) (result *mesh.MeshManager, err error) {
	result = &mesh.MeshManager{}
	err = c.client.Get().
		Resource("meshmanagers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MeshManagers that match those selectors.
func (c *meshManagers) List(ctx context.Context, opts v1.ListOptions) (result *mesh.MeshManagerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &mesh.MeshManagerList{}
	err = c.client.Get().
		Resource("meshmanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested meshManagers.
func (c *meshManagers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("meshmanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a meshManager and creates it.  Returns the server's representation of the meshManager, and an error, if there is any.
func (c *meshManagers) Create(ctx context.Context, meshManager *mesh.MeshManager, opts v1.CreateOptions) (result *mesh.MeshManager, err error) {
	result = &mesh.MeshManager{}
	err = c.client.Post().
		Resource("meshmanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(meshManager).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a meshManager and updates it. Returns the server's representation of the meshManager, and an error, if there is any.
func (c *meshManagers) Update(ctx context.Context, meshManager *mesh.MeshManager, opts v1.UpdateOptions) (result *mesh.MeshManager, err error) {
	result = &mesh.MeshManager{}
	err = c.client.Put().
		Resource("meshmanagers").
		Name(meshManager.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(meshManager).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *meshManagers) UpdateStatus(ctx context.Context, meshManager *mesh.MeshManager, opts v1.UpdateOptions) (result *mesh.MeshManager, err error) {
	result = &mesh.MeshManager{}
	err = c.client.Put().
		Resource("meshmanagers").
		Name(meshManager.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(meshManager).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the meshManager and deletes it. Returns an error if one occurs.
func (c *meshManagers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("meshmanagers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched meshManager.
func (c *meshManagers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *mesh.MeshManager, err error) {
	result = &mesh.MeshManager{}
	err = c.client.Patch(pt).
		Resource("meshmanagers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
