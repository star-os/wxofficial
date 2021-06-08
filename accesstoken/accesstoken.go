package accesstoken

import (
	"github.com/star-os/wxofficial/util"
)

const (
	//ExpireForward 提前更新的时间
	ExpireForward = 1000
)

// AccessToken 用于存放向微信请求ACCESS_TOKEN以后的返回消息
type AccessToken struct {
	AToken    string `json:"access_token"` //凭证
	ExpiresIn int64  `json:"expires_in"`
	util.ErrInfo
}
