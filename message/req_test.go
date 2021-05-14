package message

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var temp string = `<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1351776360</CreateTime>
  <MsgType><![CDATA[location]]></MsgType>
  <Location_X>23.134521</Location_X>
  <Location_Y>113.358803</Location_Y>
  <Scale>20</Scale>
  <Label><![CDATA[<TOM>位置信息</TOM>]]></Label>
  <MsgId>1234567890123456</MsgId>
</xml>`

//
func Test(t *testing.T) {
	s := new(ReqMsg)
	tb := []byte(temp)
	xml.Unmarshal(tb, s)
	fmt.Printf("%+v\n", s)
	t.Errorf("d")
}
