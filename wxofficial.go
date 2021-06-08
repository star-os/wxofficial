package wxofficial

import (
	"github.com/star-os/wxofficial/context"
)

// WeChat 调用各种方法的核心
type WeChat struct {
	context.Context
}

func NewWeChat() *WeChat {
	return &WeChat{}
}

func Do() {

}
