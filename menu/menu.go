package menu

import "github.com/star-os/wxofficial/context"

const (
	MenuCreateUrl = `https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s`
	MenuGetUrl = `https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=%s`
	MenuDeleteUrl = `https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s`
)

type Menu struct {
	context.Context
}

func (m *Menu) CreateMenu() {

}

