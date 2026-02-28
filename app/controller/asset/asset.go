package asset

import (
	"context"

	v1 "GoCEX/app/api"
	"GoCEX/internal/logic/asset"
	"GoCEX/internal/service/middleware"

	"github.com/gogf/gf/v2/util/gconv"
)

type ControllerV1 struct{}

func New() *ControllerV1 {
	return &ControllerV1{}
}

func (c *ControllerV1) AssetList(ctx context.Context, req *v1.AssetListReq) (res *v1.AssetListRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return asset.New().GetAssetList(ctx, userId)
}

func (c *ControllerV1) WalletRecords(ctx context.Context, req *v1.WalletRecordReq) (res *v1.WalletRecordRes, err error) {
	userId := gconv.Uint64(middleware.Auth.GetIdentity(ctx))
	return asset.New().GetWalletRecords(ctx, req, userId)
}
