// Copyright 2019 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubeflow/xgboost-operator/pkg/apis/xgboost/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// XGBoostJobLister helps list XGBoostJobs.
type XGBoostJobLister interface {
	// List lists all XGBoostJobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.XGBoostJob, err error)
	// XGBoostJobs returns an object that can list and get XGBoostJobs.
	XGBoostJobs(namespace string) XGBoostJobNamespaceLister
	XGBoostJobListerExpansion
}

// xGBoostJobLister implements the XGBoostJobLister interface.
type xGBoostJobLister struct {
	indexer cache.Indexer
}

// NewXGBoostJobLister returns a new XGBoostJobLister.
func NewXGBoostJobLister(indexer cache.Indexer) XGBoostJobLister {
	return &xGBoostJobLister{indexer: indexer}
}

// List lists all XGBoostJobs in the indexer.
func (s *xGBoostJobLister) List(selector labels.Selector) (ret []*v1alpha1.XGBoostJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.XGBoostJob))
	})
	return ret, err
}

// XGBoostJobs returns an object that can list and get XGBoostJobs.
func (s *xGBoostJobLister) XGBoostJobs(namespace string) XGBoostJobNamespaceLister {
	return xGBoostJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// XGBoostJobNamespaceLister helps list and get XGBoostJobs.
type XGBoostJobNamespaceLister interface {
	// List lists all XGBoostJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.XGBoostJob, err error)
	// Get retrieves the XGBoostJob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.XGBoostJob, error)
	XGBoostJobNamespaceListerExpansion
}

// xGBoostJobNamespaceLister implements the XGBoostJobNamespaceLister
// interface.
type xGBoostJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all XGBoostJobs in the indexer for a given namespace.
func (s xGBoostJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.XGBoostJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.XGBoostJob))
	})
	return ret, err
}

// Get retrieves the XGBoostJob from the indexer for a given namespace and name.
func (s xGBoostJobNamespaceLister) Get(name string) (*v1alpha1.XGBoostJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("xgboostjob"), name)
	}
	return obj.(*v1alpha1.XGBoostJob), nil
}
