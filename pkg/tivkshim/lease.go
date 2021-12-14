package tivkshim

import (
	"context"
	"go.etcd.io/etcd/api/v3/etcdserverpb"
)

type LeaseShim struct {
}

func (it *LeaseShim) LeaseGrant(ctx context.Context, request *etcdserverpb.LeaseGrantRequest) (*etcdserverpb.LeaseGrantResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *LeaseShim) LeaseRevoke(ctx context.Context, request *etcdserverpb.LeaseRevokeRequest) (*etcdserverpb.LeaseRevokeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *LeaseShim) LeaseKeepAlive(server etcdserverpb.Lease_LeaseKeepAliveServer) error {
	//TODO implement me
	panic("implement me")
}

func (it *LeaseShim) LeaseTimeToLive(ctx context.Context, request *etcdserverpb.LeaseTimeToLiveRequest) (*etcdserverpb.LeaseTimeToLiveResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *LeaseShim) LeaseLeases(ctx context.Context, request *etcdserverpb.LeaseLeasesRequest) (*etcdserverpb.LeaseLeasesResponse, error) {
	//TODO implement me
	panic("implement me")
}
