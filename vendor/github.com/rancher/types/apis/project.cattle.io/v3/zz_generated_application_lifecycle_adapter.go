package v3

import (
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type ApplicationLifecycle interface {
	Create(obj *Application) (runtime.Object, error)
	Remove(obj *Application) (runtime.Object, error)
	Updated(obj *Application) (runtime.Object, error)
}

type applicationLifecycleAdapter struct {
	lifecycle ApplicationLifecycle
}

func (w *applicationLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *applicationLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *applicationLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*Application))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *applicationLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*Application))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *applicationLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*Application))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewApplicationLifecycleAdapter(name string, clusterScoped bool, client ApplicationInterface, l ApplicationLifecycle) ApplicationHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(ApplicationGroupVersionResource)
	}
	adapter := &applicationLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *Application) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
