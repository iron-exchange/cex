package asset

import (
	"context"

	v1 "GoCEX/api/app/v1"
	"GoCEX/internal/logic/asset"

	"github.com/gogf/gf/v2/frame/g"
)

type ControllerV1 struct{}

func New() *ControllerV1 {
	return &ControllerV1{}
}

func (c *ControllerV1) AssetList(ctx context.Context, req *v1.AssetListReq) (res *v1.AssetListRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userId").Uint64()
	return asset.New().GetAssetList(ctx, userId)
}

func (c *ControllerV1) WalletRecords(ctx context.Context, req *v1.WalletRecordReq) (res *v1.WalletRecordRes, err error) {
	userId := g.RequestFromCtx(ctx).GetCtxVar("userId").Uint64()
	return asset.New().GetWalletRecords(ctx, req, userId)
}
