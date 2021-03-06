/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	application "tkestack.io/tke/api/application"
)

// AppHistoryLister helps list AppHistories.
// All objects returned here must be treated as read-only.
type AppHistoryLister interface {
	// List lists all AppHistories in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*application.AppHistory, err error)
	// AppHistories returns an object that can list and get AppHistories.
	AppHistories(namespace string) AppHistoryNamespaceLister
	AppHistoryListerExpansion
}

// appHistoryLister implements the AppHistoryLister interface.
type appHistoryLister struct {
	indexer cache.Indexer
}

// NewAppHistoryLister returns a new AppHistoryLister.
func NewAppHistoryLister(indexer cache.Indexer) AppHistoryLister {
	return &appHistoryLister{indexer: indexer}
}

// List lists all AppHistories in the indexer.
func (s *appHistoryLister) List(selector labels.Selector) (ret []*application.AppHistory, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*application.AppHistory))
	})
	return ret, err
}

// AppHistories returns an object that can list and get AppHistories.
func (s *appHistoryLister) AppHistories(namespace string) AppHistoryNamespaceLister {
	return appHistoryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppHistoryNamespaceLister helps list and get AppHistories.
// All objects returned here must be treated as read-only.
type AppHistoryNamespaceLister interface {
	// List lists all AppHistories in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*application.AppHistory, err error)
	// Get retrieves the AppHistory from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*application.AppHistory, error)
	AppHistoryNamespaceListerExpansion
}

// appHistoryNamespaceLister implements the AppHistoryNamespaceLister
// interface.
type appHistoryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppHistories in the indexer for a given namespace.
func (s appHistoryNamespaceLister) List(selector labels.Selector) (ret []*application.AppHistory, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*application.AppHistory))
	})
	return ret, err
}

// Get retrieves the AppHistory from the indexer for a given namespace and name.
func (s appHistoryNamespaceLister) Get(name string) (*application.AppHistory, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(application.Resource("apphistory"), name)
	}
	return obj.(*application.AppHistory), nil
}
