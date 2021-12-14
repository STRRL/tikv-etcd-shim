package tivkshim

import (
	"context"
	"go.etcd.io/etcd/api/v3/etcdserverpb"
)

type KVShim struct {
}

func (it *KVShim) Range(ctx context.Context, request *etcdserverpb.RangeRequest) (*etcdserverpb.RangeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *KVShim) Put(ctx context.Context, request *etcdserverpb.PutRequest) (*etcdserverpb.PutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *KVShim) DeleteRange(ctx context.Context, request *etcdserverpb.DeleteRangeRequest) (*etcdserverpb.DeleteRangeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *KVShim) Txn(ctx context.Context, request *etcdserverpb.TxnRequest) (*etcdserverpb.TxnResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *KVShim) Compact(ctx context.Context, request *etcdserverpb.CompactionRequest) (*etcdserverpb.CompactionResponse, error) {
	//TODO implement me
	panic("implement me")
}
