package accesstoken

import (
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/star-os/wxofficial"
	"github.com/star-os/wxofficial/util"
	"testing"
)

func HelpInitConfig() *wxofficial.Config {
	cfg := &wxofficial.Config{
		Token:          "TestToken",
		EncodingAESKey: "TestEncodingAESKey",
		AppId:          "wx02e9b7a7e601ce1b",
		AppSecret:      "5efa3cf01ec1c6cf58065a8f65ad54ed",
	}
	return cfg
}

func HelpInitContext() *wxofficial.Context {
	cfg := HelpInitConfig()
	accessToken := AccessToken{
		AToken:    "",
		ExpiresIn: 0,
		ErrInfo:   util.ErrInfo{},
	}

	context := &wxofficial.Context{
		Config:      *cfg,
		AccessToken: accessToken,
		UpdateChan:  nil,
	}
	return context
}

// 得到空的内容
func TestContext_GetAccessToken(t *testing.T) {
	ctx := wxofficial.InitContext(*HelpInitConfig())

	token := ctx.GetAccessToken()
	fmt.Println(ctx)
	fmt.Println(token)
	t.Errorf("dd")
}

func Test(t *testing.T) {
	js := `{"access_token":"ACCESS_TOKEN","expires_in":7200}`
	At := &AccessToken{
		AToken:    "s",
		ExpiresIn: 0,
		ErrInfo:   util.ErrInfo{},
	}
	json.Unmarshal([]byte(js), At)
	fmt.Printf("%#v", At)
}
