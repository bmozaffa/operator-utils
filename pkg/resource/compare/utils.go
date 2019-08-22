package compare

import (
	"github.com/RHsyseng/operator-utils/pkg/resource"
	"reflect"
)

type mapBuilder struct {
	resourceMap map[reflect.Type][]resource.KubernetesResource
}

func NewMapBuilder() *mapBuilder {
	this := &mapBuilder{resourceMap: make(map[reflect.Type][]resource.KubernetesResource)}
	return this
}

func (this *mapBuilder) Map() map[reflect.Type][]resource.KubernetesResource {
	return this.resourceMap
}

func (this *mapBuilder) SameTypeItems(items ...resource.KubernetesResource) *mapBuilder {
	if len(items) == 0 {
		return this
	}
	resources := this.resourceMap[reflect.ValueOf(items[0]).Elem().Type()]
	resources = append(resources, items...)
	return this
}

//func (this *mapBuilder) IdenticalArrays(arrays ...[]resource.KubernetesResource) *mapBuilder {
//	for _, resArray := range arrays {
//		if len(resArray) > 0 {
//			resType := reflect.ValueOf(resArray[0]).Elem().Type()
//
//		}
//	}
//	if len(arrays) == 0 {
//		return this
//	}
//	reflect.ValueOf(objects[0]).Elem().Type()
//}
