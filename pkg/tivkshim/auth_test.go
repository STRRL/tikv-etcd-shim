package tivkshim

import (
	"go.etcd.io/etcd/api/v3/etcdserverpb"
	"testing"
)

func Test_AuthShimIsAuthServer(t *testing.T) {
	if _, ok := ((interface{})(&AuthShim{})).(etcdserverpb.AuthServer); !ok {
		t.Errorf("AuthShim does not implement AuthServer")
	}
}
