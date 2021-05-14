package wxofficial

var cfg Config	// 存放微信的配置信息

type Config struct {
	Token          string
	EncodingAESKey string
	AppId          string
	AppSecret      string
}


// InitConfig 初始化微信配置
func InitConfig(wxConfig *Config) {
	cfg = *wxConfig
}