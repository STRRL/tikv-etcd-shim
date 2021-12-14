package tivkshim

import (
	"go.etcd.io/etcd/api/v3/etcdserverpb"
	"testing"
)

func Test_LeaseShimIsLeaseServer(t *testing.T) {
	if _, ok := ((interface{})(&LeaseShim{})).(etcdserverpb.LeaseServer); !ok {
		t.Errorf("LeaseShim does not implement LeaseServer")
	}
}
