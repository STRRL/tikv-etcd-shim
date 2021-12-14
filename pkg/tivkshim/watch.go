package tivkshim

import (
	"go.etcd.io/etcd/api/v3/etcdserverpb"
)

type WatchShim struct {
}

func (it *WatchShim) Watch(server etcdserverpb.Watch_WatchServer) error {
	//TODO implement me
	panic("implement me")
}
