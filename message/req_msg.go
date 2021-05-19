package message

// 定义微信方发送的请求消息

// ReqMsg 微信发来的消息&事件
type ReqMsg struct {
	MsgHead

	// 消息
	MsgId        int64   `xml:"MsgId"`        // 消息Id
	PicUrl       CDATA   `xml:"PicUrl"`       // 图片链接
	MediaId      CDATA   `xml:"MediaId"`      // 媒体Id
	Format       CDATA   `xml:"Format"`       // 语音格式
	Recognition  CDATA   `xml:"Recognition"`  // 语音识别结果
	ThumbMediaId CDATA   `xml:"ThumbMediaId"` // 视频消息缩略图的媒体Id
	LocationX    float64 `xml:"Location_X"`   // 纬度
	LocationY    float64 `xml:"Location_Y"`   // 经度
	Scale        float64 `xml:"Scale"`        // 地图缩放大小
	Label        CDATA   `xml:"Label"`        // 位置信息
	Title        CDATA   `xml:"Title"`        // 消息标题
	Description  CDATA   `xml:"Description"`  // 消息描述
	Url          CDATA   `xml:"Url"`          // 消息链接

	// 事件
	Event     CDATA   `xml:"Event"`     // 事件类型
	EventKey  CDATA   `xml:"EventKey"`  // 事件Key值
	Ticket    CDATA   `xml:"Ticket"`    // 二维码的ticket
	Latitude  float64 `xml:"Latitude"`  // 纬度
	Longitude float64 `xml:"Longitude"` // 经度
	Precision float64 `xml:"Precision"` // 精度
}
