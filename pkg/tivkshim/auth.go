package tivkshim

import (
	"context"
	"go.etcd.io/etcd/api/v3/etcdserverpb"
)

type AuthShim struct {
}

func (it *AuthShim) AuthEnable(ctx context.Context, request *etcdserverpb.AuthEnableRequest) (*etcdserverpb.AuthEnableResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) AuthDisable(ctx context.Context, request *etcdserverpb.AuthDisableRequest) (*etcdserverpb.AuthDisableResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) AuthStatus(ctx context.Context, request *etcdserverpb.AuthStatusRequest) (*etcdserverpb.AuthStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) Authenticate(ctx context.Context, request *etcdserverpb.AuthenticateRequest) (*etcdserverpb.AuthenticateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) UserAdd(ctx context.Context, request *etcdserverpb.AuthUserAddRequest) (*etcdserverpb.AuthUserAddResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) UserGet(ctx context.Context, request *etcdserverpb.AuthUserGetRequest) (*etcdserverpb.AuthUserGetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) UserList(ctx context.Context, request *etcdserverpb.AuthUserListRequest) (*etcdserverpb.AuthUserListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) UserDelete(ctx context.Context, request *etcdserverpb.AuthUserDeleteRequest) (*etcdserverpb.AuthUserDeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) UserChangePassword(ctx context.Context, request *etcdserverpb.AuthUserChangePasswordRequest) (*etcdserverpb.AuthUserChangePasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) UserGrantRole(ctx context.Context, request *etcdserverpb.AuthUserGrantRoleRequest) (*etcdserverpb.AuthUserGrantRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) UserRevokeRole(ctx context.Context, request *etcdserverpb.AuthUserRevokeRoleRequest) (*etcdserverpb.AuthUserRevokeRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) RoleAdd(ctx context.Context, request *etcdserverpb.AuthRoleAddRequest) (*etcdserverpb.AuthRoleAddResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) RoleGet(ctx context.Context, request *etcdserverpb.AuthRoleGetRequest) (*etcdserverpb.AuthRoleGetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) RoleList(ctx context.Context, request *etcdserverpb.AuthRoleListRequest) (*etcdserverpb.AuthRoleListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) RoleDelete(ctx context.Context, request *etcdserverpb.AuthRoleDeleteRequest) (*etcdserverpb.AuthRoleDeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) RoleGrantPermission(ctx context.Context, request *etcdserverpb.AuthRoleGrantPermissionRequest) (*etcdserverpb.AuthRoleGrantPermissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (it *AuthShim) RoleRevokePermission(ctx context.Context, request *etcdserverpb.AuthRoleRevokePermissionRequest) (*etcdserverpb.AuthRoleRevokePermissionResponse, error) {
	//TODO implement me
	panic("implement me")
}
