package access

import (
	"github.com/loststar527/wxofficial/config"
	"github.com/loststar527/wxofficial/util"
)

//校验signature
func CheckSignature(signature,nonce,timestamp string) bool {
	token := config.Token
	checkString := util.EncodingSha1(nonce, timestamp, token)
	return signature == checkString

}


