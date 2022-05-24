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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/stolostron/multicluster-mesh-addon/apis/client/clientset/versioned"
	internalinterfaces "github.com/stolostron/multicluster-mesh-addon/apis/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/stolostron/multicluster-mesh-addon/apis/client/listers/mesh/v1alpha1"
	meshv1alpha1 "github.com/stolostron/multicluster-mesh-addon/apis/mesh/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MeshFederationInformer provides access to a shared informer and lister for
// MeshFederations.
type MeshFederationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MeshFederationLister
}

type meshFederationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMeshFederationInformer constructs a new informer for MeshFederation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMeshFederationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMeshFederationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMeshFederationInformer constructs a new informer for MeshFederation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMeshFederationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeshV1alpha1().MeshFederations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeshV1alpha1().MeshFederations(namespace).Watch(context.TODO(), options)
			},
		},
		&meshv1alpha1.MeshFederation{},
		resyncPeriod,
		indexers,
	)
}

func (f *meshFederationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMeshFederationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *meshFederationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&meshv1alpha1.MeshFederation{}, f.defaultInformer)
}

func (f *meshFederationInformer) Lister() v1alpha1.MeshFederationLister {
	return v1alpha1.NewMeshFederationLister(f.Informer().GetIndexer())
}
