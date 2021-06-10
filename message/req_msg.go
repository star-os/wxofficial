package message

// 定义微信方发送的请求消息

// ReqMsg 微信发来的消息&事件
type ReqMsg struct {
	MsgHead

	// 消息，其中也包含一部分菜单-时间推送的字段
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
	Event        CDATA `xml:"Event"`    // 事件类型
	EventKey     CDATA `xml:"EventKey"` // 事件Key值
	MenuId       CDATA `xml:"MenuId"`   // 菜单ID，如果是个性化菜单，可以通过这个字段，知道是哪个规则的菜单被点击了
	ScanCodeInfo struct {
		ScanType CDATA `xml:"ScanType"` // 扫描类型
	} `xml:"ScanCodeInfo"` //扫描信息
	ScanResult   CDATA `xml:"ScanResult"` // 扫描结果
	SendPicsInfo struct {
		Count   int `xml:"Count"` // 发送的图片数量
		PicList struct {
			Item struct {
				PicMd5Sum CDATA `xml:"PicMd5Sum"` // 图片的MD5值，可用于验证收到的图片
			} `xml:"item"`
		} `xml:"PicList"` // 图片列表
	} `xml:"SendPicsInfo"` // 发送的图片信息
	Poiname   CDATA   `xml:"Poiname"`   // 朋友圈POI的名字，可能为空
	Ticket    CDATA   `xml:"Ticket"`    // 二维码的ticket
	Latitude  float64 `xml:"Latitude"`  // 纬度
	Longitude float64 `xml:"Longitude"` // 经度
	Precision float64 `xml:"Precision"` // 精度
}
