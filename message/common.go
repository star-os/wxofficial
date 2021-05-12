package message

type CDATA struct {
	Text string
}

// 通用消息头
type MsgHead struct {
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int64
	MsgType      CDATA
}
