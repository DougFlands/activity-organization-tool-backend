package config

type Wx struct {
	AppID       string        `mapstructure:"app-id" json:"appID" yaml:"app-id"`                   // AppID
	AppSecret   string        `mapstructure:"app-secret" json:"appSecret" yaml:"app-secret"`       // AppSecret
	TemplateMsg []TemplateMsg `mapstructure:"template-msg" json:"templateMsg" yaml:"template-msg"` // 消息通知
}

type TemplateMsg struct {
	Id              string `mapstructure:"id" json:"id" yaml:"id"`                                        // 模板ID
	ActivityName    string `mapstructure:"activityName" json:"activityName" yaml:"activityName"`          // 通知名称
	ActivityTime    string `mapstructure:"activityTime" json:"activityTime" yaml:"activityTime"`          // 时间
	ActivityContent string `mapstructure:"activityContent" json:"activityContent" yaml:"activityContent"` // 内容
}
