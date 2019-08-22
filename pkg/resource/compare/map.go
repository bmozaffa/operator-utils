package compare

import (
	"github.com/RHsyseng/operator-utils/pkg/resource"
	"reflect"
	logs "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var logger = logs.Log.WithName("comparator")

type MapComparator struct {
	Comparator ResourceComparator
}

func (this *MapComparator) Compare(deployed map[reflect.Type][]resource.KubernetesResource, requested map[reflect.Type][]resource.KubernetesResource) map[reflect.Type]ResourceDelta {
	delta := make(map[reflect.Type]ResourceDelta)
	for deployedType, deployedArray := range deployed {
		requestedArray := requested[deployedType]
		logger := logger.WithValues("delta", deployedType)
		for _, deployed := range deployedArray {
			logger.Info("Deployed item", "name", deployed.GetName())
		}
		for _, requested := range requestedArray {
			logger.Info("Requested item", "name", requested.GetName())
		}
		delta[deployedType] = this.compareArrays(deployedType, requestedArray, deployedArray)
	}
	return delta
}

func (this *MapComparator) compareArrays(resourceType reflect.Type, requested []resource.KubernetesResource, deployed []resource.KubernetesResource) ResourceDelta {
	requestedMap := getObjectMap(requested)
	deployedMap := getObjectMap(deployed)
	var added []resource.KubernetesResource
	var updated []resource.KubernetesResource
	var removed []resource.KubernetesResource
	for name, requestedObject := range requestedMap {
		deployedObject := deployedMap[name]
		if deployedObject == nil {
			added = append(added, requestedObject)
		} else if !this.Comparator.Compare(requestedObject, deployedObject) {
			updated = append(updated, requestedObject)
		}
	}
	for name, deployedObject := range deployedMap {
		if requestedMap[name] == nil {
			removed = append(removed, deployedObject)
		}
	}
	return ResourceDelta{
		Added:   added,
		Updated: updated,
		Removed: removed,
	}
}

func getObjectMap(objects []resource.KubernetesResource) map[string]resource.KubernetesResource {
	objectMap := make(map[string]resource.KubernetesResource)
	for index := range objects {
		objectMap[objects[index].GetName()] = objects[index]
	}
	return objectMap
}
