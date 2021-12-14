package tivkshim

import (
	"context"
	"go.etcd.io/etcd/api/v3/etcdserverpb"
)

type ClusterShim struct {
}

func (it *ClusterShim) MemberAdd(ctx context.Context, request *etcdserverpb.MemberAddRequest) (*etcdserverpb.MemberAddResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *ClusterShim) MemberRemove(ctx context.Context, request *etcdserverpb.MemberRemoveRequest) (*etcdserverpb.MemberRemoveResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *ClusterShim) MemberUpdate(ctx context.Context, request *etcdserverpb.MemberUpdateRequest) (*etcdserverpb.MemberUpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *ClusterShim) MemberList(ctx context.Context, request *etcdserverpb.MemberListRequest) (*etcdserverpb.MemberListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *ClusterShim) MemberPromote(ctx context.Context, request *etcdserverpb.MemberPromoteRequest) (*etcdserverpb.MemberPromoteResponse, error) {
	//TODO implement me
	panic("implement me")
}
