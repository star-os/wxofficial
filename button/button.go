package button

import json "github.com/json-iterator/go"

type Button struct {
	Name      string   `json:"name"`       // 菜单标题
	Type      string   `json:"type"`       // 响应的动作类型
	Key       string   `json:"key"`        // 菜单key值，用于消息接口推送
	Url       string   `json:"url"`        // 网页链接
	MediaId   string   `json:"media_id"`   // 调用新增永久素材接口返回的合法media_id
	AppId     string   `json:"appid"`      // 小程序的appid
	PagePath  string   `json:"pagepath"`   // 小程序的页面路径
	SubButton []Button `json:"sub_button"` // 二级菜单，数量为0~5个
}

func JsonToButton(st string) (*Button, error) {
	bt := new(Button)
	err := json.Unmarshal([]byte(st), bt)
	if err != nil {
		return nil, err
	}
	return bt, nil
}
