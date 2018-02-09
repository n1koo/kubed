/*
Copyright 2018 The Kubed Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/appscode/kubed/pkg/apis/kubed/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StuffLister helps list Stuffs.
type StuffLister interface {
	// List lists all Stuffs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Stuff, err error)
	// Stuffs returns an object that can list and get Stuffs.
	Stuffs(namespace string) StuffNamespaceLister
	StuffListerExpansion
}

// searchResultLister implements the StuffLister interface.
type searchResultLister struct {
	indexer cache.Indexer
}

// NewStuffLister returns a new StuffLister.
func NewStuffLister(indexer cache.Indexer) StuffLister {
	return &searchResultLister{indexer: indexer}
}

// List lists all Stuffs in the indexer.
func (s *searchResultLister) List(selector labels.Selector) (ret []*v1alpha1.Stuff, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Stuff))
	})
	return ret, err
}

// Stuffs returns an object that can list and get Stuffs.
func (s *searchResultLister) Stuffs(namespace string) StuffNamespaceLister {
	return searchResultNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StuffNamespaceLister helps list and get Stuffs.
type StuffNamespaceLister interface {
	// List lists all Stuffs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Stuff, err error)
	// Get retrieves the Stuff from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Stuff, error)
	StuffNamespaceListerExpansion
}

// searchResultNamespaceLister implements the StuffNamespaceLister
// interface.
type searchResultNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Stuffs in the indexer for a given namespace.
func (s searchResultNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Stuff, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Stuff))
	})
	return ret, err
}

// Get retrieves the Stuff from the indexer for a given namespace and name.
func (s searchResultNamespaceLister) Get(name string) (*v1alpha1.Stuff, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("stuff"), name)
	}
	return obj.(*v1alpha1.Stuff), nil
}