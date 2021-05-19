package message

// Package message 中的通用部分

import (
	"encoding/xml"
	"time"
)

type CDATA struct {
	Text string `xml:",cdata"`
}

// MsgHead 通用消息头
type MsgHead struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
}

// NewMsgHead 通过ReqMsg构建RespMsg
func NewMsgHead(req *ReqMsg, msgType string) *MsgHead {
	return &MsgHead{
		ToUserName:   req.FromUserName,
		FromUserName: req.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      CDATA{msgType},
	}
}
