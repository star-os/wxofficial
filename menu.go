package wxofficial

import (
	"errors"
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/star-os/wxofficial/button"
	"github.com/star-os/wxofficial/util"
)

const (
	MenuCreateUrl = UrlPrefix + `button/create?access_token=%s`
	MenuGetUrl    = UrlPrefix + `get_current_selfmenu_info?access_token=%s`
	MenuDeleteUrl = UrlPrefix + `button/delete?access_token=%s`
	// MenuGetConfig = `https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s`
)

// SetMenu 用于设置微信公众号的菜单
func (c *Context) SetMenu(buttons []button.Button) error {
	url := c.getUrlWithAT(MenuCreateUrl)
	data, err := util.HttpPostJSON(url, buttons)
	if err != nil {
		return err
	}

	errInfo := new(util.ErrInfo)
	err = json.Unmarshal(data, errInfo)
	if err != nil {
		return err
	}

	if errInfo.ErrCode != 0 {
		return errors.New(fmt.Sprintf("SetMenuError: %s", errInfo.ErrMsg))
	}
	return nil
}

// QueryMenu 查询接口
func (c *Context) QueryMenu() ([]byte, error) {
	url := c.getUrlWithAT(MenuGetUrl)
	return util.HttpGet(url)
}

// DeleteMenu 删除接口
func (c *Context) DeleteMenu() error {
	url := c.getUrlWithAT(MenuDeleteUrl)
	_, err := util.HttpGet(url)
	return err
}
