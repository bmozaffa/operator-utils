package kubernetes

import (
	"errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func CustomResourceDefinitionExists(gvk schema.GroupVersionKind, cfg *rest.Config) (bool, error) {
	var err error
	if cfg == nil {
		cfg, err = config.GetConfig()
		if err != nil {
			return false, err
		}
	}
	client, err := discovery.NewDiscoveryClientForConfig(cfg)
	if err != nil {
		return false, err
	}
	api, err := client.ServerResourcesForGroupVersion(gvk.GroupVersion().String())
	if err != nil {
		return false, err
	}
	for _, a := range api.APIResources {
		if a.Kind == gvk.Kind {
			return true, nil
		}
	}
	return false, errors.New(gvk.String() + " Kind not found ")
}
