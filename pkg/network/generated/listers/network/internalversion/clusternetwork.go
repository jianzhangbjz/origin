// This file was automatically generated by lister-gen

package internalversion

import (
	network "github.com/openshift/origin/pkg/network/apis/network"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterNetworkLister helps list ClusterNetworks.
type ClusterNetworkLister interface {
	// List lists all ClusterNetworks in the indexer.
	List(selector labels.Selector) (ret []*network.ClusterNetwork, err error)
	// Get retrieves the ClusterNetwork from the index for a given name.
	Get(name string) (*network.ClusterNetwork, error)
	ClusterNetworkListerExpansion
}

// clusterNetworkLister implements the ClusterNetworkLister interface.
type clusterNetworkLister struct {
	indexer cache.Indexer
}

// NewClusterNetworkLister returns a new ClusterNetworkLister.
func NewClusterNetworkLister(indexer cache.Indexer) ClusterNetworkLister {
	return &clusterNetworkLister{indexer: indexer}
}

// List lists all ClusterNetworks in the indexer.
func (s *clusterNetworkLister) List(selector labels.Selector) (ret []*network.ClusterNetwork, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*network.ClusterNetwork))
	})
	return ret, err
}

// Get retrieves the ClusterNetwork from the index for a given name.
func (s *clusterNetworkLister) Get(name string) (*network.ClusterNetwork, error) {
	key := &network.ClusterNetwork{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(network.Resource("clusternetwork"), name)
	}
	return obj.(*network.ClusterNetwork), nil
}
