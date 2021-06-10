package wxofficial

import (
	"errors"
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/star-os/wxofficial/button"
	"github.com/star-os/wxofficial/util"
)

const (
	MenuCreateUrl = UrlPrefix + `menu/create?access_token=%s`
	MenuGetUrl    = UrlPrefix + `get_current_selfmenu_info?access_token=%s`
	MenuDeleteUrl = UrlPrefix + `menu/delete?access_token=%s`
	// MenuGetConfig = `https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%s`
	ConditionalCreateUrl = UrlPrefix + `menu/addconditional?access_token=%s`
	ConditionalDeleteUrl = UrlPrefix + `delconditional?access_token=%s`
	ConditionalTryMatchUrl = UrlPrefix + `menu/trymatch?access_token=%s`
)

var (
	menu = new(Menu)
)

// Menu 包含了按钮以及个性化菜单的数据，用于Marshall成JSON进行发送
type Menu struct {
	B                 []button.Button `json:"button,omitempty"`
	*button.MatchRule `json:"matchrule,omitempty"`
}

// AddButton 为菜单添加按钮
func AddButton(b ...button.Button) {
	menu.B = b
}

// SetMenuByJson 直接通过JSON设置菜单
func (c *Context) SetMenuByJson(j []byte) error {
	url := c.getUrlWithAT(MenuCreateUrl)
	data, err := util.HttpPost(url, j)
	if err != nil {
		return err
	}

	errInfo := new(util.ErrInfo)
	err = json.Unmarshal(data, errInfo)
	if errInfo.ErrCode != 0 {
		return errors.New(fmt.Sprintf("SetMenuError: %s. ErrCode = %d", errInfo.ErrMsg, errInfo.ErrCode))
	}
	return nil
}

// SetMenu 用于设置微信公众号的菜单（上传到微信公众号）
func (c *Context) SetMenu() error {
	url := c.getUrlWithAT(MenuCreateUrl)
	data, err := util.HttpPostJSON(url, menu)
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
	errJson, err := util.HttpGet(url)
	if err != nil {
		return err
	}

	errInfo := new(util.ErrInfo)
	err = json.Unmarshal(errJson, errInfo)
	if err != nil {
		return err
	}

	return nil
}

// DelConditional 删除个性化菜单
func (c *Context) DelConditional() error {
	return nil
}