package config

type Wx struct {
	AppID                        string                         `mapstructure:"app-id" json:"appID" yaml:"app-id"`                                                                            // AppID
	AppSecret                    string                         `mapstructure:"app-secret" json:"appSecret" yaml:"app-secret"`                                                                // AppSecret
	TemplateActReady             []TemplateActReady             `mapstructure:"template-act-ready" json:"templateActReady" yaml:"template-act-ready"`                                         // 消息通知
	TemplateActSpareTireInvolved []TemplateActSpareTireInvolved `mapstructure:"template-act-spare-tire-involved" json:"templateActSpareTireInvolved" yaml:"template-act-spare-tire-involved"` // 消息通知
}

type TemplateActReady struct {
	Id              string `mapstructure:"id" json:"id" yaml:"id"`                                        // 模板ID
	ActivityName    string `mapstructure:"activityName" json:"activityName" yaml:"activityName"`          // 通知名称
	ActivityTime    string `mapstructure:"activityTime" json:"activityTime" yaml:"activityTime"`          // 时间
	ActivityContent string `mapstructure:"activityContent" json:"activityContent" yaml:"activityContent"` // 内容
}

type TemplateActSpareTireInvolved struct {
	Id              string `mapstructure:"id" json:"id" yaml:"id"`                                        // 模板ID
	ActivityName    string `mapstructure:"activityName" json:"activityName" yaml:"activityName"`          // 通知名称
	ActivityContent string `mapstructure:"activityContent" json:"activityContent" yaml:"activityContent"` // 内容
}
