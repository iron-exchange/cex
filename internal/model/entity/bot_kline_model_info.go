// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BotKlineModelInfo is the golang structure for table bot_kline_model_info.
type BotKlineModelInfo struct {
	Id       int64   `json:"id"        orm:"id"        description:"id"`
	ModelId  int64   `json:"model_id"  orm:"model_id"  description:"t_bot_kline_model 的主键"`
	DateTime int64   `json:"date_time" orm:"date_time" description:"时间戳"`
	Open     float64 `json:"open"      orm:"open"      description:"开盘价"`
	Close    float64 `json:"close"     orm:"close"     description:"封盘价"`
	High     float64 `json:"high"      orm:"high"      description:"最高价"`
	Low      float64 `json:"low"       orm:"low"       description:"最低价"`
	X        string  `json:"x"         orm:"x"         description:"x轴"`
	Y        string  `json:"y"         orm:"y"         description:"y轴"`
}
