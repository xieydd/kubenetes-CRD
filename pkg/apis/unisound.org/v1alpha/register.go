package v1alpha

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	//"k8s.io/apimachinery/pkg/runtime/schema"
)

func addKnownType(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&QueueJob{},
		&QueueJobList{},)
	metav1.AddToGroupVersion(scheme,SchemeGroupVersion)
	return nil
}
