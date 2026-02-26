// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Symbols is the golang structure for table symbols.
type Symbols struct {
	Id       int64  `json:"id"       orm:"id"       description:""`
	Slug     string `json:"slug"     orm:"slug"     description:"币种名称（ID）"`
	Symbol   string `json:"symbol"   orm:"symbol"   description:"币种符号"`
	Fullname string `json:"fullname" orm:"fullname" description:"币种全称"`
	LogoUrl  string `json:"logo_url" orm:"logo_Url" description:"图标链接"`
	Fiat     int    `json:"fiat"     orm:"fiat"     description:"是否法定货币"`
}
