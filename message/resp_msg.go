package message

// 文本消息
type Text struct {
	MsgHead
	Content CDATA
}

// 图片消息
type Image struct {
	MsgHead
	Image struct {
		MediaId CDATA
	}
}

// 语音消息
type Voice struct {
	MsgHead
	Voice struct {
		MediaId CDATA
	}
}

// 视频消息
type Video struct {
	MsgHead
	Video struct {
		MediaId     CDATA
		Title       CDATA
		Description CDATA
	}
}

// 音乐消息
type Music struct {
	MsgHead
	Music struct {
		Title        CDATA
		Description  CDATA
		MusicUrl     CDATA
		HQMuSicUrl   CDATA
		ThumbMediaId CDATA
	}
}

// 图文消息
type News struct {
	MsgHead
	ArticleCount uint
	Articles     []Article
}

// 图文消息中的信息内容
type Article struct {
	Title       CDATA
	Description CDATA
	PicUrl      CDATA
	Url         CDATA
}
