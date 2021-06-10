package button

// MatchRule 是个性化中的匹配规则
type MatchRule struct {
	TagId              string `json:"tag_id,omitempty"`               // 用户标签，可以通过用户标签接口获取
	Sex                string `json:"sex,omitempty"`                  // 性别
	ClientPlatformType string `json:"client_platform_type,omitempty"` // 客户端版本
	Country            string `json:"country,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	Language           string `json:"language,omitempty"`
}
