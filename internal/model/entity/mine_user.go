// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MineUser is the golang structure for table mine_user.
type MineUser struct {
	UserId    int64 `json:"user_id"    orm:"user_id"    description:"用户id"`
	Id        int64 `json:"id"         orm:"id"         description:"挖矿产品id"`
	TimeLimit int64 `json:"time_limit" orm:"time_limit" description:"限购次数"`
}
