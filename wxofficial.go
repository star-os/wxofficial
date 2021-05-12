package wxofficial

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/loststar527/wxofficial/config"
	"github.com/loststar527/wxofficial/util"
)

func Do() {
	gin.Default()
	fmt.Println("as")
}

//校验signature，接入微信平台
func CheckSignature(signature, nonce, timestamp string) bool {
	token := config.Token
	checkString := util.Encoding(nonce, timestamp, token)
	return signature == checkString
}