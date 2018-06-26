// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	fimcontroller_v1alpha1 "clustergarage.io/fim-k8s/pkg/apis/fimcontroller/v1alpha1"
	versioned "clustergarage.io/fim-k8s/pkg/client/clientset/versioned"
	internalinterfaces "clustergarage.io/fim-k8s/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "clustergarage.io/fim-k8s/pkg/client/listers/fimcontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FimListenerInformer provides access to a shared informer and lister for
// FimListeners.
type FimListenerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FimListenerLister
}

type fimListenerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFimListenerInformer constructs a new informer for FimListener type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFimListenerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFimListenerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFimListenerInformer constructs a new informer for FimListener type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFimListenerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FimcontrollerV1alpha1().FimListeners(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FimcontrollerV1alpha1().FimListeners(namespace).Watch(options)
			},
		},
		&fimcontroller_v1alpha1.FimListener{},
		resyncPeriod,
		indexers,
	)
}

func (f *fimListenerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFimListenerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fimListenerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&fimcontroller_v1alpha1.FimListener{}, f.defaultInformer)
}

func (f *fimListenerInformer) Lister() v1alpha1.FimListenerLister {
	return v1alpha1.NewFimListenerLister(f.Informer().GetIndexer())
}