package wxofficial

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/star-os/wxofficial/context"
)

// WeChat 调用各种方法的核心
type WeChat struct {
	context.Context
}

func Do() {
	gin.Default()
	fmt.Println("as")
}
