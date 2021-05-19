package context

import (
	"errors"
	"fmt"
	json "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	//ExpireForward 提前更新的时间
	ExpireForward = 1000
)

// AccessToken 用于存放向微信请求ACCESS_TOKEN以后的返回消息
type AccessToken struct {
	AToken    string `json:"access_token"` //凭证
	ExpiresIn int    `json:"expires_in"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

// GetAccessToken 获取AccessToken
func (c *Context) GetAccessToken() string {
	c.mAccessTokenLock.RLock()
	defer c.mAccessTokenLock.RUnlock()
	return c.AToken
}

// UpdateAccessToken 更新AccessToken
func (c *Context) UpdateAccessToken() error {
	token, err := c.pullAccessToken()
	if err != nil {
		return err
	}

	err = c.setAccessToken(token)
	if err != nil {
		return err
	}
	return nil
}

// 向微信公众号拉取AccessToken，返回json
func (c *Context) pullAccessToken() ([]byte, error) {
	url := c.buildUrl()

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("getAccessToken error: %s", err))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("getAccessToken error: %s", err))
	}

	return body, nil
}

// 用于生成请求AccessToken的URL
func (c *Context) buildUrl() string {
	return fmt.Sprintf(UrlTemplate, GrantTypeAccessToken, c.AppId, c.AppSecret)
}

// 设置AccessToken
func (c *Context) setAccessToken(tokenJson []byte) error {
	accToken := new(AccessToken)
	err := json.Unmarshal(tokenJson, accToken)
	if err != nil {
		return errors.New(fmt.Sprintf("setAccessToken error %s", err))
	}

	c.mAccessTokenLock.Lock()
	defer c.mAccessTokenLock.Unlock()
	c.AccessToken = *accToken
	return nil
}

// 检查AccessToken中错误信息
// true表示当前AccessToken正常
func (c *Context) checkAccessTokenError() (bool, error) {
	c.mAccessTokenLock.RLock()
	defer c.mAccessTokenLock.RUnlock()
	switch c.ErrCode {
	case 0: // 请求成功
		return true, nil
	case -1: // 系统繁忙
	default: // 其他问题
	}
	return false, errors.New(fmt.Sprintf("AccessTokenError: ErrCode: %d, ErrMsg: %s", c.ErrCode, c.ErrMsg))
}

// AccessTokenUpdateDaemon 用于自动刷新AccessToken
func (c *Context) AccessTokenUpdateDaemon() {
	ticker := time.NewTicker(time.Duration(c.ExpiresIn - ExpireForward))
	for {
		select {
		case <-ticker.C: // 计时主动刷新AccessToken
			err := c.UpdateAccessToken()
			if err != nil {
				// log
			}
		case <-c.UpdateChan: // 其他原因，被动刷新AccessToken
			err := c.UpdateAccessToken()
			if err != nil {
				// log
			}
			ticker.Reset(time.Duration(c.ExpiresIn - ExpireForward))
		}
		if flag, err := c.checkAccessTokenError(); flag == false {
			if err != nil {
				// log
			}
		}
	}
}
