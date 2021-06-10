package util

// ErrInfo 是微信平台返回的通用错误信息定义
// 用于避免重复定义
type ErrInfo struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *ErrInfo) IsSuccess() bool {
	if e.ErrCode == 0 {
		return true
	}
	return false
}
