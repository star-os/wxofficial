package wxofficial

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

// CheckSignature 用于校验signature，接入微信平台
func CheckSignature(signature, nonce, timestamp string) bool {
	checkString := Encoding(nonce, timestamp, cfg.Token)
	return signature == checkString
}

// Encoding 用于sha1加密
func Encoding(params ...string) string {
	sort.Strings(params)
	str := strings.Join(params, "")
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}

/**
 * Error
 * "signature=718f6ee15090aaa28188ed01b9c2f3b10681b4e1&echostr=3158210538242992831&timestamp=1620992680&nonce=314560874"
 * "signature=f5e78ba3918658c374f5cd9ce78e101ba7fe5174&echostr=3920999753636046847&timestamp=1620992727&nonce=1601008773"
 * "signature=8fee7d6745e4a28df88d8905f21b13f2bdc50135&echostr=6746511746692801247&timestamp=1620992732&nonce=203870580"
**/