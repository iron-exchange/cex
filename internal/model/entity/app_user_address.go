// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AppUserAddress is the golang structure for table app_user_address.
type AppUserAddress struct {
	Id           int64  `json:"id"            orm:"id"            description:""`
	UserId       int64  `json:"user_id"       orm:"user_id"       description:""`
	Symbol       string `json:"symbol"        orm:"symbol"        description:"钱包类型"`
	Address      string `json:"address"       orm:"address"       description:"钱包地址"`
	BinanceEmail string `json:"binance_email" orm:"binance_email" description:"币安子账号地址"`
}
