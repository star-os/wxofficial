package context

import (
	"errors"
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/star-os/wxofficial/menu"
	"github.com/star-os/wxofficial/util"
)

const (
	MenuCreateUrl = `https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s`
	MenuGetUrl    = `https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=%s`
	MenuDeleteUrl = `https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s`
	// MenuGetConfig = `https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s`
)

// SetMenu 用于设置微信公众号的菜单
func (c *Context) SetMenu(buttons []menu.Button) error {
	url := c.buildUrlWithAT(MenuCreateUrl)
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
	url := c.buildUrlWithAT(MenuGetUrl)
	return util.HttpGet(url)
}

// DeleteMenu 删除接口
func (c *Context) DeleteMenu() error {
	url := c.buildUrlWithAT(MenuDeleteUrl)
	_, err := util.HttpGet(url)
	return err
}
