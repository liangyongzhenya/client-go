// Code generated by client-gen. DO NOT EDIT.

package fake

import "k8s.io/client-go/testing"


type FakeCoreV1Interface interface {
	PodsGetter
}

// CoreV1Client is used to interact with features provided by the  group.
type FakeCoreV1 struct {
	*testing.Fake
}

func (c *FakeCoreV1) Pods(namespace string) PodInterface {
	return &FakePods{
		Fake: c,
		ns: namespace,
	}
}
