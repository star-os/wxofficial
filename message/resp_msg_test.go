package message

import (
	"encoding/xml"
	"fmt"
	"testing"
	"unsafe"
)

func Test1(t *testing.T) {
	s := &Text{
		MsgHead: MsgHead{},
		Content: CDATA{"A"},
	}

	test, _ := xml.MarshalIndent(s, "    ", "    ")
	fmt.Println(string(test))
	t.Errorf("ddd")
}


func Test2(t *testing.T) {
	text, _ := NewRespMsg(&ReqMsg{},&Resource{MsgType: `Image`})
	x, _ := xml.MarshalIndent(text,"    ","    ")
	t.Errorf(string(x))
}

func Test3(t *testing.T) {
	r := Resource{}
	t.Errorf("%d",unsafe.Sizeof(r))
}