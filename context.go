package wxofficial

import (
	"github.com/star-os/wxofficial/accesstoken"
	"sync"
)

const (
	UrlPrefix = "https://api.weixin.qq.com/cgi-bin/"
)

// Context 中控部分
type Context struct {
	Config

	// 控制部分
	accesstoken.AccessToken
	UpdateChan       chan int
	mAccessTokenLock sync.RWMutex // AccessToken的读写锁

}

// NewContext 初始化一个Context，并启动AccessToken的中控goroutine
func NewContext(token, encodingAESKey, appId, appSecret string) *Context {
	config := NewConfig(token, encodingAESKey, appId, appSecret)
	context := &Context{
		Config:           *config,
		AccessToken:      accesstoken.AccessToken{},
		UpdateChan:       make(chan int, 5),
		mAccessTokenLock: sync.RWMutex{},
	}

	go context.accessTokenUpdateDaemon()
	return context
}
