// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserCoin is the golang structure for table user_coin.
type UserCoin struct {
	Id     int64  `json:"id"      orm:"id"      description:"主键"`
	UserId int64  `json:"user_id" orm:"user_id" description:"用户id"`
	Coin   string `json:"coin"    orm:"coin"    description:"币种"`
	Icon   string `json:"icon"    orm:"icon"    description:"图标"`
}
