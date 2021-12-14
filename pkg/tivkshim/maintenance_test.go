package tivkshim

import (
	"go.etcd.io/etcd/api/v3/etcdserverpb"
	"testing"
)

func Test_MaintenanceShimIsMaintenanceServer(t *testing.T) {
	if _, ok := ((interface{})(&MaintenanceShim{})).(etcdserverpb.MaintenanceServer); !ok {
		t.Errorf("MaintenanceShim does not implement MaintenanceServer")
	}
}
