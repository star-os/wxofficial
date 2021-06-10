package wxofficial

import (
	"github.com/star-os/wxofficial/accesstoken"
	"github.com/star-os/wxofficial/util"
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
		Config: *config,
		AccessToken: accesstoken.AccessToken{
			AToken:    "",
			ExpiresIn: 7200,
			ErrInfo:   util.ErrInfo{},
		},
		UpdateChan:       make(chan int, 5),
		mAccessTokenLock: sync.RWMutex{},
	}

	err := context.UpdateAccessToken()
	if err != nil {
		panic(err)
	}
	go context.accessTokenUpdateDaemon()
	return context
}
