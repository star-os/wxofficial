package context

import (
	"sync"
)

const (
	UrlTemplate          = "https://api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s"
	GrantTypeAccessToken = "client_credential"
	// 需要保存错误信息嘛

)

// Context 中控部分
type Context struct {
	Config

	// 控制部分
	AccessToken
	UpdateChan       chan int
	mAccessTokenLock sync.RWMutex // AccessToken的读写锁

}

// InitContext 初始化一个Context，并启动AccessToken的中控goroutine
func InitContext(cfg Config) *Context {
	c := &Context{
		Config: cfg,
	}

	// 初始化AccessToken部分
	err := c.UpdateAccessToken()
	c.UpdateChan = make(chan int, 5)
	if err != nil {
		panic(err)
	}
	go c.AccessTokenUpdateDaemon()
	return c
}