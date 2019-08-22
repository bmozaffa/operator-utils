package compare

import (
	"github.com/RHsyseng/operator-utils/pkg/resource"
	"reflect"
)

type ResourceDelta struct {
	Added   []resource.KubernetesResource
	Updated []resource.KubernetesResource
	Removed []resource.KubernetesResource
}

type ResourceComparator interface {
	SetDefaultComparator(compFunc func(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool)
	GetDefaultComparator() func(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool
	AddComparator(resourceType reflect.Type, compFunc func(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool)
	GetComparator(resourceType reflect.Type) func(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool
	Compare(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool
}

func DefaultComparator() ResourceComparator {
	return &resourceComparator{
		deepEquals,
		defaultMap(),
	}
}

func SimpleComparator() ResourceComparator {
	return &resourceComparator{
		deepEquals,
		make(map[reflect.Type]func(resource.KubernetesResource, resource.KubernetesResource) bool),
	}
}
