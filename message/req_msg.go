package message

// 微信发来的消息&事件
type ReqMsg struct {
	MsgHead

	// 消息
	MsgId        int64   // 消息Id
	PicUrl       CDATA   // 图片链接
	MediaId      CDATA   // 媒体Id
	Format       CDATA   // 语音格式
	Recognition  CDATA   // 语音识别结果
	ThumbMediaId CDATA   // 视频消息缩略图的媒体Id
	Location_X   float64 // 纬度
	Location_Y   float64 // 经度
	Scale        float64 // 地图缩放大小
	Label        CDATA   // 位置信息
	Title        CDATA   // 消息标题
	Description  CDATA   // 消息描述
	Url          CDATA   // 消息链接

	// 事件
	Event     CDATA   // 事件类型
	EventKey  CDATA   // 事件Key值
	Ticket    CDATA   // 二维码的ticket
	Latitude  float64 // 纬度
	Longitude float64 // 经度
	Precision float64 // 精度
}
