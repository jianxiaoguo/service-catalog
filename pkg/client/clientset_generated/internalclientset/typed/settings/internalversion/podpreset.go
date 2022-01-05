/*
Copyright 2022 The Kubernetes Authors.

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

package internalversion

import (
	"context"
	"time"

	settings "github.com/kubernetes-sigs/service-catalog/pkg/apis/settings"
	scheme "github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PodPresetsGetter has a method to return a PodPresetInterface.
// A group's client should implement this interface.
type PodPresetsGetter interface {
	PodPresets(namespace string) PodPresetInterface
}

// PodPresetInterface has methods to work with PodPreset resources.
type PodPresetInterface interface {
	Create(ctx context.Context, podPreset *settings.PodPreset, opts v1.CreateOptions) (*settings.PodPreset, error)
	Update(ctx context.Context, podPreset *settings.PodPreset, opts v1.UpdateOptions) (*settings.PodPreset, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*settings.PodPreset, error)
	List(ctx context.Context, opts v1.ListOptions) (*settings.PodPresetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *settings.PodPreset, err error)
	PodPresetExpansion
}

// podPresets implements PodPresetInterface
type podPresets struct {
	client rest.Interface
	ns     string
}

// newPodPresets returns a PodPresets
func newPodPresets(c *SettingsClient, namespace string) *podPresets {
	return &podPresets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the podPreset, and returns the corresponding podPreset object, and an error if there is any.
func (c *podPresets) Get(ctx context.Context, name string, options v1.GetOptions) (result *settings.PodPreset, err error) {
	result = &settings.PodPreset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podpresets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PodPresets that match those selectors.
func (c *podPresets) List(ctx context.Context, opts v1.ListOptions) (result *settings.PodPresetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &settings.PodPresetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podpresets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested podPresets.
func (c *podPresets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("podpresets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a podPreset and creates it.  Returns the server's representation of the podPreset, and an error, if there is any.
func (c *podPresets) Create(ctx context.Context, podPreset *settings.PodPreset, opts v1.CreateOptions) (result *settings.PodPreset, err error) {
	result = &settings.PodPreset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("podpresets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podPreset).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a podPreset and updates it. Returns the server's representation of the podPreset, and an error, if there is any.
func (c *podPresets) Update(ctx context.Context, podPreset *settings.PodPreset, opts v1.UpdateOptions) (result *settings.PodPreset, err error) {
	result = &settings.PodPreset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("podpresets").
		Name(podPreset.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(podPreset).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the podPreset and deletes it. Returns an error if one occurs.
func (c *podPresets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podpresets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *podPresets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podpresets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched podPreset.
func (c *podPresets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *settings.PodPreset, err error) {
	result = &settings.PodPreset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("podpresets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
