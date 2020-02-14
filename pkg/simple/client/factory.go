/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/
package client

import (
	"rest-shell/pkg/simple/client/k8s"
	"fmt"
	"sync"
)

type ClientSetNotEnabledError struct {
	err error
}

func (e ClientSetNotEnabledError) Error() string {
	return fmt.Sprintf("client set not enabled: %v", e.err)
}

type ClientSetOptions struct {
	kubernetesOptions   *k8s.KubernetesOptions
}

func NewClientSetOptions() *ClientSetOptions {
	return &ClientSetOptions{
		kubernetesOptions:   k8s.NewKubernetesOptions(),
	}
}

func (c *ClientSetOptions) SetKubernetesOptions(options *k8s.KubernetesOptions) *ClientSetOptions {
	c.kubernetesOptions = options
	return c
}

// ClientSet provide best of effort service to initialize clients,
// but there is no guarantee to return a valid client instance,
// so do validity check before use
type ClientSet struct {
	csoptions *ClientSetOptions
	stopCh    <-chan struct{}
	k8sClient           *k8s.KubernetesClient
}

var mutex sync.Mutex

// global clientsets instance
var sharedClientSet *ClientSet

func ClientSets() *ClientSet {
	return sharedClientSet
}

func NewClientSetFactory(c *ClientSetOptions, stopCh <-chan struct{}) *ClientSet {
	sharedClientSet = &ClientSet{csoptions: c, stopCh: stopCh}
	if c.kubernetesOptions != nil {
		sharedClientSet.k8sClient = k8s.NewKubernetesClientOrDie(c.kubernetesOptions)
	}

	return sharedClientSet
}

// jinxin added
func CreateClientSet(conf *ClientSetOptions, stopCh <-chan struct{}) error {
	NewClientSetFactory(conf, stopCh)
	return nil
}
