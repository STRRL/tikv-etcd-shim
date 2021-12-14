package tivkshim

import (
	"go.etcd.io/etcd/api/v3/etcdserverpb"
	"testing"
)

func Test_ClusterShimIsClusterServer(t *testing.T) {
	if _, ok := ((interface{})(&ClusterShim{}).(etcdserverpb.ClusterServer)); !ok {
		t.Errorf("ClusterShim does not implement ClusterServer")
	}
}
