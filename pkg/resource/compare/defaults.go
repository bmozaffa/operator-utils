package compare

import (
	"github.com/RHsyseng/operator-utils/pkg/resource"
	oappsv1 "github.com/openshift/api/apps/v1"
	routev1 "github.com/openshift/api/route/v1"
	corev1 "k8s.io/api/core/v1"
	"reflect"
)

type resourceComparator struct {
	defaultCompareFunc func(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool
	compareFuncMap     map[reflect.Type]func(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool
}

func (this *resourceComparator) SetDefaultComparator(compFunc func(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool) {
	this.defaultCompareFunc = compFunc
}

func (this *resourceComparator) GetDefaultComparator() func(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool {
	return this.defaultCompareFunc
}

func (this *resourceComparator) AddComparator(resourceType reflect.Type, compFunc func(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool) {
	this.compareFuncMap[resourceType] = compFunc
}

func (this *resourceComparator) GetComparator(resourceType reflect.Type) func(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool {
	return this.compareFuncMap[resourceType]
}

func (this *resourceComparator) Compare(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool {
	compareFunc := this.GetDefaultComparator()
	type1 := reflect.ValueOf(resource1).Elem().Type()
	type2 := reflect.ValueOf(resource2).Elem().Type()
	if type1 == type2 {
		if comparator, exists := this.compareFuncMap[type1]; exists {
			compareFunc = comparator
		}
	}
	return compareFunc(resource1, resource2)
}

func deepEquals(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool {
	struct1 := reflect.ValueOf(resource1).Elem().Type()
	if field1, found1 := struct1.FieldByName("Spec"); found1 {
		struct2 := reflect.ValueOf(resource2).Elem().Type()
		if field2, found2 := struct2.FieldByName("Spec"); found2 {
			return reflect.DeepEqual(field1, field2)
		}
	}
	return reflect.DeepEqual(resource1, resource2)
}

func defaultMap() map[reflect.Type]func(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool {
	equalsMap := make(map[reflect.Type]func(resource.KubernetesResource, resource.KubernetesResource) bool)
	equalsMap[reflect.TypeOf(oappsv1.DeploymentConfig{})] = equalsDeploymentConfig
	equalsMap[reflect.TypeOf(corev1.Service{})] = equalsService
	equalsMap[reflect.TypeOf(routev1.Route{})] = equalsRoute
	return equalsMap
}

func equalsDeploymentConfig(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool {
	dc1 := resource1.(*oappsv1.DeploymentConfig)
	dc2 := resource2.(*oappsv1.DeploymentConfig)
	if !reflect.DeepEqual(dc1.ObjectMeta, dc2.ObjectMeta) {
		return false
	}
	if !reflect.DeepEqual(dc1.TypeMeta, dc2.TypeMeta) {
		return false
	}
	if !reflect.DeepEqual(dc1.Spec, dc2.Spec) {
		return false
	}
	return true
}

func equalsService(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool {
	service1 := resource1.(*corev1.Service)
	service2 := resource2.(*corev1.Service)
	if !reflect.DeepEqual(service1.ObjectMeta, service2.ObjectMeta) {
		return false
	}
	if !reflect.DeepEqual(service1.TypeMeta, service2.TypeMeta) {
		return false
	}
	if !reflect.DeepEqual(service1.Spec, service2.Spec) {
		return false
	}
	return true
}

func equalsRoute(resource1 resource.KubernetesResource, resource2 resource.KubernetesResource) bool {
	route1 := resource1.(*routev1.Route)
	route2 := resource2.(*routev1.Route)
	if !reflect.DeepEqual(route1.ObjectMeta, route2.ObjectMeta) {
		return false
	}
	if !reflect.DeepEqual(route1.TypeMeta, route2.TypeMeta) {
		return false
	}
	if !reflect.DeepEqual(route1.Spec, route2.Spec) {
		return false
	}
	return true
}
