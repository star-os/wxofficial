package wxofficial

import (
	"errors"
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/star-os/wxofficial/accesstoken"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	UrlSuffixAccesstoken = UrlPrefix + "token?grant_type=%s&appid=%s&secret=%s"
	GrantTypeAccessToken = "client_credential"
)

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

// accessTokenUpdateDaemon 用于主动/被动刷新AccessToken
func (c *Context) accessTokenUpdateDaemon() {
	ticker := time.NewTicker(time.Duration(c.ExpiresIn - accesstoken.ExpireForward))

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
			ticker.Reset(time.Duration(c.ExpiresIn - accesstoken.ExpireForward))
		}

		go func() {
			err := c.checkAccessToken()
			if err != nil {
				// log
			}
		}()
	}
}

// 用于生成请求AccessToken的URL
func (c *Context) buildUrl() string {
	return fmt.Sprintf(UrlSuffixAccesstoken, GrantTypeAccessToken, c.AppId, c.AppSecret)
}

// 用于生成携带accesstoken的url
func (c *Context) getUrlWithAT(url string) string {
	return fmt.Sprintf(url, c.AToken)
}

// 检查AccessToken中错误信息
// true表示当前AccessToken正常
func (c *Context) checkAccessToken() error {
	c.mAccessTokenLock.RLock()
	defer c.mAccessTokenLock.RUnlock()
	if c.AToken == "" {
		c.UpdateChan <- c.ErrCode
		return errors.New(fmt.Sprintf("AccessToken Not Set"))
	}
	switch c.ErrCode {
	case 0: // 请求成功
		return nil
	case -1: // 系统繁忙
		fallthrough
	default: // 其他问题
		c.UpdateChan <- c.ErrCode
	}
	return errors.New(fmt.Sprintf("AccessTokenError: ErrCode: %d, ErrMsg: %s", c.ErrCode, c.ErrMsg))
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

// 设置AccessToken
func (c *Context) setAccessToken(tokenJson []byte) error {
	accToken := new(accesstoken.AccessToken)
	err := json.Unmarshal(tokenJson, accToken)
	if err != nil {
		return errors.New(fmt.Sprintf("setAccessToken error %s", err))
	}

	c.mAccessTokenLock.Lock()
	defer c.mAccessTokenLock.Unlock()
	c.AccessToken = *accToken
	return nil
}
