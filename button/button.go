package button

import json "github.com/json-iterator/go"

type Button struct {
	Name      string   `json:"name"`                 // 菜单标题
	Type      string   `json:"type"`                 // 响应的动作类型
	Key       string   `json:"key,omitempty"`        // 菜单key值，用于消息接口推送
	Url       string   `json:"url,omitempty"`        // 网页链接
	MediaId   string   `json:"media_id,omitempty"`   // 调用新增永久素材接口返回的合法media_id
	AppId     string   `json:"appid,omitempty"`      // 小程序的appid
	PagePath  string   `json:"pagepath,omitempty"`   // 小程序的页面路径
	SubButton []Button `json:"sub_button,omitempty"` // 二级菜单，数量为0~5个
}

//Click 创建click按钮
func Click(name, key string) *Button {
	return &Button{
		Name: name,
		Type: "click",
		Key:  key,
	}
}

// View 创建view按钮
func View(name, url string) *Button {
	return &Button{
		Name: name,
		Type: "view",
		Url:  url,
	}
}

// MiniProgram 创建小程序按钮
func MiniProgram(name, url, appid, pagepath string) *Button {
	return &Button{
		Name:     name,
		Type:     "miniprogram",
		Url:      url,
		AppId:    appid,
		PagePath: pagepath,
	}
}

// ScanCodePush 创建scancode_push按钮
func ScanCodePush(name, key string) *Button {
	return &Button{
		Name: name,
		Type: "scancode_push",
		Key:  key,
	}
}

// ScanCodeWaitMsg 创建scancode_waitmsg按钮
func ScanCodeWaitMsg(name, key string) *Button {
	return &Button{
		Name: name,
		Type: "scancode_waitmsg",
		Key:  key,
	}
}

func PicSysPhoto(name, key string) *Button {
	return &Button{
		Name: name,
		Type: "pic_sysphoto",
		Key:  key,
	}
}

func PicWeiXin(name, key string) *Button {
	return &Button{
		Name: name,
		Type: "pic_weixin",
		Key:  key,
	}
}
func LocationSelect(name, key string) *Button {
	return &Button{
		Name: name,
		Type: "location_select",
		Key:  key,
	}
}

// MediaId 图片
func MediaId(name, mediaid string) *Button {
	return &Button{
		Name:    name,
		Type:    "media_id",
		MediaId: mediaid,
	}
}

// ViewLimited 图文消息
func ViewLimited(name, mediaid string) *Button {
	return &Button{
		Name:    name,
		Type:    "view_limited",
		MediaId: mediaid,
	}
}

func JsonToButton(st string) (*Button, error) {
	bt := new(Button)
	err := json.Unmarshal([]byte(st), bt)
	if err != nil {
		return nil, err
	}
	return bt, nil
}
