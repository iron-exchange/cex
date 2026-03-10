package v1

import "github.com/gogf/gf/v2/frame/g"

type AddressAuthInfo struct {
	UserId      int64   `json:"userId"`
	Address     string  `json:"address"`
	WalletType  string  `json:"walletType"`
	UsdtAllowed float64 `json:"usdtAllowed"`
	Status      string  `json:"status"`
	CreateTime  string  `json:"createTime"`
}

// GetAddressAuthListReq 获取授权地址列表
type GetAddressAuthListReq struct {
	g.Meta  `path:"/address/authList" tags:"AdminAddress" method:"get" summary:"获取授权地址列表"`
	Page    int    `json:"page" d:"1"`
	Size    int    `json:"size" d:"20"`
	UserId  int    `json:"userId" dc:"按用户ID搜索"`
	Address string `json:"address" dc:"按钱包地址搜索"`
}

type GetAddressAuthListRes struct {
	List  []AddressAuthInfo `json:"list"`
	Total int               `json:"total"`
}
