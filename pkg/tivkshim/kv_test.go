package tivkshim

import (
	"go.etcd.io/etcd/api/v3/etcdserverpb"
	"testing"
)

func Test_KVShimIsKVServer(t *testing.T) {
	if _, ok := ((interface{})(&KVShim{}).(etcdserverpb.KVServer)); !ok {
		t.Errorf("KVShim does not implement KVServer")
	}
}
