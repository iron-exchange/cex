// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Markets is the golang structure for table markets.
type Markets struct {
	Id          int64  `json:"id"           orm:"id"           description:""`
	Slug        string `json:"slug"         orm:"slug"         description:"交易所名称(ID)"`
	Fullname    string `json:"fullname"     orm:"fullname"     description:"交易所全称"`
	WebsiteUrl  string `json:"website_url"  orm:"website_url"  description:"交易所官网链接"`
	Status      string `json:"status"       orm:"status"       description:"状态: [enable, disable]. disable为停止更新数据"`
	Kline       int    `json:"kline"        orm:"kline"        description:"是否接入K线数据"`
	Spot        int    `json:"spot"         orm:"spot"         description:"是否支持现货"`
	Futures     int    `json:"futures"      orm:"futures"      description:"是否支持期货"`
	SearchValue string `json:"search_value" orm:"search_value" description:""`
}
