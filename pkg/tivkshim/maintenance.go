package tivkshim

import (
	"context"
	"go.etcd.io/etcd/api/v3/etcdserverpb"
)

type MaintenanceShim struct {
}

func (it *MaintenanceShim) Alarm(ctx context.Context, request *etcdserverpb.AlarmRequest) (*etcdserverpb.AlarmResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *MaintenanceShim) Status(ctx context.Context, request *etcdserverpb.StatusRequest) (*etcdserverpb.StatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *MaintenanceShim) Defragment(ctx context.Context, request *etcdserverpb.DefragmentRequest) (*etcdserverpb.DefragmentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *MaintenanceShim) Hash(ctx context.Context, request *etcdserverpb.HashRequest) (*etcdserverpb.HashResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *MaintenanceShim) HashKV(ctx context.Context, request *etcdserverpb.HashKVRequest) (*etcdserverpb.HashKVResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *MaintenanceShim) Snapshot(request *etcdserverpb.SnapshotRequest, server etcdserverpb.Maintenance_SnapshotServer) error {
	//TODO implement me
	panic("implement me")
}

func (it *MaintenanceShim) MoveLeader(ctx context.Context, request *etcdserverpb.MoveLeaderRequest) (*etcdserverpb.MoveLeaderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *MaintenanceShim) Downgrade(ctx context.Context, request *etcdserverpb.DowngradeRequest) (*etcdserverpb.DowngradeResponse, error) {
	//TODO implement me
	panic("implement me")
}
