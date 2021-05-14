package message

import (
	"errors"
	"fmt"
)

var (
	errMsgType string = `Unknown MsgType %s` // 用于生成错误输出
)

// 定义被动回复微信的消息类型
type (
	// Text 文本消息
	Text struct {
		MsgHead
		Content CDATA `xml:"Content"`
	}

	// Image 图片消息
	Image struct {
		MsgHead
		MediaId CDATA `xml:"Image>MediaId"`
	}

	// Voice 语音消息
	Voice struct {
		MsgHead
		MediaId CDATA `xml:"Voice>MediaId"`
	}

	// Video 视频消息
	Video struct {
		MsgHead
		MediaId     CDATA `xml:"Video>MediaId"`
		Title       CDATA `xml:"Video>Title"`
		Description CDATA `xml:"Video>Description"`
	}

	// Music 音乐消息
	Music struct {
		MsgHead
		Title        CDATA `xml:"Music>Title"`
		Description  CDATA `xml:"Music>Description"`
		MusicUrl     CDATA `xml:"Music>MusicUrl"`
		HQMuSicUrl   CDATA `xml:"Music>HQMuSicUrl"`
		ThumbMediaId CDATA `xml:"Music>ThumbMediaId"`
	}

	// News 图文消息
	News struct {
		MsgHead
		ArticleCount uint `xml:"ArticleCount"`
		Articles     []Article
	}
)

// NewText 初始化Text消息
func NewText(mh *MsgHead, rsc *Resource) *Text {
	return &Text{
		MsgHead: *mh,
		Content: CDATA{rsc.Content},
	}
}

// NewImage 初始化Image消息
func NewImage(mh *MsgHead, rsc *Resource) *Image {
	return &Image{
		MsgHead: *mh,
		MediaId: CDATA{rsc.MediaId},
	}
}

// NewVoice 初始化Voice消息
func NewVoice(mh *MsgHead, rsc *Resource) *Voice {
	return &Voice{
		MsgHead: *mh,
		MediaId: CDATA{rsc.MediaId},
	}
}

// NewVideo 初始化Video消息
func NewVideo(mh *MsgHead, rsc *Resource) *Video {
	return &Video{
		MsgHead:     *mh,
		MediaId:     CDATA{rsc.MediaId},
		Title:       CDATA{rsc.Title},
		Description: CDATA{rsc.Description},
	}
}

// NewNews 初始化News消息
func NewNews(mh *MsgHead, rsc *Resource) *News {
	return &News{
		MsgHead:      *mh,
		ArticleCount: uint(len(rsc.Articles)),
		Articles:     rsc.Articles,
	}
}

// NewMusic 初始化Music消息
func NewMusic(mh *MsgHead, rsc *Resource) *Music {
	return &Music{
		MsgHead:      *mh,
		Title:        CDATA{rsc.Title},
		Description:  CDATA{rsc.Description},
		MusicUrl:     CDATA{rsc.MusicUrl},
		HQMuSicUrl:   CDATA{rsc.HQMuSicUrl},
		ThumbMediaId: CDATA{rsc.ThumbMediaId},
	}
}

// 辅助结构体
type (
	// Article 图文消息中的信息内容
	Article struct {
		Title       CDATA `xml:"Title"`
		Description CDATA `xml:"Description"`
		PicUrl      CDATA `xml:"PicUrl"`
		Url         CDATA `xml:"Url"`
	}

	// Resource 结构用于初始化RespMsg
	// 你只需要填充生成消息所需要的字段即可
	Resource struct {
		MsgType      string
		Content      string
		Title        string
		MediaId      string
		Description  string
		MusicUrl     string
		HQMuSicUrl   string
		ThumbMediaId string
		Articles     []Article
	}
)

//NewRespMsg 生成RespMsg
func NewRespMsg(req *ReqMsg, rsc *Resource) (interface{}, error) {
	var resp interface{}
	msgHead := NewMsgHead(req, rsc.MsgType)

	switch rsc.MsgType {
	case `Text`:
		resp = NewText(msgHead, rsc)
	case `Image`:
		resp = NewImage(msgHead, rsc)
	case `Voice`:
		resp = NewVoice(msgHead, rsc)
	case `Video`:
		resp = NewVideo(msgHead, rsc)
	case `Music`:
		resp = NewMusic(msgHead, rsc)
	case `News`:
		resp = NewNews(msgHead, rsc)
	default:
		errStr := fmt.Sprintf(errMsgType, rsc.MsgType)
		return nil, errors.New(errStr)
	}

	return resp, nil
}
