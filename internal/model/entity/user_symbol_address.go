// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserSymbolAddress is the golang structure for table user_symbol_address.
type UserSymbolAddress struct {
	Id          int    `json:"id"           orm:"id"           description:"主键id"`
	UserId      int    `json:"user_id"      orm:"user_id"      description:"用户id"`
	Symbol      string `json:"symbol"       orm:"symbol"       description:"币种"`
	Address     string `json:"address"      orm:"address"      description:"充值地址"`
	SearchValue string `json:"search_value" orm:"search_value" description:""`
}
