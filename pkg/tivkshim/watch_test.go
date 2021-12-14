package tivkshim

import (
	"go.etcd.io/etcd/api/v3/etcdserverpb"
	"testing"
)

func Test_WatchShimIsWatchServer(t *testing.T) {
	if _, ok := ((interface{})(&WatchShim{})).(etcdserverpb.WatchServer); !ok {
		t.Errorf("WatchShim does not implement WatchServer")
	}
}
