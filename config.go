package wxofficial

type Config struct {
	Token          string
	EncodingAESKey string
	AppId          string
	AppSecret      string
}

func NewConfig(token, encodingAESKey, appId, appSecret string) *Config {
	return &Config{
		Token:          token,
		EncodingAESKey: encodingAESKey,
		AppId:          appId,
		AppSecret:      appSecret,
	}
}
