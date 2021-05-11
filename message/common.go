package message

import ()

type CDATA struct {
	Text string
}

type msgHead struct {
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int64
	MsgType      CDATA
}
